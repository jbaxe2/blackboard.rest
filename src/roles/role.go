package roles

/**
 * The [Role] type...
 */
type Role struct {
  Id, RoleId, Name, Description string

  Custom bool
}

/**
 * The [SystemRole] type...
 */
type SystemRole Role

/**
 * The [InstitutionRole] type...
 */
type InstitutionRole Role

/**
 * The [CourseRole] type...
 */
type CourseRole struct {
  Id, RoleId, NameForCourses, NameForOrganizations, Description string

  ActAsInstructor bool

  Availability CourseRoleAvailability
}

/**
 * The [CourseRoleAvailability] type...
 */
type CourseRoleAvailability string

const (
  Course        CourseRoleAvailability = "Course"
  Organization  CourseRoleAvailability = "Organization"
  Both          CourseRoleAvailability = "Both"
  None          CourseRoleAvailability = "None"
)
