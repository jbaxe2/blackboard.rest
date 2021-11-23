package blackboard_rest_test

import (
  "io/ioutil"
  "net/http"
  "strings"
)

/**
 * A default improper request message, if none are passed as part of the
 * conditions and response bodies.
 */
const improperRequest = `{"status":400,"message":"Improper request"}`

/**
 * The [ResponseBuilder] interface provides the base type for building out an
 * HTTP Response from a provided HTTP Request.
 */
type ResponseBuilder interface {
  Build (*http.Request) *http.Response
}

/**
 * The [_ResponseBuilder] type implements the Response Builder interface.
 */
type _ResponseBuilder struct {
  conditions []bool

  responseBodies []string

  ResponseBuilder
}

/**
 * The [NewResponseBuilder] function creates a new response builder instance.
 * It requires a slice of ordered conditions and associated response bodies.
 * The order of the conditions should go from more specific to generic, as the
 * response will be built from the first condition satisfied.
 */
func NewResponseBuilder (conditions []bool, responseBodies []string) ResponseBuilder {
  if len (conditions) != len (responseBodies) {
    return nil
  }

  return &_ResponseBuilder {
    conditions: conditions,
    responseBodies: responseBodies,
  }
}

/**
 * The [Build] method builds an HTTP Response from a provided HTTP Request along
 * with the conditions and response bodies.
 */
func (builder *_ResponseBuilder) Build (request *http.Request) *http.Response {
  request.Response = &http.Response {
    Request: request,
    Header: make (http.Header),
  }

  request.Response.StatusCode = 200

  builder._buildResponseBody (request.Response)

  return request.Response
}

/**
 * The [_buildResponseBody] method builds out the response body text.
 */
func (builder *_ResponseBuilder) _buildResponseBody (response *http.Response) {
  responseBody := ""

  for i, condition := range builder.conditions {
    if condition {
      responseBody = builder.responseBodies[i]
      break
    }
  }

  if "" == responseBody {
    responseBody = improperRequest
    response.StatusCode = 400
  }

  response.Body = ioutil.NopCloser (strings.NewReader (responseBody))
}
