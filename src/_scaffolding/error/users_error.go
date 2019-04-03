package error

/**
 * The [UsersError] type...
 */
type UsersError struct {
  RestableError

  RestError
}

/**
 * The [Error] method...
 */
func (usersError UsersError) Error() string {
  return usersError.RestError.Error()
}
