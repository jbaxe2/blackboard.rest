package factory

import "github.com/jbaxe2/blackboard.rest.go/src/course_group_users"

/**
 * The [NewCourseGroupUsers] function...
 */
func NewCourseGroupUsers (
  rawGroupUsers []map[string]interface{},
) []course_group_users.GroupMembership {
  groupUsers := make ([]course_group_users.GroupMembership, len (rawGroupUsers))

  for i, rawGroupUser := range rawGroupUsers {
    groupUsers[i] = NewCourseGroupUser (rawGroupUser)
  }

  return groupUsers
}

/**
 * The [NewCourseGroupUser] function...
 */
func NewCourseGroupUser (
  rawGroupUser map[string]interface{},
) course_group_users.GroupMembership {
  return course_group_users.GroupMembership {
    UserId: rawGroupUser["userId"].(string),
  }
}
