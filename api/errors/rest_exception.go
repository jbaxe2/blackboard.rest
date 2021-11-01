package errors

import "net/url"

/**
 * The [RestException] interface provides the base type for Blackboard's REST
 * API errors.
 */
type RestException interface {
  Status() int

  Code() string

  Message() string

  DeveloperMessage() string

  ExtraInfo() *url.URL
}

/**
 * The [_RestException] type implements the REST Exception interface.
 */
type _RestException struct {
  RestException
}

/**
 * The [NewRestException] function creates a new REST Exception instance.
 */
func NewRestException (
  status int, code, message, developerMessage string, extraInfo *url.URL,
) RestException {
  if 400 > status || 499 < status {
    return nil
  }

  return new (_RestException)
}
