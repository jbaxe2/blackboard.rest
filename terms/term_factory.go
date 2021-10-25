package terms

import (
  "time"
)

/**
 * The [NewTerms] function...
 */
func NewTerms (rawTerms []map[string]interface{}) []Term {
  theTerms := make ([]Term, len (rawTerms))

  for i, rawTerm := range rawTerms {
    theTerms[i] = NewTerm (rawTerm)
  }

  return theTerms
}

/**
 * The [NewTerm] function...
 */
func NewTerm (rawTerm map[string]interface{}) Term {
  return Term {
    Id: rawTerm["id"].(string),
    ExternalId: rawTerm["externalId"].(string),
    Name: rawTerm["name"].(string),
    Availability: _parseTermAvailability (rawTerm["availability"]),
  }
}

/**
 * The [_parseTermAvailability] function...
 */
func _parseTermAvailability (rawAvailability interface{}) TermAvailability {
  mappedAvailability := rawAvailability.(map[string]interface{})

  return TermAvailability {
    Available: Availability (mappedAvailability["available"].(string)),
    Duration: _parseDuration (mappedAvailability["duration"]),
  }
}

/**
 * The [_parseDuration] function...
 */
func _parseDuration (rawDuration interface{}) TermDuration {
  var start, end time.Time

  mappedDuration := rawDuration.(map[string]interface{})

  if nil != mappedDuration["start"] {
    start, _ = time.Parse (time.RFC3339, mappedDuration["start"].(string))
  }

  if nil != mappedDuration["end"] {
    end, _ = time.Parse (time.RFC3339, mappedDuration["end"].(string))
  }

  termDuration := TermDuration {
    Type: DurationType(mappedDuration["type"].(string)),
    Start: start,
    End: end,
  }

  if nil != mappedDuration["daysOfUse"] {
    termDuration.DaysOfUse = mappedDuration["daysOfUse"].(int)
  }

  return termDuration
}
