package test

import (
  "testing"

  blackboardRest "github.com/jbaxe2/blackboard.rest"
  "github.com/jbaxe2/blackboard.rest/_scaffolding/config"
)

/**
 * The [CourseGroupsTester] type...
 */
type CourseGroupsTester struct {
  t *testing.T

  Testable
}

/**
 * The [Run] method...
 */
func (tester *CourseGroupsTester) Run() {
  println ("Course groups:\n")

  _testGetGroupsForCourse (tester.t)
}

/**
 * The [_getCourseGroupsInstance] function...
 */
func _getCourseGroupsInstance() blackboardRest.CourseGroups {
  authorizer := TestersAuthorizer{}
  _ = authorizer.AuthorizeForTests()

  return blackboardRest.GetCourseGroupsInstance (
    config.Host, authorizer.accessToken,
  )
}

/**
 * The [_testGetGroupsForCourse] function...
 */
func _testGetGroupsForCourse (t *testing.T) {
  println ("Obtain the groups for a course by its external ID.")

  groupsService := _getCourseGroupsInstance()

  _, err := groupsService.GetGroups ("externalId:wsu_educ_cap_2020spring")

  if nil != err {
    t.Error ("Failed to retrieve the groups for the course.")
    t.Error (err.Error())

    t.FailNow()
  }
}
