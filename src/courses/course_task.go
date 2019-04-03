package courses

import "time"

/**
 * The [CourseTask] type...
 */
type CourseTask struct {
  Id string

  Status CourseTaskStatus

  PercentComplete int

  Started time.Time
}

/**
 * The [CourseTaskStatus] type...
 */
type CourseTaskStatus string

const (
  Waiting             CourseTaskStatus = "Waiting"
  Assigned            CourseTaskStatus = "Assigned"
  Running             CourseTaskStatus = "Running"
  Complete            CourseTaskStatus = "Complete"
  CompleteWithErrors  CourseTaskStatus = "CompleteWithErrors"
  Incomplete          CourseTaskStatus = "Incomplete"
)
