package test

import (
  "testing"

  blackboardRest "github.com/jbaxe2/blackboard.rest"
  "github.com/jbaxe2/blackboard.rest/_scaffolding/config"
)

/**
 * The [CourseGradeAttemptsTester] type...
 */
type CourseGradeAttemptsTester struct {
  t *testing.T

  Testable
}

/**
 * The [Run] method...
 */
func (tester *CourseGradeAttemptsTester) Run() {
  println ("\nCourse Grade Attempts:")

  _testGetValidCourseGradeAttemptsInstance (tester.t)
  _testGetAttemptFileMetadataList (tester.t)
}

/**
 * The [_getCourseGradeAttemptsInstance] function...
 */
func _getCourseGradeAttemptsInstance() blackboardRest.CourseGradeAttempts {
  authorizer := TestersAuthorizer{}
  _ = authorizer.AuthorizeForTests()

  return blackboardRest.GetCourseGradeAttemptsInstance (
    config.Host, authorizer.accessToken,
  )
}

/**
 * The [_testGetValidCourseGradeAttemptsInstance] function...
 */
func _testGetValidCourseGradeAttemptsInstance (t *testing.T) {
  println ("Obtain a valid course grade attempts instance.")

  if nil == _getCourseGradeAttemptsInstance() {
    t.Error ("Obtaining a course grade attempts instance failed.")
    t.FailNow()
  }
}

/**
 * The [_testGetAttemptFileMetadataList] function...
 */
func _testGetAttemptFileMetadataList (t *testing.T) {
  println ("Get a list of the files metadata for a course and attempt.")

  attemptsService := _getCourseGradeAttemptsInstance()

  attemptFiles, err := attemptsService.GetAttemptFileMetadataList (
    "externalId:wsu_educ_cap_2020fall", "_21235_1",
  )

  if nil != err {
    t.Error ("Error in obtaining the attempt file metadata list.")
    t.Error (err.Error())

    t.FailNow()
  }

  if 0 == len (attemptFiles) {
    t.Error ("The attempt files for the course and attempt should not be empty.")
    t.FailNow()
  }
}
