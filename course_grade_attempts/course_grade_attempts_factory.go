package course_grade_attempts

import (
  "net/url"
)

/**
 * The [NewAttemptFiles] function...
 */
func NewAttemptFiles (rawAttemptFiles []map[string]interface{}) []AttemptFile {
  attemptFiles := make ([]AttemptFile, len (rawAttemptFiles))

  for i, rawAttemptFile := range rawAttemptFiles {
    attemptFiles[i] = NewAttemptFile (rawAttemptFile)
  }

  return attemptFiles
}

/**
 * The [NewAttemptFile] function...
 */
func NewAttemptFile (rawAttemptFile map[string]interface{}) AttemptFile {
  name, _ := rawAttemptFile["name"].(string)
  viewUrl, _ := url.Parse (rawAttemptFile["viewUrl"].(string))

  return AttemptFile {
    Id: rawAttemptFile["id"].(string),
    Name: name,
    ViewUrl: *viewUrl,
  }
}
