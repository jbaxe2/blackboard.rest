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

/**
 * The [AsMap] method returns a map containing the information contained in some
 * content instance as a raw string-based map.
 */
func (content *Content) AsMap() map[string]interface{} {
  return map[string]interface{} {
    "id": content.Id,
    "parentId": content.ParentId,
    "title": content.Title,
    "body": content.Body,
    "description": content.Description,
    "created": content.Created.Format (time.RFC3339),
    "modified": content.Modified.Format (time.RFC3339),
    "position": content.Position,
    "hasChildren": content.HasChildren,
    "hasGradebookColumn": content.HasGradebookColumn,
    "hasAssociatedGroups": content.HasAssociatedGroups,
    "launchInNewWindow": content.LaunchInNewWindow,
    "reviewable": content.Reviewable,
    "availability": map[string]interface{} {
      "available": string (content.Availability.Available),
      "allowGuests": content.Availability.AllowGuests,
      "adaptiveRelease": map[string]interface{} {
        "start": content.Availability.AdaptiveRelease.Start.Format (time.RFC3339),
        "end": content.Availability.AdaptiveRelease.End.Format (time.RFC3339),
      },
    },
    "contentHandler": content.ContentHandler.AsMap(),
  }
}

/**
 * The [AsNewContentMap] method returns a map of the content containing only the
 * information accepted by a Learn server for creating the content.
 */
func (content *Content) AsNewContentMap() map[string]interface{} {
  return map[string]interface{} {
    "parentId": content.ParentId,
    "title": content.Title,
    "body": content.Body,
    "description": content.Description,
    "position": content.Position,
    "launchInNewWindow": content.LaunchInNewWindow,
    "reviewable": content.Reviewable,
    "availability": map[string]interface{}{
      "available":  "Yes",
    },
    "contentHandler": content.ContentHandler.AsMap(),
  }
}
