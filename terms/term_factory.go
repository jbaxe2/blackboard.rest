package terms

import (
  "time"
)

/**
 * The [NewTerms] function creates new term instances from a slice of raw term
 * information, typically provided as part of a response to the terms service.
 */
func NewTerms (rawTerms []map[string]interface{}) []Term {
  theTerms := make ([]Term, len (rawTerms))

  for i, rawTerm := range rawTerms {
    theTerms[i] = NewTerm (rawTerm)
  }

  return theTerms
}

/**
 * The [NewTerm] function creates a new term instance from raw term information,
 * typically provided as part of a response to the REST API terms service.
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
 * The [_parseTermAvailability] function parses the term availability as per the
 * REST API's documentation specifying appropriate values.
 */
func _parseTermAvailability (rawAvailability interface{}) TermAvailability {
  mappedAvailability := rawAvailability.(map[string]interface{})

  return TermAvailability {
    Available: Availability (mappedAvailability["available"].(string)),
    Duration: _parseDuration (mappedAvailability["duration"]),
  }
}

/**
 * The [_parseDuration] function parses the duration of the term as per the REST
 * API's documentation specifying appropriate values.
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
    if daysOfUse, isInt := mappedDuration["daysOfUse"].(int); isInt {
      termDuration.DaysOfUse = daysOfUse
    } else {
      termDuration.DaysOfUse = int (mappedDuration["daysOfUse"].(float64))
    }
  }

  return termDuration
}
