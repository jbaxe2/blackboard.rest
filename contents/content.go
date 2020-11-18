package contents

import "time"

/**
 * The [Content] type...
 */
type Content struct {
  Id, ParentId, Title, Body, Description string

  Created time.Time

  Position int

  HasChildren, HasGradebookColumn, HasAssociatedGroups bool

  Availability ContentAvailability

  ContentHandler ContentHandler
}

/**
 * The [ContentAvailability] type...
 */
type ContentAvailability struct {
  Available ContentAvailable

  AllowGuests bool

  AdaptiveRelease AdaptiveRelease
}

/**
 * The [ContentHandler] type...
 */
type ContentHandler struct {
  Id string
}

/**
 * The [AdaptiveRelease] type...
 */
type AdaptiveRelease struct {
  Start, End time.Time
}

/**
 * The [ContentAvailable] type...
 */
type ContentAvailable string

const (
  Yes              ContentAvailable = "Yes"
  No               ContentAvailable = "No"
  PartiallyVisible ContentAvailable = "PartiallyVisible"
)
