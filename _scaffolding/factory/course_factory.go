package factory

import (
  "time"

  "github.com/jbaxe2/blackboard.rest/courses"
)

/**
 * The [NewCourse] function...
 */
func NewCourse (rawCourse map[string]interface{}) courses.Course {
  created, _ := time.Parse (time.RFC3339, rawCourse["created"].(string))

  return courses.Course {
    Id: rawCourse["id"].(string),
    Uuid: rawCourse["uuid"].(string),
    ExternalId: rawCourse["externalId"].(string),
    DataSourceId: rawCourse["dataSourceId"].(string),
    CourseId: rawCourse["courseId"].(string),
    Name: rawCourse["name"].(string),
    TermId: rawCourse["termId"].(string),
    Organization: rawCourse["organization"].(bool),
    Created: created,
  }
}
