package factory

import "github.com/jbaxe2/blackboard.rest/course_groups"

/**
 * The [NewCourseGroups] function...
 */
func NewCourseGroups (
  rawCourseGroups []map[string]interface{},
) []course_groups.Group {
  courseGroups := make ([]course_groups.Group, len (rawCourseGroups))

  for i, rawCourseGroup := range rawCourseGroups {
    courseGroups[i] = NewCourseGroup (rawCourseGroup)
  }

  return courseGroups
}

/**
 * The [NewCourseGroup] function...
 */
func NewCourseGroup (rawCourseGroup map[string]interface{}) course_groups.Group {
  return course_groups.Group{}
}
