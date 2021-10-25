package course_memberships

import (
  "time"

  "github.com/jbaxe2/blackboard.rest/users"
)

/**
 * The [NewMemberships] function...
 */
func NewMemberships (rawMemberships []map[string]interface{}) []Membership {
  memberships := make ([]Membership, len (rawMemberships))

  for i, rawMembership := range rawMemberships {
    memberships[i] = NewMembership (rawMembership)
  }

  return memberships
}

/**
 * The [NewMembership] function...
 */
func NewMembership (rawMembership map[string]interface{}) Membership {
  created, _ := time.Parse (time.RFC3339, rawMembership["created"].(string))

  return Membership {
    Id: rawMembership["id"].(string),
    CourseId: rawMembership["courseId"].(string),
    UserId: rawMembership["userId"].(string),
    User: _parseUser (rawMembership["user"]),
    Created: created,
    CourseRoleId: _parseCourseRole (rawMembership["courseRoleId"].(string)),
    Availability:
      _parseAvailability (rawMembership["availability"].(map[string]interface{})),
  }
}

/**
 * The [_parseUser] function...
 */
func _parseUser (user interface{}) users.User {
  var newUser users.User

  if rawUser, haveUser := user.(map[string]interface{}); haveUser {
    newUser = users.NewUser (rawUser)
  }

  return newUser
}

/**
 * The [_parseAvailability] function...
 */
func _parseAvailability (
  availability map[string]interface{},
) MembershipAvailability {
  return MembershipAvailability(
    availability["available"].(string),
  )
}

/**
 * The [_parseCourseRole] function...
 */
func _parseCourseRole (roleId string) MembershipRole {
  return MembershipRole (roleId)
}
