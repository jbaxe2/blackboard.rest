package blackboard_rest

import (
  "github.com/jbaxe2/blackboard.rest/api"
  "github.com/jbaxe2/blackboard.rest/course_grade_attempts"
)

/**
 * The [CourseGradeAttempts] interface provides the base interface for interacting
 * with the REST API's course grade attempts service.
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
 * The [_CourseGradeAttempts] type implements the Course Grade Attempts interface.
 */
type _CourseGradeAttempts struct {
  service api.Service

  CourseGradeAttempts
}

/**
 * The [NewCourseGradeAttempts] function creates a new Course Grade Attempts
 * instance.
 */
func NewCourseGradeAttempts (service api.Service) CourseGradeAttempts {
  if nil == service {
    return nil
  }

  return &_CourseGradeAttempts {
    service: service,
  }
}
