package contents

import (
  "time"

  "github.com/jbaxe2/blackboard.rest/utils"
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
  if nil == rawContent {
    return Content{}
  }

  id, _ := rawContent["id"].(string)
  parentId, _ := rawContent["parentId"].(string)
  body, _ := rawContent["body"].(string)
  title, _ := rawContent["title"].(string)
  rawCreated, _ := rawContent["created"].(string)
  rawModified, _ := rawContent["modified"].(string)
  launchInNewWindow, _ := rawContent["launchInNewWindow"].(bool)
  reviewable, _ := rawContent["reviewable"].(bool)

  created, _ := time.Parse (time.RFC3339, rawCreated)
  modified, _ := time.Parse (time.RFC3339, rawModified)

  hasChildren, _ := rawContent["hasChildren"].(bool)
  hasGradebookColumns, _ := rawContent["hasGradebookColumns"].(bool)
  hasAssociatedGroups, _ := rawContent["hasAssociatedGroups"].(bool)

  contentHandler, _ := rawContent["contentHandler"].(map[string]interface{})

  return Content {
    Id: id,
    ParentId: parentId,
    Title: title,
    Body: body,
    Created: created,
    Modified: modified,
    Position: _parsePosition (rawContent["position"]),
    HasChildren: hasChildren,
    HasGradebookColumn: hasGradebookColumns,
    HasAssociatedGroups: hasAssociatedGroups,
    LaunchInNewWindow: launchInNewWindow,
    Reviewable: reviewable,
    ContentHandler: NewContentHandler (contentHandler),
    Availability:
      _parseAvailability (rawContent["availability"].(map[string]interface{})),
  }
}

func _parsePosition (rawPosition interface{}) int {
  return int (utils.NormalizeNumeric (rawPosition))
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
