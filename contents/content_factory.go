package contents

import (
  "time"
)

/**
 * The [NewContents] function creates new content instances from a slice of raw
 * maps containing content information, presumably returned from a REST API call.
 */
func NewContents (rawContents []map[string]interface{}) []Content {
  newContents := make ([]Content, len (rawContents))

  for i, rawContent := range rawContents {
    newContents[i] = NewContent (rawContent)
  }

  return newContents
}

/**
 * The [NewContent] function creates a new content instance from a raw map of
 * content information, presumably returned from a REST API call.
 */
func NewContent (rawContent map[string]interface{}) Content {
  parentId, _ := rawContent["parentId"].(string)
  body, _ := rawContent["body"].(string)

  created, _ := time.Parse (time.RFC3339, rawContent["created"].(string))
  modified, _ := time.Parse (time.RFC3339, rawContent["modified"].(string))

  hasChildren, _ := rawContent["hasChildren"].(bool)
  hasGradebookColumns, _ := rawContent["hasGradebookColumns"].(bool)
  hasAssociatedGroups, _ := rawContent["hasAssociatedGroups"].(bool)

  return Content {
    Id: rawContent["id"].(string),
    ParentId: parentId,
    Title: rawContent["title"].(string),
    Body: body,
    Created: created,
    Modified: modified,
    Position: _parsePosition (rawContent["position"]),
    HasChildren: hasChildren,
    HasGradebookColumn: hasGradebookColumns,
    HasAssociatedGroups: hasAssociatedGroups,
    LaunchInNewWindow: rawContent["launchInNewWindow"].(bool),
    Reviewable: rawContent["reviewable"].(bool),
    Availability:
      _parseAvailability (rawContent["availability"].(map[string]interface{})),
  }
}

func _parsePosition (rawPosition interface{}) int {
  var position = 0

  if intPosition, isInt := rawPosition.(int); isInt {
    position = intPosition
  } else {
    position = int (rawPosition.(float64))
  }

  return position
}

/**
 * The [_parseAvailability] function parses the availability of the content.
 */
func _parseAvailability (rawAvailability map[string]interface{}) Availability {
  return Availability {
    Available: Available (rawAvailability["available"].(string)),
    AllowGuests: rawAvailability["allowGuests"].(bool),
    AdaptiveRelease: _parseAdaptiveRelease (
      rawAvailability["adaptiveRelease"].(map[string]interface{}),
    ),
  }
}

/**
 * The [_parseAdaptiveRelease] function parses the times of the adaptive release
 * rules, if any, for this content.
 */
func _parseAdaptiveRelease (
  rawAdaptiveRelease map[string]interface{},
) AdaptiveRelease {
  var start, end time.Time

  if rawStart, haveStart := rawAdaptiveRelease["start"].(string); haveStart {
    start, _ = time.Parse (time.RFC3339, rawStart)
  }

  if rawEnd, haveEnd := rawAdaptiveRelease["end"].(string); haveEnd {
    end, _ = time.Parse (time.RFC3339, rawEnd)
  }

  return AdaptiveRelease {
    Start: start,
    End: end,
  }
}
