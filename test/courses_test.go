package test

import (
  "testing"

  "github.com/jbaxe2/blackboard.rest"
  "github.com/jbaxe2/blackboard.rest/_scaffolding/config"
)

/**
 * The [CoursesTester] type...
 */
type CoursesTester struct {
  t *testing.T

  Testable
}

/**
 * The [Run] method...
 */
func (tester *CoursesTester) Run() {
  println ("\nCourses:")

  _testGetCoursesInstance (tester.t)
  _testGetCourseByPrimaryId (tester.t)
}

/**
 * The [_getUsersInstance] function...
 */
func _getCoursesInstance() blackboard_rest.Courses {
  authorizer := TestersAuthorizer{}
  _ = authorizer.AuthorizeForTests()

  return blackboard_rest.GetCoursesInstance (
    config.Host, authorizer.accessToken,
  )
}

/**
 * The [_testGetCoursesInstance] function...
 */
func _testGetCoursesInstance (t *testing.T) {
  println ("Obtain a valid Courses service instance.")

  coursesService := _getCoursesInstance()

  if nil == coursesService {
    t.Error ("Obtaining a valid Courses service instance failed.")
    t.FailNow()
  }
}

/**
 * The [_testGetCourseByPrimaryId] method...
 */
func _testGetCourseByPrimaryId (t *testing.T) {
  println ("Get a course by its primary ID.")

  coursesService := _getCoursesInstance()
  course, err := coursesService.GetCourse ("_101_1")

  if nil != err {
    t.Error ("Error while retrieving the course:\n" + err.Error())
    t.FailNow()
  }

  if "_101_1" != course.Id {
    t.Error ("The course retrieved does not match the one selected.")
    t.FailNow()
  }
}
