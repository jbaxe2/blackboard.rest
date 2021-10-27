package blackboard_rest

import (
  "github.com/jbaxe2/blackboard.rest/course_group_users"
)

/**
 * The [CourseGroupUsers] interface...
 */
type CourseGroupUsers interface {
  GetGroupMemberships (
    courseId string, groupId string,
  ) ([]course_group_users.GroupMembership, error)

  GetGroupMembership (
    courseId string, groupId string, userId string,
  ) (course_group_users.GroupMembership, error)
}
