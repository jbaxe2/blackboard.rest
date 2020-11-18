package factory

import "github.com/jbaxe2/blackboard.rest/course_grades"

/**
 * The [NewGradeColumns] function...
 */
func NewGradeColumns (
  rawGradeColumns []map[string]interface{},
) []course_grades.GradeColumn {
  gradeColumns := make ([]course_grades.GradeColumn, len (rawGradeColumns))

  for i, rawGradeColumn := range rawGradeColumns {
    gradeColumns[i] = NewGradeColumn (rawGradeColumn)
  }

  return gradeColumns
}

/**
 * The [NewGradeColumn] function...
 */
func NewGradeColumn (
  rawGradeColumn map[string]interface{},
) course_grades.GradeColumn {
  return course_grades.GradeColumn {
    Id: rawGradeColumn["id"].(string),
    Name: rawGradeColumn["name"].(string),
    Score: _parseScore (rawGradeColumn["score"].(map[string]interface{})),
    Availability:
      _parseAvailability (rawGradeColumn["availability"].(map[string]interface{})),
    Grading:
      _parseGrading (rawGradeColumn["grading"].(map[string]interface{})),
  }
}

/**
 * The [_parseScore] function...
 */
func _parseScore (rawScore map[string]interface{}) course_grades.Scoring {
  return course_grades.Scoring {
    Possible: rawScore["possible"].(float64),
  }
}

/**
 * The [_parseAvailability] function...
 */
func _parseAvailability (
  rawAvailability map[string]interface{},
) course_grades.GradeAvailability {
  return course_grades.GradeAvailability (rawAvailability["available"].(string))
}

/**
 * The [_parseGrading] function...
 */
func _parseGrading (rawGrading map[string]interface{}) course_grades.Grading {
  return course_grades.Grading {
    Type: course_grades.GradingType (rawGrading["type"].(string)),
    SchemaId: rawGrading["schemaId"].(string),
  }
}
