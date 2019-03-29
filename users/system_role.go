package users

/**
 * The [SystemRole] type...
 */
type SystemRole string

const (
  SystemAdmin   SystemRole = "SystemAdmin"
  SystemSupport SystemRole = "SystemSupport"
  CourseCreator SystemRole = "CourseCreator"
  CourseSupport SystemRole = "CourseSupport"
  AccountAdmin  SystemRole = "AccountAdmin"
  Guest         SystemRole = "Guest"
  Observer      SystemRole = "Observer"
  Integration   SystemRole = "Integration"
  Portal        SystemRole = "Portal"
)
