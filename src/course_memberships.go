package blackboard_rest

import (
  "github.com/jbaxe2/blackboard.rest.go/src/_scaffolding/config"
  error2 "github.com/jbaxe2/blackboard.rest.go/src/_scaffolding/error"
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

  service services.BlackboardRestServices

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

  membershipsService.service.SetAccessToken (accessToken)

  return membershipsService
}

/**
 * The [GetMembershipsForCourse] method...
 */
func (restMemberships *_BbRestCourseMemberships) GetMembershipsForCourse (
  courseId string,
) ([]course_memberships.Membership, error) {
  endpoint := config.CourseMembershipsEndpoints()["course_memberships"]
  endpoint = strings.Replace (endpoint, "{courseId}", courseId, -1)

  return restMemberships._getMemberships (endpoint, make(map[string]interface{}))
}

/**
 * The [GetMembershipsForUser] method...
 */
func (restMemberships *_BbRestCourseMemberships) GetMembershipsForUser (
  userId string,
) ([]course_memberships.Membership, error) {
  endpoint := config.CourseMembershipsEndpoints()["user_memberships"]
  endpoint = strings.Replace (endpoint, "{userId}", userId, -1)

  return restMemberships._getMemberships (endpoint, make(map[string]interface{}))
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

  if (nil != err) && (error2.RestError{} != err) {
    err = restMemberships.HandleError (err.(error2.RestError))

    return courseMemberships, err
  }

  rawMemberships := result.(map[string]interface{})["results"]

  courseMemberships = factory.NewMemberships (
    rawMemberships.([]map[string]interface{}),
  )

  return courseMemberships, err
}

/**
 * The [HandleError] method...
 */
func (restMemberships *_BbRestCourseMemberships) HandleError (
  err error2.RestError,
) error2.CourseMembershipsError {
  membershipsErr := error2.CourseMembershipsError{}

  membershipsErr.SetStatus (err.Status())
  membershipsErr.SetCode (err.Code())
  membershipsErr.SetMessage (err.Message())
  membershipsErr.SetDeveloperMessage (err.DeveloperMessage())
  membershipsErr.SetExtraInfo (err.ExtraInfo())

  return membershipsErr
}
