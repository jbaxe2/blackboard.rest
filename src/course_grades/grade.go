package course_grades

import "time"

/**
 * The [Grade] type...
 */
type Grade struct {
  UserId, ColumnId, Text, Notes, Feedback, GradeNotationId string

  GradeStatus GradeStatus

  DisplayGrade DisplayGrade

  Score float64

  Overridden time.Time

  Exempt, Corrupt bool

  ChangeIndex int
}

/**
 * The [DisplayGrade] type...
 */
type DisplayGrade struct {
  ScaleType ScaleType

  Score float64

  Text string
}

/**
 * The [GradeStatus] type...
 */
type GradeStatus string

const (
  Graded            GradeStatus = "Graded"
  GradeNeedsGrading GradeStatus = "Needs Grading"
)

/**
 * The [ScaleType] type...
 */
type ScaleType string

const (
  Percent   ScaleType = "Percent"
  Score     ScaleType = "Score"
  Tabular   ScaleType = "Tabular"
  Text      ScaleType = "Text"
  Complete  ScaleType = "Complete"
)
