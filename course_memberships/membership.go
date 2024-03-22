package course_memberships

import (
  "time"

  "github.com/jbaxe2/blackboard.rest/courses"
  "github.com/jbaxe2/blackboard.rest/users"
)

/**
 * The [Membership] type...
 */
type Membership struct {
  Id, UserId, CourseId, ChildCourseId, DataSourceId string

  User users.User
  Course courses.Course

  Created, LastAccess, BypassCourseAvailabilityUntil time.Time

  Availability MembershipAvailability
  CourseRoleId MembershipRole
}

/**
 * The [MembershipAvailability] type...
 */
type MembershipAvailability string

const (
  Yes      MembershipAvailability = "Yes"
  No       MembershipAvailability = "No"
  Disabled MembershipAvailability = "Disabled"
)

/**
 * The [MembershipRole] type...
 */
type MembershipRole string

const (
  Instructor        MembershipRole = "Instructor"
  TeachingAssistant MembershipRole = "TeachingAssistant"
  CourseBuilder     MembershipRole = "CourseBuilder"
  Grader            MembershipRole = "Grader"
  Student           MembershipRole = "Student"
  Guest             MembershipRole = "Guest"
)
