package blackboard_rest

import (
  "github.com/jbaxe2/blackboard.rest/course_groups"
)

/**
 * The [CourseGroups] interface...
 */
type CourseGroups interface {
  GetGroups (courseId string) ([]course_groups.Group, error)

  CreateGroup (courseId string, group course_groups.Group) error

  GetGroupSets (courseId string) ([]course_groups.Group, error)

  CreateGroupSet (courseId string, groupSet course_groups.Group) error

  GetGroupSet (courseId string, groupSetId string)  (course_groups.Group, error)

  GetGroupSetGroups (
    courseId string, groupSetId string,
  ) ([]course_groups.Group, error)

  CreateGroupInSet (
    courseId string, groupSetId string, group course_groups.Group,
  ) error

  GetGroup (courseId string, groupSetId string) (course_groups.Group, error)
}
