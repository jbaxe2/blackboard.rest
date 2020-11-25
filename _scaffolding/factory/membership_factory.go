package factory

import (
  "time"

  "github.com/jbaxe2/blackboard.rest/course_memberships"
  "github.com/jbaxe2/blackboard.rest/users"
)

/**
 * The [NewMemberships] function...
 */
func NewMemberships (
  rawMemberships []map[string]interface{},
) []course_memberships.Membership {
  memberships := make ([]course_memberships.Membership, len (rawMemberships))

  for i, rawMembership := range rawMemberships {
    memberships[i] = NewMembership (rawMembership)
  }

  return memberships
}

/**
 * The [NewMembership] function...
 */
func NewMembership (
  rawMembership map[string]interface{},
) course_memberships.Membership {
  created, _ := time.Parse (time.RFC3339, rawMembership["created"].(string))

  return course_memberships.Membership {
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
    newUser = NewUser (rawUser)
  }

  return newUser
}

/**
 * The [_parseAvailability] function...
 */
func _parseAvailability (
  availability map[string]interface{},
) course_memberships.MembershipAvailability {
  return course_memberships.MembershipAvailability (
    availability["available"].(string),
  )
}

/**
 * The [_parseCourseRole] function...
 */
func _parseCourseRole (roleId string) course_memberships.MembershipRole {
  return course_memberships.MembershipRole (roleId)
}
