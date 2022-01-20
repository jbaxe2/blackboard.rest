package blackboard_rest

import (
  "github.com/jbaxe2/blackboard.rest/api"
  "github.com/jbaxe2/blackboard.rest/course_memberships"
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
