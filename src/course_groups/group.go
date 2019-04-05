package course_groups

/**
 * The [Group] type...
 */
type Group struct {
  Id, ExternalId, ParentId, Name, Description string

  Availability GroupAvailability

  Enrollment GroupEnrollment
}

/**
 * The [GroupAvailability] type...
 */
type GroupAvailability struct {
  Available GroupAvailable
}

/**
 * The [GroupEnrollment] type...
 */
type GroupEnrollment struct {
  Type GroupEnrollmentType

  Limit int

  SignupSheet SignupSheet
}

/**
 * The [GroupMembership] type...
 */
type GroupMembership struct {
  UserId string
}

/**
 * The [SignupSheet] type...
 */
type SignupSheet struct {
  Name, Description string

  ShowMembers bool
}

/**
 * The [GroupAvailable] type...
 */
type GroupAvailable string

const (
  Yes         GroupAvailable = "Yes"
  No          GroupAvailable = "No"
  SignupOnly  GroupAvailable = "SignupOnly"
)

/**
 * The [GroupEnrollmentType] type...
 */
type GroupEnrollmentType string

const (
  InstructorOnly  GroupEnrollmentType = "InstructorOnly"
  SelfEnroll      GroupEnrollmentType = "SelfEnroll"
)
