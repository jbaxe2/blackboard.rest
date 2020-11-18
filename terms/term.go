package terms

import "time"

/**
 * The [Term] type...
 */
type Term struct {
  Id, ExternalId, DataSourceId, Name, Description string

  Availability TermAvailability
}

/**
 * The [TermAvailability] type...
 */
type TermAvailability struct {
  Available Availability

  Duration TermDuration
}

/**
 * The [TermDuration] type...
 */
type TermDuration struct {
  Type DurationType

  Start, End time.Time

  DaysOfUse int
}

/**
 * The [Availability] type...
 */
type Availability string

const (
  Yes Availability = "Yes"
  No  Availability = "No"
)

/**
 * The [DurationType] type...
 */
type DurationType string

const (
  Continuous   DurationType = "Continuous"
  DateRange    DurationType = "DateRange"
  FixedNumDays DurationType = "FixedNumDays"
)
