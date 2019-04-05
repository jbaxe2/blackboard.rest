package error

/**
 * The [RestableError] interface...
 */
type RestableError interface {
  error
}

/**
 * The [RestError] type...
 */
type RestError struct {
  Status float64

  Code, Message, DeveloperMessage, ExtraInfo string

  RestableError
}

/**
 * The [CoursesError] type...
 */
type CoursesError RestError

/**
 * The [UsersError] type...
 */
type UsersError RestError

/**
 * The [CourseMembershipsError] type...
 */
type CourseMembershipsError RestError

/**
 * The [SystemError] type...
 */
type SystemError RestError

/**
 * The [Error] method...
 */
func (err RestError) Error() string {
  return "Error (" + err.Code + "): " + err.Message
}
