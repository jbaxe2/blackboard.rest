package course_grades

import "time"

/**
 * The [Attempt] type...
 */
type Attempt struct {
  Id, UserId, GroupAttemptId, Text, Notes, Feedback, StudentComments,
  StudentSubmission string

  GroupOverride, Exempt bool

  Status AttemptStatus

  DisplayGrade DisplayGrade

  Score float64

  Created time.Time
}

/**
 * The [AttemptStatus] type...
 */
type AttemptStatus string

const (
  NotAttempted     AttemptStatus = "NotAttempted"
  Abandoned        AttemptStatus = "Abandoned"
  InProgress       AttemptStatus = "InProgress"
  Suspended        AttemptStatus = "Suspended"
  Canceled         AttemptStatus = "Canceled"
  NeedsGrading     AttemptStatus = "NeedsGrading"
  Completed        AttemptStatus = "Completed"
  InMoreProgress   AttemptStatus = "InMoreProgress"
  NeedsMoreGrading AttemptStatus = "NeedsMoreGrading"
)
