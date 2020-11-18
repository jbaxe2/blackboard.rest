package course_grades

import "net/url"

/**
 * The [AttemptFile] type...
 */
type AttemptFile struct {
  Id, Name string

  ViewUrl url.URL
}

/**
 * The [UploadedFile] type...
 */
type UploadedFile struct {
  Name, UploadId string
}
