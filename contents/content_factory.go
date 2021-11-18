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
  created, _ := time.Parse (time.RFC3339, rawContent["created"].(string))
  modified, _ := time.Parse (time.RFC3339, rawContent["modified"].(string))

  return Content {
    Id: rawContent["id"].(string),
    ParentId: rawContent["parentId"].(string),
    Title: rawContent["title"].(string),
    Body: rawContent["body"].(string),
    Created: created,
    Modified: modified,
    Position: rawContent["position"].(int),
    HasChildren: rawContent["hasChildren"].(bool),
    HasGradebookColumn: rawContent["hasGradebookColumns"].(bool),
    HasAssociatedGroups: rawContent["hasAssociatedGroups"].(bool),
    LaunchInNewWindow: rawContent["launchInNewWindow"].(bool),
    Reviewable: rawContent["reviewable"].(bool),
    Availability:
      _parseAvailability (rawContent["availability"].(map[string]interface{})),
  }
}

/**
 * The [_parseAvailability] function parses the availability of the content.
 */
func _parseAvailability (rawAvailability map[string]interface{}) Availability {
  return Availability {
    Available: Available (rawAvailability["available"].(string)),
    AllowGuests: rawAvailability["allowGuests"].(bool),
    AdaptiveRelease: _parseAdaptiveRelease (
      rawAvailability["adaptiveRelease"].(map[string]string),
    ),
  }
}

/**
 * The [_parseAdaptiveRelease] function parses the times of the adaptive release
 * rules, if any, for this content.
 */
func _parseAdaptiveRelease (
  rawAdaptiveRelease map[string]string,
) AdaptiveRelease {
  start, _ := time.Parse (time.RFC3339, rawAdaptiveRelease["start"])
  end, _ := time.Parse (time.RFC3339, rawAdaptiveRelease["end"])

  return AdaptiveRelease {
    Start: start,
    End: end,
  }
}
