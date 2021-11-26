package course_groups

import (
  "github.com/google/uuid"
)

/**
 * The [NewCourseGroups] function...
 */
func NewCourseGroups (rawCourseGroups []map[string]interface{}) []Group {
  courseGroups := make ([]Group, len (rawCourseGroups))

  for i, rawCourseGroup := range rawCourseGroups {
    courseGroups[i] = NewCourseGroup (rawCourseGroup)
  }

  return courseGroups
}

/**
 * The [NewCourseGroup] function...
 */
func NewCourseGroup (rawCourseGroup map[string]interface{}) Group {
  groupSetId, _ := rawCourseGroup["groupSetId"].(string)
  description, _ := rawCourseGroup["description"].(string)
  groupUuid, _ := uuid.Parse (rawCourseGroup["uuid"].(string))

  return Group {
    Id: rawCourseGroup["id"].(string),
    ExternalId: rawCourseGroup["externalId"].(string),
    GroupSetId: groupSetId,
    Name: rawCourseGroup["name"].(string),
    Description: description,
    Uuid: groupUuid,
  }
}
