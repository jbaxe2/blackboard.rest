package courses

import "time"

/**
 * The [CourseAvailability] type...
 */
type CourseAvailability struct {
  Availability string

  Duration CourseDuration
}

/**
 * The [CourseDuration] type...
 */
type CourseDuration struct {
  DurationType CourseDurationType

  Start, End time.Time

  DaysOfUse int
}

/**
 * The [CourseDurationType] type...
 */
type CourseDurationType string

const (
  Continuous   CourseDurationType = "Continuous"
  DateRange    CourseDurationType = "DateRange"
  FixedNumDays CourseDurationType = "FixedNumDays"
  Term         CourseDurationType = "Term"
)
