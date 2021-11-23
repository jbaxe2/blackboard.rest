package contents

import (
  "net/url"
  "time"
)

/**
 * The [Content] type...
 */
type Content struct {
  Id, ParentId, Title, Body, Description string

  Created, Modified time.Time

  Position int

  HasChildren, HasGradebookColumn, HasAssociatedGroups, LaunchInNewWindow,
    Reviewable bool

  Availability Availability

  ContentHandler ContentHandler

  Links []Link
}

/**
 * The [Availability] type...
 */
type Availability struct {
  Available Available

  AllowGuests bool

  AdaptiveRelease AdaptiveRelease
}

/**
 * The [AdaptiveRelease] type...
 */
type AdaptiveRelease struct {
  Start, End time.Time
}

/**
 * The [Link] type...
 */
type Link struct {
  Href *url.URL

  Rel, Title, Type string
}

/**
 * The [Available] type...
 */
type Available string

const (
  Yes              Available = "Yes"
  No               Available = "No"
  PartiallyVisible Available = "PartiallyVisible"
)
