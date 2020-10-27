package blackboard_rest

import (
  "net/url"
  "strings"

  "github.com/jbaxe2/blackboard.rest.go/src/_scaffolding"
  "github.com/jbaxe2/blackboard.rest.go/src/_scaffolding/config"
  "github.com/jbaxe2/blackboard.rest.go/src/_scaffolding/factory"
  "github.com/jbaxe2/blackboard.rest.go/src/course_groups"
  "github.com/jbaxe2/blackboard.rest.go/src/oauth2"
)

/**
 * The [CourseGroups] interface...
 */
type CourseGroups interface {
  GetGroups (courseId string) ([]course_groups.Group, error)

  CreateGroup (courseId string, group course_groups.Group) error

  GetGroupSets (courseId string) ([]course_groups.Group, error)

  CreateGroupSet (courseId string, groupSet course_groups.Group) error

  GetGroupSet (courseId string, groupSetId string)  (course_groups.Group, error)

  GetGroupSetGroups (
    courseId string, groupSetId string,
  ) ([]course_groups.Group, error)

  CreateGroupInSet (
    courseId string, groupSetId string, group course_groups.Group,
  ) error

  GetGroup (courseId string, groupSetId string) (course_groups.Group, error)
}

/**
 * The [_BbRestCourseGroups] type...
 */
type _BbRestCourseGroups struct {
  _BlackboardRest

  CourseGroups
}

/**
 * The [GetCourseGroupsInstance] function...
 */
func GetCourseGroupsInstance (
  host string, accessToken oauth2.AccessToken,
) CourseGroups {
  hostUri, _ := url.Parse (host)

  courseGroupsService := new (_BbRestCourseGroups)

  courseGroupsService.host = *hostUri
  courseGroupsService.accessToken = accessToken

  courseGroupsService.service.SetHost (host)
  courseGroupsService.service.SetAccessToken (accessToken)

  return courseGroupsService
}

/**
 * The [GetGroups] method...
 */
func (restCourseGroups *_BbRestCourseGroups) GetGroups (
  courseId string,
) ([]course_groups.Group, error) {
  endpoint := config.CourseGroupsEndpoints["groups"]
  endpoint = strings.Replace (endpoint, "{courseId}", courseId, -1)

  result, err := restCourseGroups.service.Connector.SendBbRequest (
    endpoint, "GET", make (map[string]interface{}), 2,
  )

  if nil != err {
    return []course_groups.Group{}, err
  }

  rawCourseGroups := result.(map[string]interface{})["results"]

  courseGroups := factory.NewCourseGroups (
    _scaffolding.NormalizeRawResponse (rawCourseGroups.([]interface{})),
  )

  return courseGroups, err
}
