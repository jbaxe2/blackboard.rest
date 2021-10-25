package course_grades

import (
  "time"
)

/**
 * The [NewColumnAttempts] function...
 */
func NewColumnAttempts (rawColumnAttempts []map[string]interface{}) []Attempt {
  attempts := make ([]Attempt, len (rawColumnAttempts))

  for i, rawAttempt := range rawColumnAttempts {
    attempts[i] = NewColumnAttempt (rawAttempt)
  }

  return attempts
}

/**
 * The [NewColumnAttempt] function...
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
    Score: rawColumnAttempt["score"].(float64),
    Exempt: rawColumnAttempt["exempt"].(bool),
    Created: created,
    Text: rawColumnAttempt["text"].(string),
  }
}

/**
 * The [_parseDisplayGrade] function...
 */
func _parseDisplayGrade (rawDisplayGrade interface{}) DisplayGrade {
  displayGrade := new (DisplayGrade)
  semiRawDisplayGrade, haveDisplayGrade := rawDisplayGrade.(map[string]interface{})

  if !haveDisplayGrade {
    return *displayGrade
  }

  displayGrade.ScaleType =
    ScaleType(semiRawDisplayGrade["scaleType"].(string))
  displayGrade.Score = semiRawDisplayGrade["score"].(float64)

  if nil == semiRawDisplayGrade["text"] {
    semiRawDisplayGrade["text"] = ""
  }

  displayGrade.Text = semiRawDisplayGrade["text"].(string)

  return *displayGrade
}
