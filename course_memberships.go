package blackboard_rest

import (
  "strings"

  "github.com/jbaxe2/blackboard.rest/api"
  "github.com/jbaxe2/blackboard.rest/course_memberships"
  "github.com/jbaxe2/blackboard.rest/utils"
)

/**
 * The [CourseMemberships] interface provides the base interface for interacting
 * with the REST API's course memberships service.
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
) ([]course_memberships.Membership, error) {
  memberships.service.SetRequestOption ("expand", "user")

  endpoint :=
    strings.Replace (string (api.CourseMemberships), "{courseId}", courseId, 1)

  rawMemberships, err := memberships.service.Request (endpoint, "GET", nil, 1)

  if nil != err {
    return nil, err
  }

  return course_memberships.NewMemberships (
    utils.NormalizeRawResponse (rawMemberships["results"].([]interface{})),
  ), nil
}
