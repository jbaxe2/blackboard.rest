package blackboard_rest

import (
  "net/url"
  "strings"

  "github.com/jbaxe2/blackboard.rest/_scaffolding"
  "github.com/jbaxe2/blackboard.rest/_scaffolding/config"
  "github.com/jbaxe2/blackboard.rest/_scaffolding/factory"
  "github.com/jbaxe2/blackboard.rest/course_grade_attempts"
  "github.com/jbaxe2/blackboard.rest/oauth2"
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
  _BlackboardRest

  CourseGradeAttempts
}

/**
 * The [GetCourseGradeAttemptsInstance] function...
 */
func GetCourseGradeAttemptsInstance (
  host string, accessToken oauth2.AccessToken,
) CourseGradeAttempts {
  hostUri, _ := url.Parse (host)

  courseGradeAttemptsService := new (_BbRestCourseGradeAttempts)

  courseGradeAttemptsService.host = *hostUri
  courseGradeAttemptsService.accessToken = accessToken

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
  endpoint := config.CourseGradeAttemptsEndpoints["file_metadata_list"]
  endpoint = strings.Replace (endpoint, "{courseId}", courseId, -1)
  endpoint = strings.Replace (endpoint, "{attemptId}", attemptId, -1)

  result, err := restGradeAttempts.service.Connector.SendBbRequest (
    endpoint, "GET", make (map[string]interface{}), 1,
  )

  if nil != err {
    return []course_grade_attempts.AttemptFile{}, err
  }

  rawAttemptFiles := result.(map[string]interface{})["results"]

  attemptFiles := factory.NewAttemptFiles (
    _scaffolding.NormalizeRawResponse (rawAttemptFiles.([]interface{})),
  )

  return attemptFiles, err
}

/**
 * The [DownloadAttemptFile] method...
 */
func (restGradeAttempts *_BbRestCourseGradeAttempts) DownloadAttemptFile (
  courseId string, attemptId string, attemptFileId string,
) ([]byte, error) {
  endpoint := config.CourseGradeAttemptsEndpoints["download"]
  endpoint = strings.Replace (endpoint, "{courseId}", courseId, -1)
  endpoint = strings.Replace (endpoint, "{attemptId}", attemptId, -1)
  endpoint = strings.Replace (endpoint, "{attemptFileId}", attemptFileId, -1)

  result, err := restGradeAttempts.service.Connector.SendBbRequest (
    endpoint, "GET", make (map[string]interface{}), 1,
  )

  if nil != err {
    return []byte{}, err
  }

  return result.([]byte), err
}
