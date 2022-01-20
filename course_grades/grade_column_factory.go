package course_grades

import "github.com/jbaxe2/blackboard.rest/utils"

/**
 * The [NewGradeColumns] function...
 */
func NewGradeColumns (rawGradeColumns []map[string]interface{}) []GradeColumn {
  gradeColumns := make ([]GradeColumn, len (rawGradeColumns))

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
) GradeColumn {
  return GradeColumn {
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
func _parseScore (rawScore map[string]interface{}) Scoring {
  return Scoring {
    Possible: utils.NormalizeNumeric (rawScore["possible"]),
  }
}

/**
 * The [_parseAvailability] function...
 */
func _parseAvailability (rawAvailability map[string]interface{}) GradeAvailability {
  return GradeAvailability (rawAvailability["available"].(string))
}

/**
 * The [_parseGrading] function...
 */
func _parseGrading (rawGrading map[string]interface{}) Grading {
  return Grading {
    Type: GradingType (rawGrading["type"].(string)),
    SchemaId: rawGrading["schemaId"].(string),
  }
}
