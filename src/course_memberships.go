package blackboard_rest

import (
  "github.com/jbaxe2/blackboard.rest.go/src/_scaffolding/config"
  "github.com/jbaxe2/blackboard.rest.go/src/_scaffolding/factory"
  "github.com/jbaxe2/blackboard.rest.go/src/_scaffolding/services"
  "github.com/jbaxe2/blackboard.rest.go/src/course_memberships"
  "github.com/jbaxe2/blackboard.rest.go/src/oauth2"
  "net/url"
  "strings"
)

/**
 * The [CourseMemberships] interface...
 */
type CourseMemberships interface {
  GetMembershipsForCourse (
    courseId string,
  ) ([]course_memberships.Membership, error)

  GetMembershipsForUser (userId string) ([]course_memberships.Membership, error)

  GetMembership (
    courseId string, userId string,
  ) (course_memberships.Membership, error)

  UpdateMembership (
    courseId string, userId string, membership course_memberships.Membership,
  ) error

  CreateMembership (
    courseId string, userId string, membership course_memberships.Membership,
  ) error
}

/**
 * The [_BbRestCourseMemberships] type...
 */
type _BbRestCourseMemberships struct {
  host url.URL

  accessToken oauth2.AccessToken

  service services.BlackboardRestService

  CourseMemberships
}

/**
 * The [GetCourseMembershipsInstance] function...
 */
func GetCourseMembershipsInstance (
  host string, accessToken oauth2.AccessToken,
) CourseMemberships {
  hostUri, _ := url.Parse (host)

  membershipsService := &_BbRestCourseMemberships {
    host: *hostUri, accessToken: accessToken,
  }

  membershipsService.service.SetHost (host)
  membershipsService.service.SetAccessToken (accessToken)

  return membershipsService
}

/**
 * The [GetMembershipsForCourse] method...
 */
func (restMemberships *_BbRestCourseMemberships) GetMembershipsForCourse (
  courseId string,
) ([]course_memberships.Membership, error) {
  endpoint := config.CourseMembershipsEndpoints["course_memberships"]
  endpoint = strings.Replace (endpoint, "{courseId}", courseId, -1)

  return restMemberships._getMemberships (endpoint, make (map[string]interface{}))
}

/**
 * The [GetMembershipsForUser] method...
 */
func (restMemberships *_BbRestCourseMemberships) GetMembershipsForUser (
  userId string,
) ([]course_memberships.Membership, error) {
  endpoint := config.CourseMembershipsEndpoints["user_memberships"]
  endpoint = strings.Replace (endpoint, "{userId}", userId, -1)

  return restMemberships._getMemberships (endpoint, make (map[string]interface{}))
}

/**
 * The [GetMembership] method...
 */
func (restMemberships *_BbRestCourseMemberships) GetMembership (
  courseId string, userId string,
) (course_memberships.Membership, error) {
  var courseMembership course_memberships.Membership
  var err error
  var result interface{}

  endpoint := config.CourseMembershipsEndpoints["membership"]
  endpoint = strings.Replace (endpoint, "{courseId}", courseId, -1)
  endpoint = strings.Replace (endpoint, "{userId}", userId, -1)

  result, err = restMemberships.service.Connector.SendBbRequest (
    endpoint, "GET", make (map[string]interface{}), 1,
  )

  if nil != err {
    return courseMembership, err
  }

  courseMembership = factory.NewMembership (result.(map[string]interface{}))

  return courseMembership, err
}

/**
 * The [_getMemberships] method...
 */
func (restMemberships *_BbRestCourseMemberships) _getMemberships (
  endpoint string, data map[string]interface{},
) ([]course_memberships.Membership, error) {
  var courseMemberships []course_memberships.Membership
  var err error
  var result interface{}

  result, err = restMemberships.service.Connector.SendBbRequest (
    endpoint, "GET", data, 1,
  )

  if nil != err {
    return courseMemberships, err
  }

  rawMemberships := result.(map[string]interface{})["results"]

  courseMemberships = factory.NewMemberships (
    _normalizeRawMemberships (rawMemberships.([]interface{})),
  )

  return courseMemberships, err
}

/**
 * The [_normalizeRawMemberships] function...
 */
func _normalizeRawMemberships (
  rawMemberships []interface{},
) []map[string]interface{} {
  mappedRawMemberships := make ([]map[string]interface{}, len (rawMemberships))

  for i, rawMembership := range rawMemberships {
    mappedRawMemberships[i] = rawMembership.(map[string]interface{})
  }

  return mappedRawMemberships
}
