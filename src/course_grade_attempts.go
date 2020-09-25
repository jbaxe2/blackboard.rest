package blackboard_rest

import (
  "net/url"
  "strings"

  "github.com/jbaxe2/blackboard.rest.go/src/_scaffolding/config"
  "github.com/jbaxe2/blackboard.rest.go/src/_scaffolding/services"
  "github.com/jbaxe2/blackboard.rest.go/src/course_grade_attempts"
  "github.com/jbaxe2/blackboard.rest.go/src/oauth2"
)

/**
 * The [CourseGradeAttempts] interface...
 */
type CourseGradeAttempts interface {
  GetAttemptFileMetadataList (
    courseId string, attemptId string,
  ) ([]course_grade_attempts.AttemptFile, error)

  AttachFile() error

  GetAttemptFileMetadata (
    courseId string, attemptId string, attemptFileId string,
  ) (course_grade_attempts.AttemptFile, error)

  DownloadAttemptFile (
    courseId string, attemptId string, attemptFileId string,
  ) ([]byte, error)
}

/**
 * The [_BbRestCourseGradeAttempts] type...
 */
type _BbRestCourseGradeAttempts struct {
  host url.URL

  accessToken oauth2.AccessToken

  service services.BlackboardRestService

  CourseGradeAttempts
}

/**
 * The [GetCourseGradeAttemptsInstance] function...
 */
func GetCourseGradeAttemptsInstance (
  host string, accessToken oauth2.AccessToken,
) CourseGradeAttempts {
  hostUri, _ := url.Parse (host)

  courseGradeAttemptsService := &_BbRestCourseGradeAttempts {
    host: *hostUri, accessToken: accessToken,
  }

  courseGradeAttemptsService.service.SetHost (host)
  courseGradeAttemptsService.service.SetAccessToken (accessToken)

  return courseGradeAttemptsService
}

/**
 * The [GetAttemptFileMetadataList] method...
 */
func (restGradeAttempts *_BbRestCourseGradeAttempts) GetAttemptFileMetadataList (
  courseId string, attemptId string,
) ([]course_grade_attempts.AttemptFile, error) {
  var attemptFiles []course_grade_attempts.AttemptFile
  var err error
  var result interface{}

  endpoint := config.CourseGradeAttemptsEndpoints["file_metadata_list"]
  endpoint = strings.Replace (endpoint, "{courseId}", courseId, -1)
  endpoint = strings.Replace (endpoint, "{attemptId}", attemptId, -1)

  result, err = restGradeAttempts.service.Connector.SendBbRequest (
    endpoint, "GET", make (map[string]interface{}), 1,
  )

  return attemptFiles, err
}
