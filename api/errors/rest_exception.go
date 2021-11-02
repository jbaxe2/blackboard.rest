package errors

import (
  "net/url"
  "strconv"
)

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

  error
}

/**
 * The [_RestException] type implements the REST Exception interface.
 */
type _RestException struct {
  status int

  code string

  message string

  developerMessage string

  extraInfo *url.URL

  RestException
}

/**
 * The [NewRestException] function creates a new REST Exception instance.
 */
func NewRestException (
  status int, code, message, developerMessage string, extraInfo *url.URL,
) RestException {
  if 400 > status || 499 < status || "" == message {
    return nil
  }

  return &_RestException {
    status: status,
    code: code,
    message: message,
    developerMessage: developerMessage,
    extraInfo: extraInfo,
  }
}

func (exception *_RestException) Status() int {
  return exception.status
}

func (exception *_RestException) Code() string {
  return exception.code
}

func (exception *_RestException) Message() string {
  return exception.message
}

func (exception *_RestException) DeveloperMessage() string {
  return exception.developerMessage
}

func (exception *_RestException) ExtraInfo() *url.URL {
  return exception.extraInfo
}

func (exception *_RestException) Error() string {
  return "(" + strconv.Itoa (exception.status) + ") " + exception.message
}
