package terms_test

import (
  "testing"

  "github.com/jbaxe2/blackboard.rest/terms"
)

/**
 * The [TestCreateNewTerm] function...
 */
func TestCreateNewTerm (t *testing.T) {
  println ("Create a new term instance.")

  newTerm := terms.NewTerm (rawTerm)

  if !(newTerm.Id == rawTerm["id"] &&
       string (newTerm.Availability.Available) == "Yes") {
    t.Error ("Creating a new term instance should have expected value.")
  }
}

/**
 * The [TestCreateNewTerms] function...
 */
func TestCreateNewTerms (t *testing.T) {
  println ("Create multiple new term instances.")

  newTerms := terms.NewTerms (rawTerms)

  if !(2 == len (newTerms) && newTerms[0].Id == rawTerm["id"]) {
    t.Error ("Creating new term instances should have expected values.")
  }
}

/**
 * Mocked instances to run the above tests with.
 */
var rawTerms = []map[string]interface{} {rawTerm, rawTerm2}

var rawTerm = map[string]interface{} {
  "id": "termId",
  "externalId": "externalTermId",
  "dataSourceId": "term.data.source.id",
  "name": "Term Name",
  "description": "",
  "availability": map[string]interface{} {
    "available": "Yes",
    "duration": map[string]interface{} {
      "type": "Continuous",
      "start": "2021-11-16T20:38:45.738Z",
      "end": "2021-11-16T20:38:45.738Z",
      "daysOfUse": 0,
    },
  },
}

var rawTerm2 = map[string]interface{} {
  "id": "termId2",
  "externalId": "externalTermId2",
  "dataSourceId": "term.data.source.id",
  "name": "Term Name 2",
  "description": "",
  "availability": map[string]interface{} {
    "available": "Yes",
    "duration": map[string]interface{} {
      "type": "Continuous",
      "start": "2021-11-16T20:38:45.738Z",
      "end": "2021-11-16T20:38:45.738Z",
      "daysOfUse": 0,
    },
  },
}
