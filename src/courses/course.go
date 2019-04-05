package courses

import "time"

/**
 * The [Course] type...
 */
type Course struct {
  Id, Uuid, ExternalId, DataSourceId, CourseId, Name, Description, TermId,
  ParentId, ExternalAccessUrl, GuestAccessUrl string

  Organization, AllowGuests, ReadOnly, HasChildren bool

  Created time.Time

  UltraStatus UltraStatus

  Availability CourseAvailability

  Enrollment CourseEnrollment

  Locale Locale
}

/**
 * The [CourseChild] type...
 */
type CourseChild struct {
  Id, ParentId, DataSourceId string

  Created time.Time
}

/**
 * The [Locale] type...
 */
type Locale struct {
  Id string

  Force bool
}

/**
 * The [UltraStatus] type...
 */
type UltraStatus string

const (
  Undecided     UltraStatus = "Undecided"
  Classic       UltraStatus = "Classic"
  Ultra         UltraStatus = "Ultra"
  UltraPreview  UltraStatus = "UltraPreview"
)
