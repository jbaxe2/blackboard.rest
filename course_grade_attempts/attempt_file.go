package course_grade_attempts

import "net/url"

/**
 * The [AttemptFile] type...
 */
type AttemptFile struct {
  Id string

  Name string

  ViewUrl url.URL
}
