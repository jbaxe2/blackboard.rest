package blackboard_rest

import (
  "github.com/jbaxe2/blackboard.rest/course_memberships"
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
