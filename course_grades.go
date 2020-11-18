package blackboard_rest

import (
  "net/url"
  "strings"

  "github.com/jbaxe2/blackboard.rest/_scaffolding"
  "github.com/jbaxe2/blackboard.rest/_scaffolding/config"
  "github.com/jbaxe2/blackboard.rest/_scaffolding/factory/course_grades"
  "github.com/jbaxe2/blackboard.rest/course_grades"
  "github.com/jbaxe2/blackboard.rest/oauth2"
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

/**
 * The [_BbRestCourseGrades] type...
 */
type _BbRestCourseGrades struct {
  _BlackboardRest

  CourseGrades
}

/**
 * The [GetCourseGradesInstance] function...
 */
func GetCourseGradesInstance (
  host string, accessToken oauth2.AccessToken,
) CourseGrades {
  hostUri, _ := url.Parse (host)

  courseGradesService := new (_BbRestCourseGrades)

  courseGradesService.host = *hostUri
  courseGradesService.accessToken = accessToken

  courseGradesService.service.SetHost (host)
  courseGradesService.service.SetAccessToken (accessToken)

  return courseGradesService
}

/**
 * The [GetGradeColumns] method...
 */
func (restGrades *_BbRestCourseGrades) GetGradeColumns (
  courseId string,
) ([]course_grades.GradeColumn, error) {
  endpoint := config.CourseGradesEndpoints["grade_columns"]
  endpoint = strings.Replace (endpoint, "{courseId}", courseId, -1)

  result, err := restGrades.service.Connector.SendBbRequest (
    endpoint, "GET", make (map[string]interface{}), 2,
  )

  if nil != err {
    return []course_grades.GradeColumn{}, err
  }

  rawColumns := result.(map[string]interface{})["results"]

  columns := factory.NewGradeColumns (
    _scaffolding.NormalizeRawResponse (rawColumns.([]interface{})),
  )

  return columns, err
}

/**
 * The [GetGradeColumn] method...
 */
func (restGrades *_BbRestCourseGrades) GetGradeColumn (
  courseId string, columnId string,
) (course_grades.GradeColumn, error) {
  var column course_grades.GradeColumn

  endpoint := config.CourseGradesEndpoints["grade_column"]
  endpoint = strings.Replace (endpoint, "{courseId}", courseId, -1)
  endpoint = strings.Replace (endpoint, "{columnId}", columnId, -1)

  result, err := restGrades.service.Connector.SendBbRequest (
    endpoint, "GET", make (map[string]interface{}), 2,
  )

  if nil != err {
    return column, err
  }

  column = factory.NewGradeColumn (result.(map[string]interface{}))

  return column, err
}

/**
 * The [GetColumnAttempts] method...
 */
func (restGrades *_BbRestCourseGrades) GetColumnAttempts (
  courseId string, columnId string,
) ([]course_grades.Attempt, error) {
  endpoint := config.CourseGradesEndpoints["column_attempts"]
  endpoint = strings.Replace (endpoint, "{courseId}", courseId, -1)
  endpoint = strings.Replace (endpoint, "{columnId}", columnId, -1)

  result, err := restGrades.service.Connector.SendBbRequest (
    endpoint, "GET", make (map[string]interface{}), 2,
  )

  if nil != err {
    return []course_grades.Attempt{}, err
  }

  rawAttempts := result.(map[string]interface{})["results"]

  attempts := factory.NewColumnAttempts (
    _scaffolding.NormalizeRawResponse (rawAttempts.([]interface{})),
  )

  return attempts, err
}

/**
 * The [GetColumnAttempt] method...
 */
func (restGrades *_BbRestCourseGrades) GetColumnAttempt (
  courseId string, columnId string, attemptId string,
) (course_grades.Attempt, error) {
  var attempt course_grades.Attempt

  endpoint := config.CourseGradesEndpoints["column_attempt"]
  endpoint = strings.Replace (endpoint, "{courseId}", courseId, -1)
  endpoint = strings.Replace (endpoint, "{columnId}", columnId, -1)
  endpoint = strings.Replace (endpoint, "{attemptId}", attemptId, -1)

  result, err := restGrades.service.Connector.SendBbRequest (
    endpoint, "GET", make (map[string]interface{}), 2,
  )

  if nil != err {
    return attempt, err
  }

  attempt = factory.NewColumnAttempt (result.(map[string]interface{}))

  return attempt, err
}
