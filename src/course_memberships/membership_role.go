package course_memberships

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
