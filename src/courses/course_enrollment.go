package courses

import "time"

/**
 * The [CourseEnrollment] type...
 */
type CourseEnrollment struct {
  EnrollmentType CourseEnrollmentType

  Start, End time.Time

  AccessCode string
}

/**
 * The [CourseEnrollmentType] type...
 */
type CourseEnrollmentType string

const (
  InstructorLed   CourseEnrollmentType = "InstructorLed"
  SelfEnrollment  CourseEnrollmentType = "SelfEnrollment"
  EmailEnrollment CourseEnrollmentType = "EmailEnrollment"
)
