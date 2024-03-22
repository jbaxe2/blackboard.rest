package course_memberships

import (
  "time"

  "github.com/jbaxe2/blackboard.rest/courses"
  "github.com/jbaxe2/blackboard.rest/users"
)

/**
 * The [NewMemberships] function creates a slice of Membership instances based on
 * a slice of raw string-based maps containing the memberships' information.
 */
func NewMemberships (rawMemberships []map[string]interface{}) []Membership {
  memberships := make ([]Membership, len (rawMemberships))

  for i, rawMembership := range rawMemberships {
    memberships[i] = NewMembership (rawMembership)
  }

  return memberships
}

/**
 * The [NewMembership] function creates a Membership instance based on a raw
 * string-based map containing the membership's information.
 */
func NewMembership (rawMembership map[string]interface{}) Membership {
  created, _ := time.Parse (time.RFC3339, rawMembership["created"].(string))

  return Membership {
    Id: rawMembership["id"].(string),
    CourseId: rawMembership["courseId"].(string),
    UserId: rawMembership["userId"].(string),
    User: _parseUser (rawMembership["user"]),
    Course: _parseCourse (rawMembership["course"]),
    Created: created,
    CourseRoleId: _parseCourseRole (rawMembership["courseRoleId"].(string)),
    Availability:
      _parseAvailability (rawMembership["availability"].(map[string]interface{})),
  }
}

/**
 * The [_parseUser] function checks to see if there is user information contained
 * in the raw data, and if so creates the correlating User instance.
 */
func _parseUser (user interface{}) users.User {
  var newUser users.User

  if rawUser, haveUser := user.(map[string]interface{}); haveUser {
    newUser = users.NewUser (rawUser)
  }

  return newUser
}

func _parseCourse (rawCourse interface{}) courses.Course {
  var newCourse courses.Course

  if someCourse, haveCourse := rawCourse.(map[string]interface{}); haveCourse {
    newCourse = courses.NewCourse (someCourse)
  }

  return newCourse
}

/**
 * The [_parseAvailability] function returns the Membership Availability for the
 * membership.
 */
func _parseAvailability (
  availability map[string]interface{},
) MembershipAvailability {
  return MembershipAvailability (
    availability["available"].(string),
  )
}

/**
 * The [_parseCourseRole] function returns the Membership Role for the membership.
 */
func _parseCourseRole (roleId string) MembershipRole {
  return MembershipRole (roleId)
}
