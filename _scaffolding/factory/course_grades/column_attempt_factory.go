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

  return course_grades.Attempt {
    Id: rawColumnAttempt["id"].(string),
    UserId: rawColumnAttempt["userId"].(string),
    Status: course_grades.AttemptStatus (rawColumnAttempt["status"].(string)),
    DisplayGrade:
      _parseDisplayGrade (rawColumnAttempt["displayGrade"].(map[string]interface{})),
    Text: rawColumnAttempt["text"].(string),
    Score: rawColumnAttempt["score"].(float64),
    Exempt: rawColumnAttempt["exempt"].(bool),
    Created: created,
  }
}

/**
 * The [_parseDisplayGrade] function...
 */
func _parseDisplayGrade (
  rawDisplayGrade map[string]interface{},
) course_grades.DisplayGrade {
  displayGrade := new (course_grades.DisplayGrade)
  displayGrade.ScaleType =
    course_grades.ScaleType (rawDisplayGrade["scaleType"].(string))
  displayGrade.Score = rawDisplayGrade["score"].(float64)

  if nil == rawDisplayGrade["text"] {
    rawDisplayGrade["text"] = ""
  }

  displayGrade.Text = rawDisplayGrade["text"].(string)

  return *displayGrade
}
