package course_grades

import "time"

/**
 * The [GradeColumn] type...
 */
type GradeColumn struct {
  Id, ExternalId, Name, DisplayName, Description, ContentId string

  ExternalGrade bool

  Created time.Time

  Score Scoring

  Availability GradeAvailability

  Grading Grading
}

/**
 * The [Score] type...
 */
type Scoring struct {
  Possible float64
}

/**
 * The [GradeAvailability] type...
 */
type GradeAvailability string

const (
  Yes GradeAvailability = "Yes"
  No  GradeAvailability = "No"
)
