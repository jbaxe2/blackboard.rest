package course_grades

import (
  "time"

  "github.com/jbaxe2/blackboard.rest/utils"
)

/**
 * The [NewColumnAttempts] function creates a slice of Column Attempt instances
 * from a slice of raw maps.
 */
func NewColumnAttempts (rawColumnAttempts []map[string]interface{}) []Attempt {
  attempts := make ([]Attempt, len (rawColumnAttempts))

  for i, rawAttempt := range rawColumnAttempts {
    attempts[i] = NewColumnAttempt (rawAttempt)
  }

  return attempts
}

/**
 * The [NewColumnAttempt] function creates a new Column Attempt instance from a
 * raw map.
 */
func NewColumnAttempt (rawColumnAttempt map[string]interface{}) Attempt {
  created, _ := time.Parse (time.RFC3339, rawColumnAttempt["created"].(string))

  if nil == rawColumnAttempt["text"] {
    rawColumnAttempt["text"] = ""
  }

  if nil == rawColumnAttempt["groupAttemptId"] {
    rawColumnAttempt["groupAttemptId"] = ""
  }

  return Attempt {
    Id: rawColumnAttempt["id"].(string),
    UserId: rawColumnAttempt["userId"].(string),
    GroupAttemptId: rawColumnAttempt["groupAttemptId"].(string),
    Status: AttemptStatus(rawColumnAttempt["status"].(string)),
    DisplayGrade:
      _parseDisplayGrade (rawColumnAttempt["displayGrade"]),
    Score: utils.NormalizeNumeric (rawColumnAttempt["score"]),
    Exempt: rawColumnAttempt["exempt"].(bool),
    Created: created,
    Text: rawColumnAttempt["text"].(string),
  }
}

/**
 * The [_parseDisplayGrade] function parses the display grade for a column attempt.
 */
func _parseDisplayGrade (rawDisplayGrade interface{}) DisplayGrade {
  displayGrade := new (DisplayGrade)
  semiRawDisplayGrade, haveDisplayGrade := rawDisplayGrade.(map[string]interface{})

  if !haveDisplayGrade {
    return *displayGrade
  }

  displayGrade.ScaleType =
    ScaleType (semiRawDisplayGrade["scaleType"].(string))
  displayGrade.Score = utils.NormalizeNumeric (semiRawDisplayGrade["score"])

  if nil == semiRawDisplayGrade["text"] {
    semiRawDisplayGrade["text"] = ""
  }

  displayGrade.Text = semiRawDisplayGrade["text"].(string)

  return *displayGrade
}
