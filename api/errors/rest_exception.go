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

/**
 * The [NewRestExceptionFromRaw] function creates a new REST Exception based on
 * a raw map (most likely as a response decoded from a JSON structure).
 */
func NewRestExceptionFromRaw (rawException map[string]interface{}) RestException {
  if nil == rawException {
    return nil
  }

  statusStr, haveStatus := rawException["status"].(string)
  code, _ := rawException["code"].(string)
  message, _ := rawException["message"].(string)
  developerMsg, _ := rawException["developerMessage"].(string)
  extraInfo, haveExtraInfo := rawException["extraInfo"].(string)

  var status int
  var extraInfoUri *url.URL

  if haveStatus {
    status, _ = strconv.Atoi (statusStr)
  }

  if haveExtraInfo {
    extraInfoUri, _ = url.Parse (extraInfo)
  }

  return NewRestException (status, code, message, developerMsg, extraInfoUri)
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
