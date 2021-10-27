package blackboard_rest

import (
  "github.com/jbaxe2/blackboard.rest/course_grade_attempts"
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
