package course_group_users

/**
 * The [NewCourseGroupUsers] function...
 */
func NewCourseGroupUsers (
  rawGroupUsers []map[string]interface{},
) []GroupMembership {
  groupUsers := make ([]GroupMembership, len (rawGroupUsers))

  for i, rawGroupUser := range rawGroupUsers {
    groupUsers[i] = NewCourseGroupUser (rawGroupUser)
  }

  return groupUsers
}

/**
 * The [NewCourseGroupUser] function...
 */
func NewCourseGroupUser (rawGroupUser map[string]interface{}) GroupMembership {
  return GroupMembership {
    UserId: rawGroupUser["userId"].(string),
  }
}
