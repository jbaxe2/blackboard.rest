package course_grade_attempts_test

import (
  "testing"

  "github.com/jbaxe2/blackboard.rest/course_grade_attempts"
)

/**
 * The [TestCreateNewCourseGradeAttemptsFile] function...
 */
func TestCreateNewCourseGradeAttemptsFile (t *testing.T) {
  println ("Create a new course grade attempts file instance.")

  gradeAttemptsFile :=
    course_grade_attempts.NewAttemptFile (rawCourseGradeAttemptFile)

  if !(gradeAttemptsFile.Id == rawCourseGradeAttemptFile["id"] &&
       gradeAttemptsFile.Name == rawCourseGradeAttemptFile["name"] &&
       gradeAttemptsFile.ViewUrl.String() == rawCourseGradeAttemptFile["viewUrl"]) {
    t.Error (
      "Creating a new course grade attempts file instance should have expected value.",
    )
  }
}

/**
 * The [TestCreateNewCourseGradeAttemptsFiles] function...
 */
func TestCreateNewCourseGradeAttemptsFiles (t *testing.T) {
  println ("Create multiple course grade attempts file instances.")

  gradeAttemptsFiles :=
    course_grade_attempts.NewAttemptFiles (rawCourseGradeAttemptFiles)

  if !(gradeAttemptsFiles[0].Id == rawCourseGradeAttemptFile["id"] &&
       gradeAttemptsFiles[1].Id == rawCourseGradeAttemptFile2["id"] &&
       gradeAttemptsFiles[2].Id == rawCourseGradeAttemptFile3["id"]) {
    t.Error (
      "Creating multiple course grade attempts files should have expected values.",
    )
  }
}

/**
 * Mocked instances to run the above tests with.
 */
var rawCourseGradeAttemptFiles = []map[string]interface{} {
  rawCourseGradeAttemptFile, rawCourseGradeAttemptFile2, rawCourseGradeAttemptFile3,
}

var rawCourseGradeAttemptFile = map[string]interface{} {
  "id": "gradeAttemptFileId",
  "name": "Grade Attempt File Name",
  "viewUrl": "localhost/some/attempt/file",
}

var rawCourseGradeAttemptFile2 = map[string]interface{} {
  "id": "gradeAttemptFileId2",
  "name": "Grade Attempt File Name #2",
  "viewUrl": "localhost/some/attempt/file/2",
}

var rawCourseGradeAttemptFile3 = map[string]interface{} {
  "id": "gradeAttemptFileId3",
  "name": "Grade Attempt File Name #23",
  "viewUrl": "localhost/some/attempt/file/3",
}
