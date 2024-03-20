package blackboard_rest

import (
  "strings"

  "github.com/jbaxe2/blackboard.rest/api"
  courseMemberships "github.com/jbaxe2/blackboard.rest/course_memberships"
  "github.com/jbaxe2/blackboard.rest/utils"
)

/**
 * The [CourseMemberships] interface provides the base interface for interacting
 * with the REST API's course memberships service.
 */
type CourseMemberships interface {
  GetMembershipsForCourse (
    courseId string,
  ) ([]courseMemberships.Membership, error)

  GetMembershipsForUser (userId string) ([]courseMemberships.Membership, error)

  GetMembership (
    courseId string, userId string,
  ) (courseMemberships.Membership, error)

  UpdateMembership (
    courseId string, userId string, membership courseMemberships.Membership,
  ) error

  CreateMembership (
    courseId string, userId string, membership courseMemberships.Membership,
  ) error
}

/**
 * The [_CourseMemberships] type implements the Course Memberships interface.
 */
type _CourseMemberships struct {
  service api.Service

  CourseMemberships
}

/**
 * The [NewCourseMemberships] function creates a new course memberships instance.
 */
func NewCourseMemberships (service api.Service) CourseMemberships {
  if nil == service {
    return nil
  }

  return &_CourseMemberships {
    service: service,
  }
}

/**
 * The [GetMembershipsForCourse] method obtains the collection of course
 * memberships, based on the course's ID.
 */
func (memberships *_CourseMemberships) GetMembershipsForCourse (
  courseId string,
) ([]courseMemberships.Membership, error) {
  return memberships._getMemberships (
    string (api.CourseMemberships), "{courseId}", courseId,
  )
}

/**
 * The [GetMembershipsForCourse] method obtains the collection of course
 * memberships, based on the user's ID.
 */
func (memberships *_CourseMemberships) GetMembershipsForUser (
  userId string,
) ([]courseMemberships.Membership, error) {
  return memberships._getMemberships (
    string (api.UserMemberships), "{userId}", userId,
  )
}

/**
 * The [_getMemberships] method obtains a collection of course memberships, for
 * either the course or user based on the provided context type and context ID.
 */
func (memberships *_CourseMemberships) _getMemberships (
  rawEndpoint, contextType, contextId string,
) ([]courseMemberships.Membership, error) {
  memberships.service.SetRequestOption ("expand", "user")

  endpoint := strings.Replace (rawEndpoint, contextType, contextId, 1)
  rawMemberships, err := memberships.service.Request (endpoint, "GET", nil, 1)

  if nil != err {
    return nil, err
  }

  return courseMemberships.NewMemberships (
    utils.NormalizeRawResponse (rawMemberships["results"].([]interface{})),
  ), nil
}
