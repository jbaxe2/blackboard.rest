package course_grades

/**
 * The [GradeSchema] type...
 */
type GradeSchema struct {
  Id, ExternalId, Title, Description string

  ScaleType ScaleType

  Symbols []GradeSymbol
}

/**
 * The [GradeSymbol] type...
 */
type GradeSymbol struct {
  Text string

  AbsoluteValue, LowerBound, UpperBound float64
}
