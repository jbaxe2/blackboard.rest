package course_memberships

import "time"

/**
 * The [Membership] type...
 */
type Membership struct {
  UserId, CourseId, ChildCourseId, DataSourceId string

  Created, LastAccess, BypassCourseAvailabilityUntil time.Time

  Availability MembershipAvailability

  CourseRoleId MembershipRole
}
