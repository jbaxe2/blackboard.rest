package error

/**
 * The [CoursesError] type...
 */
type CoursesError struct {
  RestableError

  RestError
}

/**
 * The [Error] method...
 */
func (coursesError CoursesError) Error() string {
  return coursesError.RestError.Error()
}
