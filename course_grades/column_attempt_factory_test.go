package course_grades_test

import (
  "testing"

  courseGrades "github.com/jbaxe2/blackboard.rest/course_grades"
)

/**
 * The [TestCreateNewColumnAttempt] function...
 */
func TestCreateNewColumnAttempt (t *testing.T) {
  println ("Create a new column attempt instance.")

  attempt := courseGrades.NewColumnAttempt (rawAttempt1)

  if attempt.Id != rawAttempt1["id"] {
    t.Error ("Creating a new column attempt should have the expected results.")
  }
}

/**
 * The [TestCreateNewColumnAttempts] function...
 */
func TestCreateNewColumnAttempts (t *testing.T) {
  println ("Create multiple column attempts instances.")

  attempts :=
    courseGrades.NewColumnAttempts ([]map[string]interface{} {rawAttempt1, rawAttempt2})

  if !(2 == len (attempts) && attempts[0].Id == rawAttempt1["id"] &&
       attempts[1].Id == rawAttempt2["id"]) {
    t.Error ("Creating multiple column attempts should have the expected results.")
  }
}

/**
 * Mocked instances to run the above tests with.
 */
var rawAttempt1 = map[string]interface{} {
  "id": "gradeAttemptId1",
  "userId": "userId1",
  "groupAttemptId": "groupAttemptId1",
  "groupOverride": true,
  "status": "NotAttempted",
  "displayGrade": map[string]interface{} {
    "scaleType": "Percent",
    "score": 100,
    "text": "Display Text",
  },"text": "Text Grade",
  "score": 100,
  "reconciliationMode": "Average",
  "notes": "Some grade attempt notes.",
  "feedback": "Some grade attempt feedback.",
  "studentComments": "Student comments for the attempt.",
  "studentSubmission": "The student's textual submission.",
  "exempt": false,
  "created": "2021-11-26T14:03:07.808Z",
  "attemptDate": "2021-11-26T14:03:07.808Z",
  "modified": "2021-11-26T14:03:07.808Z",
  "attemptReceipt": map[string]interface{} {
    "receiptId": "receiptId1",
    "submissionDate": "2021-11-26T14:03:07.808Z",
  },
}

var rawAttempt2 = map[string]interface{} {
  "id": "gradeAttemptId2",
  "userId": "userId1",
  "groupAttemptId": "groupAttemptId2",
  "groupOverride": true,
  "status": "NotAttempted",
  "displayGrade": map[string]interface{} {
    "scaleType": "Percent",
    "score": 100,
    "text": "Display Text",
  },"text": "Text Grade",
  "score": 100,
  "reconciliationMode": "Average",
  "notes": "Some grade attempt notes.",
  "feedback": "Some grade attempt feedback.",
  "studentComments": "Student comments for the attempt.",
  "studentSubmission": "The student's textual submission.",
  "exempt": false,
  "created": "2021-11-26T14:03:07.808Z",
  "attemptDate": "2021-11-26T14:03:07.808Z",
  "modified": "2021-11-26T14:03:07.808Z",
  "attemptReceipt": map[string]interface{} {
    "receiptId": "receiptId1",
    "submissionDate": "2021-11-26T14:03:07.808Z",
  },
}
