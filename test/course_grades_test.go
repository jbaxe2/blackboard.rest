package test

import (
  "testing"

  blackboardRest "github.com/jbaxe2/blackboard.rest"
  "github.com/jbaxe2/blackboard.rest/_scaffolding/config"
  "github.com/jbaxe2/blackboard.rest/course_grades"
)

/**
 * The [CourseGradesTester] type...
 */
type CourseGradesTester struct {
  t *testing.T

  Testable
}

/**
 * The [Run] method...
 */
func (tester *CourseGradesTester) Run() {
  println ("\nCourse Grades:")

  _testGetValidCourseGradesInstance (tester.t)
  _testGetGradeColumnsForCourseByPrimaryId (tester.t)
  _testGetGradeColumnForCourseByPrimaryIds (tester.t)
  _testGetColumnAttemptsByColumnPrimaryId (tester.t)
  _testGetColumnAttemptByColumnAndAttemptPrimaryId (tester.t)
}

/**
 * The [_getCourseGradesInstance] function...
 */
func _getCourseGradesInstance() blackboardRest.CourseGrades {
  authorizer := TestersAuthorizer{}
  _ = authorizer.AuthorizeForTests()

  return blackboardRest.GetCourseGradesInstance (
    config.Host, authorizer.accessToken,
  )
}

/**
 * The [_testGetValidCourseGradesInstance] function...
 */
func _testGetValidCourseGradesInstance (t *testing.T) {
  println ("Obtain a valid CourseGrades service instance.")

  if nil == _getCourseGradesInstance() {
    t.Error ("Obtaining a valid CourseGrades service instance failed.")
    t.FailNow()
  }
}

/**
 * The [_testGetGradeColumnsForCourseByPrimaryId] function...
 */
func _testGetGradeColumnsForCourseByPrimaryId (t *testing.T) {
  println ("Get grade columns for a course by its primary ID.")

  gradesService := _getCourseGradesInstance()

  columns, err := gradesService.GetGradeColumns ("_121_1")

  if (nil == columns) || (nil != err) {
    t.Error ("Failed to obtain the grade columns for the course.")
    t.FailNow()
  }

  if 0 == len (columns) {
    t.Error ("Retrieved an empty list of columns that should not be empty.")
    t.FailNow()
  }
}

/**
 * The [_testGetGradeColumnForCourseByPrimaryIds] function...
 */
func _testGetGradeColumnForCourseByPrimaryIds (t *testing.T) {
  println ("Get a grade column for a course by primary ID's.")

  gradesService := _getCourseGradesInstance()

  columns, err := gradesService.GetGradeColumns ("_121_1")
  column, err := gradesService.GetGradeColumn ("_121_1", columns[0].Id)

  if (course_grades.GradeColumn{} == column) || (nil != err) {
    t.Error ("Failed to retrieve the grade column.")
    t.FailNow()
  }

  if columns[0].Id != column.Id {
    t.Error ("The retrieved column does not match what was specified.")
    t.FailNow()
  }
}

/**
 * The [_testGetColumnAttemptsByColumnPrimaryId] function...
 */
func _testGetColumnAttemptsByColumnPrimaryId (t *testing.T) {
  println ("Get column attempts by a column's primary ID.")

  gradesService := _getCourseGradesInstance()

  columns, err := gradesService.GetGradeColumns ("_121_1")
  attempts, err := gradesService.GetColumnAttempts ("_121_1", columns[1].Id)

  if (nil == attempts) || (nil != err) {
    t.Error ("Failed to obtain the attempts for the grade column.")
    t.FailNow()
  }

  if 0 == len (attempts) {
    t.Error ("Retrieved an empty list of attempts that should not be empty.")
    t.FailNow()
  }
}

/**
 * The [_testGetColumnAttemptByColumnAndAttemptPrimaryId] function...
 */
func _testGetColumnAttemptByColumnAndAttemptPrimaryId (t *testing.T) {
  println ("Get an attempt by the column and attempt primary ID's.")

  gradesService := _getCourseGradesInstance()

  columns, err := gradesService.GetGradeColumns ("_121_1")
  attempts, err := gradesService.GetColumnAttempts ("_121_1", columns[1].Id)

  attempt, err := gradesService.GetColumnAttempt (
    "_121_1", columns[1].Id, attempts[0].Id,
  )

  if (course_grades.Attempt{} == attempt) || (nil != err) {
    t.Error ("Failed to retrieve the column attempt.")
    t.FailNow()
  }

  if attempts[0].Id != attempt.Id {
    t.Error ("The retrieved attempt does not match what was specified.")
    t.FailNow()
  }
}
