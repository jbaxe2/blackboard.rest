package blackboard_rest

import (
  "net/url"
  "strings"

  "github.com/jbaxe2/blackboard.rest/_scaffolding"
  "github.com/jbaxe2/blackboard.rest/_scaffolding/config"
  "github.com/jbaxe2/blackboard.rest/course_group_users"
  "github.com/jbaxe2/blackboard.rest/oauth2"
)

/**
 * The [CourseGroupUsers] interface...
 */
type CourseGroupUsers interface {
  GetGroupMemberships (
    courseId string, groupId string,
  ) ([]course_group_users.GroupMembership, error)

  GetGroupMembership (
    courseId string, groupId string, userId string,
  ) (course_group_users.GroupMembership, error)
}

/**
 * The [_BbRestCourseGroupUsers] type...
 */
type _BbRestCourseGroupUsers struct {
  _BlackboardRest

  CourseGroupUsers
}

/**
 * The [GetCourseGroupUsersInstance] function...
 */
func GetCourseGroupUsersInstance (
  host string, accessToken oauth2.AccessToken,
) CourseGroupUsers {
  hostUri, _ := url.Parse (host)

  groupUsersService := new (_BbRestCourseGroupUsers)

  groupUsersService.host = *hostUri
  groupUsersService.accessToken = accessToken

  groupUsersService.service.SetHost (host)
  groupUsersService.service.SetAccessToken (accessToken)

  return groupUsersService
}

/**
 * The [GetGroupMemberships] method...
 */
func (restGroupUsers *_BbRestCourseGroupUsers) GetGroupMemberships (
  courseId string, groupId string,
) ([]course_group_users.GroupMembership, error) {
  endpoint := config.CourseGroupUsersEndpoints["group_memberships"]
  endpoint = strings.Replace (endpoint, "{courseId}", courseId, -1)
  endpoint = strings.Replace (endpoint, "{groupId}", groupId, -1)

  result, err := restGroupUsers.service.Connector.SendBbRequest (
    endpoint, "GET", make (map[string]interface{}), 2,
  )

  if nil != err {
    return []course_group_users.GroupMembership{}, err
  }

  rawGroupUsers := result.(map[string]interface{})["results"]

  groupUsers := course_group_users.NewCourseGroupUsers(
    _scaffolding.NormalizeRawResponse (rawGroupUsers.([]interface{})),
  )

  return groupUsers, err
}
