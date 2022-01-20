package blackboard_rest

import (
  "github.com/jbaxe2/blackboard.rest/api"
  "github.com/jbaxe2/blackboard.rest/course_groups"
)

/**
 * The [CourseGroups] interface provides the base interface for interacting with
 * the REST API's course groups service.
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

/**
 * The [_CourseGroups] type implements the Course Groups interface.
 */
type _CourseGroups struct {
  service api.Service

  CourseGroups
}

/**
 * The [NewCourseGroups] function creates a new course groups instance.
 */
func NewCourseGroups (service api.Service) CourseGroups {
  if nil == service {
    return nil
  }

  return &_CourseGroups {
    service: service,
  }
}
