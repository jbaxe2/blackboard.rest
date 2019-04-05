package factory

import (
  "github.com/jbaxe2/blackboard.rest.go/src/course_grades"
  "time"
)

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
  return course_grades.DisplayGrade {
    ScaleType: course_grades.ScaleType (rawDisplayGrade["scaleType"].(string)),
    Score: rawDisplayGrade["score"].(float64),
    Text: rawDisplayGrade["text"].(string),
  }
}
