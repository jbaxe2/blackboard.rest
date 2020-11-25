package course_memberships

import (
  "time"

  "github.com/jbaxe2/blackboard.rest/users"
)

/**
 * The [Membership] type...
 */
type Membership struct {
  Id, UserId, CourseId, ChildCourseId, DataSourceId string

  User users.User

  Created, LastAccess, BypassCourseAvailabilityUntil time.Time

  Availability MembershipAvailability

  CourseRoleId MembershipRole
}
