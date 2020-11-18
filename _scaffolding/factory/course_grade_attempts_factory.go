package factory

import (
  "net/url"

  "github.com/jbaxe2/blackboard.rest/course_grade_attempts"
)

/**
 * The [NewAttemptFiles] function...
 */
func NewAttemptFiles (
  rawAttemptFiles []map[string]interface{},
) []course_grade_attempts.AttemptFile {
  attemptFiles := make ([]course_grade_attempts.AttemptFile, len (rawAttemptFiles))

  for i, rawAttemptFile := range rawAttemptFiles {
    attemptFiles[i] = NewAttemptFile (rawAttemptFile)
  }

  return attemptFiles
}

/**
 * The [NewAttemptFile] function...
 */
func NewAttemptFile (
  rawAttemptFile map[string]interface{},
) course_grade_attempts.AttemptFile {
  name, _ := rawAttemptFile["name"].(string)
  viewUrl, _ := url.Parse (rawAttemptFile["viewUrl"].(string))

  return course_grade_attempts.AttemptFile {
    Id: rawAttemptFile["id"].(string),
    Name: name,
    ViewUrl: *viewUrl,
  }
}
