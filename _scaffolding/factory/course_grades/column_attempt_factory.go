package factory

import (
  "time"

  "github.com/jbaxe2/blackboard.rest/course_grades"
)

/**
 * The [NewColumnAttempts] function...
 */
func NewColumnAttempts (
  rawColumnAttempts []map[string]interface{},
) []course_grades.Attempt {
  attempts := make ([]course_grades.Attempt, len (rawColumnAttempts))

  for i, rawAttempt := range rawColumnAttempts {
    attempts[i] = NewColumnAttempt (rawAttempt)
  }

  return attempts
}

/**
 * The [NewColumnAttempt] function...
 */
func NewColumnAttempt (
  rawColumnAttempt map[string]interface{},
) course_grades.Attempt {
  created, _ := time.Parse (time.RFC3339, rawColumnAttempt["created"].(string))

  if nil == rawColumnAttempt["text"] {
    rawColumnAttempt["text"] = ""
  }

  if nil == rawColumnAttempt["groupAttemptId"] {
    rawColumnAttempt["groupAttemptId"] = ""
  }

  return course_grades.Attempt {
    Id: rawColumnAttempt["id"].(string),
    UserId: rawColumnAttempt["userId"].(string),
    GroupAttemptId: rawColumnAttempt["groupAttemptId"].(string),
    Status: course_grades.AttemptStatus (rawColumnAttempt["status"].(string)),
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
func _parseDisplayGrade (
  rawDisplayGrade interface{}, // map[string]interface{},
) course_grades.DisplayGrade {
  displayGrade := new (course_grades.DisplayGrade)

  semiRawDisplayGrade, haveDisplayGrade := rawDisplayGrade.(map[string]interface{})

  if !haveDisplayGrade {
    return *displayGrade
  }

  displayGrade.ScaleType =
    course_grades.ScaleType (semiRawDisplayGrade["scaleType"].(string))
  displayGrade.Score = semiRawDisplayGrade["score"].(float64)

  if nil == semiRawDisplayGrade["text"] {
    semiRawDisplayGrade["text"] = ""
  }

  displayGrade.Text = semiRawDisplayGrade["text"].(string)

  return *displayGrade
}
