package course_grades

import "time"

/**
 * The [Grading] type...
 */
type Grading struct {
  Type GradingType

  Due time.Time

  AttemptsAllowed int

  ScoringModel ScoringModel

  SchemaId string

  AnonymousGrading AnonymousGrading
}

/**
 * The [AnonymousGrading] type...
 */
type AnonymousGrading struct {
  Type AnonymousGradingType

  ReleaseAfter time.Time
}

/**
 * The [GradingType] type...
 */
type GradingType string

const (
  Attempts    GradingType = "Attempts"
  Calculated  GradingType = "Calculated"
  Manual      GradingType = "Manual"
)

/**
 * The [ScoringModel] type...
 */
type ScoringModel string

const (
  Last    ScoringModel = "Last"
  Highest ScoringModel = "Highest"
  Lowest  ScoringModel = "Lowest"
  First   ScoringModel = "First"
  Average ScoringModel = "Average"
)

/**
 * The [AnonymousGradingType] type...
 */
type AnonymousGradingType string

const (
  None            AnonymousGradingType = "None"
  AfterAllGraded  AnonymousGradingType = "AfterAllGraded"
  Date            AnonymousGradingType = "Date"
)
