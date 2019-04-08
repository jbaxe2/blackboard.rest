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
  Status, Code, Message, DeveloperMessage, ExtraInfo string

  RestableError
}

/**
 * The [OAuth2Error] type...
 */
type OAuth2Error = RestError

/**
 * The [CourseGradesError] type...
 */
type CourseGradesError = RestError

/**
 * The [CourseMembershipsError] type...
 */
type CourseMembershipsError = RestError

/**
 * The [CoursesError] type...
 */
type CoursesError = RestError

/**
 * The [RolesError] type...
 */
type RolesError = RestError

/**
 * The [SystemError] type...
 */
type SystemError = RestError

/**
 * The [TermsError] type...
 */
type TermsError = RestError

/**
 * The [UsersError] type...
 */
type UsersError = RestError

/**
 * The [Error] method...
 */
func (err RestError) Error() string {
  return "Error (" + err.Code + "): " + err.Message
}
