package blackboard_rest

import (
  "github.com/jbaxe2/blackboard.rest/course_grades"
)

/**
 * The [CourseGrades] interface...
 */
type CourseGrades interface {
  GetGradeColumns (courseId string) ([]course_grades.GradeColumn, error)

  CreateGradeColumn (courseId string, column course_grades.GradeColumn) error

  GetGradeColumn (
    courseId string, columnId string,
  ) (course_grades.GradeColumn, error)

  UpdateGradeColumn (
    courseId string, columnId string, column course_grades.GradeColumn,
  ) error

  GetColumnAttempts (
    courseId string, columnId string,
  ) ([]course_grades.Attempt, error)

  CreateColumnAttempt (columnId string, attempt course_grades.Attempt) error

  GetColumnAttempt (
    courseId string, columnId string, attemptId string,
  ) (course_grades.Attempt, error)

  UpdateColumnAttempt (
    columnId string, attemptId string, attempt course_grades.Attempt,
  ) error

  GetColumnGrades (courseId string, columnId string) ([]course_grades.Grade, error)

  GetColumnGradeLastChanged (
    courseId string, columnId string,
  ) (course_grades.Grade, error)

  GetColumnGrade (
    courseId string, columnId string, userId string,
  ) (course_grades.Grade, error)

  UpdateColumnGrade (
    courseId string, columnId string, userId string, grade course_grades.Grade,
  ) error

  GetUserGrades (courseId string, userId string) ([]course_grades.Grade, error)
}
