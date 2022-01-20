package blackboard_rest

import (
  "github.com/jbaxe2/blackboard.rest/api"
  "github.com/jbaxe2/blackboard.rest/course_group_users"
)

/**
 * The [CourseGroupUsers] interface provides the base interface for interacting
 * with the REST API's course group users endpoints.
 */
type CourseGroupUsers interface {
  GetGroupMemberships (
    courseId string, groupId string,
  ) ([]course_group_users.GroupMembership, error)

  GetGroupMembership (
    courseId string, groupId string, userId string,
  ) (course_group_users.GroupMembership, error)
}

/**
 * The [_CourseGroupUsers] type implements the Course Group Users interface.
 */
type _CourseGroupUsers struct {
  CourseGroupUsers
}

/**
 * The [NewCourseGroupUsers] function creates a new course group users instance.
 */
func NewCourseGroupUsers (service api.Service) CourseGroupUsers {
  return new (_CourseGroupUsers)
}
