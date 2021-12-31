package courses

import (
  "time"
)

/**
 * The [NewCourses] function creates a slice of new Course instances from slice
 * of raw maps.
 */
func NewCourses (rawCourses []map[string]interface{}) []Course {
  newCourses := make ([]Course, len (rawCourses))

  for i, rawCourse := range rawCourses {
    newCourses[i] = NewCourse (rawCourse)
  }

  return newCourses
}

/**
 * The [NewCourse] function creates a new Course instance from a raw map.
 */
func NewCourse (rawCourse map[string]interface{}) Course {
  created, _ := time.Parse (time.RFC3339, rawCourse["created"].(string))
  parentId, _ := rawCourse["parentId"].(string)
  termId, _ := rawCourse["termId"].(string)

  return Course {
    Id: rawCourse["id"].(string),
    Uuid: rawCourse["uuid"].(string),
    ExternalId: rawCourse["externalId"].(string),
    DataSourceId: rawCourse["dataSourceId"].(string),
    CourseId: rawCourse["courseId"].(string),
    Name: rawCourse["name"].(string),
    TermId: termId,
    ParentId: parentId,
    Organization: rawCourse["organization"].(bool),
    Created: created,
  }
}
