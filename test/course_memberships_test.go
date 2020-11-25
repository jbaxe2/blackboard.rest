package test

import (
  "testing"

  "github.com/jbaxe2/blackboard.rest"
  "github.com/jbaxe2/blackboard.rest/_scaffolding/config"
)

/**
 * The [CourseMembershipsTester] type...
 */
type CourseMembershipsTester struct {
  t *testing.T

  Testable
}

/**
 * The [Run] method...
 */
func (tester *CourseMembershipsTester) Run() {
  println ("\nCourse Memberships:")

  _testGetCourseMembershipsInstance (tester.t)
  _testGetCourseMembershipsByCoursePrimaryId (tester.t)
  _testGetCourseMembershipsByUserPrimaryId (tester.t)
  _testGetMembershipByCourseAndUserPrimaryIds (tester.t)
}

/**
 * The [_getCourseMembershipsInstance] function...
 */
func _getCourseMembershipsInstance() blackboard_rest.CourseMemberships {
  authorizer := TestersAuthorizer{}
  _ = authorizer.AuthorizeForTests()

  return blackboard_rest.GetCourseMembershipsInstance (
    config.Host, authorizer.accessToken,
  )
}

/**
 * The [_testGetCourseMembershipsInstance] function...
 */
func _testGetCourseMembershipsInstance (t *testing.T) {
  println ("Obtain a valid CourseMemberships service instance.")

  courseMembershipsService := _getCourseMembershipsInstance()

  if nil == courseMembershipsService {
    t.Error ("Obtaining a valid CourseMemberships instance failed.\n")
    t.FailNow()
  }
}

/**
 * The [_testGetCourseMembershipsByCoursePrimaryId] function...
 */
func _testGetCourseMembershipsByCoursePrimaryId (t *testing.T) {
  println ("Get a list of course memberships by the course primary ID.")

  membershipsService := _getCourseMembershipsInstance()

  memberships, err := membershipsService.GetMembershipsForCourse ("_121_1")

  if (nil == memberships) || (nil != err) {
    t.Error ("Failed to obtain the list of course memberships (course).\n")
    t.FailNow()
  }

  if 0 == len (memberships) {
    t.Error ("Retrieved an empty list of enrollments that should not be empty.")
    t.FailNow()
  }

  for _, membership := range memberships {
    if "_121_1" != membership.CourseId {
      t.Error ("Membership retrieved does not match what was specified.")
      t.FailNow()
    }
  }
}

/**
 * The [_testGetCourseMembershipsByUserPrimaryId] function...
 */
func _testGetCourseMembershipsByUserPrimaryId (t *testing.T) {
  println ("Get a list of course memberships by the user primary ID.")

  membershipsService := _getCourseMembershipsInstance()

  memberships, err := membershipsService.GetMembershipsForUser ("_27_1")

  if (nil == memberships) || (nil != err) {
    t.Error ("Failed to obtain the list of course memberships (user).\n")
    t.FailNow()
  }

  if 0 == len (memberships) {
    t.Error ("Retrieved an empty list of enrollments that should not be empty.")
    t.FailNow()
  }

  for _, membership := range memberships {
    if "_27_1" != membership.UserId {
      t.Error ("Membership retrieved does not match what was specified.")
      t.FailNow()
    }
  }
}

/**
 * The [_testGetMembershipByCourseAndUserPrimaryIds] function...
 */
func _testGetMembershipByCourseAndUserPrimaryIds (t *testing.T) {
  println ("Get a membership by the course and user primary ID's.")

  membershipsService := _getCourseMembershipsInstance()

  membership, err :=
    membershipsService.GetMembership ("_121_1", "_27_1")

  if (nil != err) || ("" == membership.Id) {
    t.Error ("Failed to obtain the membership for the course and user.")
    t.FailNow()
  }

  if ("_121_1" != membership.CourseId) && ("_27_1" != membership.UserId) {
    t.Error ("Membership retrieved does not match what was specified.")
  }
}
