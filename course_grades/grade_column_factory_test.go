package course_grades_test

import (
  "testing"

  courseGrades "github.com/jbaxe2/blackboard.rest/course_grades"
)

/**
 * The [TestCreateNewGradeColumn] function...
 */
func TestCreateNewGradeColumn (t *testing.T) {
  println ("Create a new grade column instance.")

  gradeColumn := courseGrades.NewGradeColumn (rawGradeColumn1)

  if gradeColumn.Id != rawGradeColumn1["id"] {
    t.Error ("Creating the grade column instance should have expected values.")
  }
}

/**
 * The [TestCreateNewGradeColumns] function...
 */
func TestCreateNewGradeColumns (t *testing.T) {
  println ("Create multiple new grade column instances.")

  gradeColumns := courseGrades.NewGradeColumns (
    []map[string]interface{} {rawGradeColumn1, rawGradeColumn2},
  )

  if !(2 == len (gradeColumns) && gradeColumns[0].Id == rawGradeColumn1["id"] &&
       gradeColumns[1].Id == rawGradeColumn2["id"]) {
    t.Error ("Creating the grade column instances should have expected values.")
  }
}

/**
 * Mocked instances to run the above tests with.
 */
var rawGradeColumn1 = map[string]interface{} {
  "id": "gradeColumnId1",
  "externalId": "externalGradeColumnId1",
  "externalToolId": "externalToolId1",
  "name": "Grade Column 1",
  "displayName": "Grade Column 1",
  "description": "Grading column.",
  "externalGrade": true,
  "created": "2021-11-26T16:07:41.875Z",
  "contentId": "contentId1",
  "score": map[string]interface{} {
    "possible": 0,
  },
  "availability": map[string]interface{} {
    "available": "Yes",
  },
  "grading": map[string]interface{} {
    "type": "Attempts",
    "due": "2021-11-26T16:07:41.875Z",
    "attemptsAllowed": 0,
    "scoringModel": "Last",
    "schemaId": "schemaId1",
    "anonymousGrading": map[string]interface{} {
      "type": "None",
      "releaseAfter": "2021-11-26T16:07:41.875Z",
    },
  },
  "gradebookCategoryId": "categoryId1",
  "formula": map[string]interface{} {
    "formula": "string",
    "aliases": map[string]interface{} {
      "key1": "value1",
      "key2": "value2",
    },
  },
  "includeInCalculations": true,
  "showStatisticsToStudents": true,
  "scoreProviderHandle": "string",
}

var rawGradeColumn2 = map[string]interface{} {
  "id": "gradeColumnId2",
  "externalId": "externalGradeColumnId2",
  "externalToolId": "externalToolId2",
  "name": "Grade Column 2",
  "displayName": "Grade Column 2",
  "description": "Grading column.",
  "externalGrade": true,
  "created": "2021-11-26T16:07:41.875Z",
  "contentId": "contentId2",
  "score": map[string]interface{} {
    "possible": 0,
  },
  "availability": map[string]interface{} {
    "available": "Yes",
  },
  "grading": map[string]interface{} {
    "type": "Attempts",
    "due": "2021-11-26T16:07:41.875Z",
    "attemptsAllowed": 0,
    "scoringModel": "Last",
    "schemaId": "schemaId1",
    "anonymousGrading": map[string]interface{} {
      "type": "None",
      "releaseAfter": "2021-11-26T16:07:41.875Z",
    },
  },
  "gradebookCategoryId": "categoryId1",
  "formula": map[string]interface{} {
    "formula": "string",
    "aliases": map[string]interface{} {
      "key1": "value1",
      "key2": "value2",
    },
  },
  "includeInCalculations": true,
  "showStatisticsToStudents": true,
  "scoreProviderHandle": "string",
}
