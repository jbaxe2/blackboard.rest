package users

import (
  "time"
)

/**
 * The [User] type...
 */
type User struct {
  Id, Uuid, ExternalId, DataSourceId, UserName, Email, StudentId string

  Created, LastLogin time.Time

  InstitutionRoleIds []string

  SystemRoleIds []SystemRole

  Availability UserAvailability

  Name Name
}

/**
 * The [Name] type...
 */
type Name struct {
  Given, Family, Middle, Other, Suffix, Title string
}

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

/**
 * The [UserAvailability] type...
 */
type UserAvailability string

const (
  Yes       UserAvailability = "Yes"
  No        UserAvailability = "No"
  Disabled  UserAvailability = "Disabled"
)
