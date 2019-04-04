package error

/**
 * The [CourseMembershipsError] type...
 */
type CourseMembershipsError struct {
  RestableError

  RestError
}

/**
 * The [Error] method...
 */
func (membershipsError CourseMembershipsError) Error() string {
  return membershipsError.RestError.Error()
}
