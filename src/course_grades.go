package blackboard_rest

import (
  "github.com/jbaxe2/blackboard.rest.go/src/_scaffolding/config"
  error2 "github.com/jbaxe2/blackboard.rest.go/src/_scaffolding/error"
  "github.com/jbaxe2/blackboard.rest.go/src/_scaffolding/factory/course_grades"
  "github.com/jbaxe2/blackboard.rest.go/src/_scaffolding/services"
  "github.com/jbaxe2/blackboard.rest.go/src/course_grades"
  "github.com/jbaxe2/blackboard.rest.go/src/oauth2"
  "net/url"
  "strings"
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
  host url.URL

  accessToken oauth2.AccessToken

  service services.BlackboardRestService

  CourseGrades
}

/**
 * The [GetCourseGradesInstance] function...
 */
func GetCourseGradesInstance (
  host string, accessToken oauth2.AccessToken,
) CourseGrades {
  hostUri, _ := url.Parse (host)

  courseGradesService := &_BbRestCourseGrades {
    host: *hostUri, accessToken: accessToken,
  }

  courseGradesService.service.SetAccessToken (accessToken)

  return courseGradesService
}

/**
 * The [GetGradeColumns] method...
 */
func (restGrades *_BbRestCourseGrades) GetGradeColumns (
  courseId string,
) ([]course_grades.GradeColumn, error) {
  var columns []course_grades.GradeColumn
  var err error
  var result interface{}

  endpoint := config.CourseGradesEndpoints["grade_columns"]
  endpoint = strings.Replace (endpoint, "{courseId}", courseId, -1)

  result, err = restGrades.service.Connector.SendBbRequest (
    endpoint, "GET", make (map[string]interface{}), 2,
  )

  if (nil != err) && (error2.RestError{} != err) {
    return columns, err.(error2.CourseGradesError)
  }

  rawColumns := result.(map[string]interface{})["results"]

  columns = factory.NewGradeColumns (
    _normalizeRawResponse (rawColumns.([]interface{})),
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
  var err error
  var result interface{}

  endpoint := config.CourseGradesEndpoints["grade_column"]
  endpoint = strings.Replace (endpoint, "{courseId}", courseId, -1)
  endpoint = strings.Replace (endpoint, "{columnId}", columnId, -1)

  result, err = restGrades.service.Connector.SendBbRequest (
    endpoint, "GET", make (map[string]interface{}), 2,
  )

  if (nil != err) && (error2.RestError{} != err) {
    return column, err.(error2.CourseGradesError)
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
  var attempts []course_grades.Attempt
  var err error
  var result interface{}

  endpoint := config.CourseGradesEndpoints["column_attempts"]
  endpoint = strings.Replace (endpoint, "{courseId}", courseId, -1)
  endpoint = strings.Replace (endpoint, "{columnId}", columnId, -1)

  result, err = restGrades.service.Connector.SendBbRequest (
    endpoint, "GET", make (map[string]interface{}), 2,
  )

  if (nil != err) && (error2.RestError{} != err) {
    return attempts, err.(error2.CourseGradesError)
  }

  rawAttempts := result.(map[string]interface{})["results"]

  attempts = factory.NewColumnAttempts (
    _normalizeRawResponse (rawAttempts.([]interface{})),
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
  var err error
  var result interface{}

  endpoint := config.CourseGradesEndpoints["column_attempt"]
  endpoint = strings.Replace (endpoint, "{courseId}", courseId, -1)
  endpoint = strings.Replace (endpoint, "{columnId}", columnId, -1)
  endpoint = strings.Replace (endpoint, "{attemptId}", attemptId, -1)

  result, err = restGrades.service.Connector.SendBbRequest (
    endpoint, "GET", make (map[string]interface{}), 2,
  )

  if (nil != err) && (error2.RestError{} != err) {
    return attempt, err.(error2.CourseGradesError)
  }

  attempt = factory.NewColumnAttempt (result.(map[string]interface{}))

  return attempt, err
}

/**
 * The [_normalizeRawResponse] function...
 */
func _normalizeRawResponse(rawResponse []interface{}) []map[string]interface{} {
  mappedResponse := make ([]map[string]interface{}, len (rawResponse))

  for i, rawColumn := range rawResponse {
    mappedResponse[i] = rawColumn.(map[string]interface{})
  }

  return mappedResponse
}
