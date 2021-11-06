package api

import (
  "errors"
  "net/http"

  "github.com/jbaxe2/blackboard.rest/oauth2"
  "github.com/jbaxe2/blackboard.rest/utils"
)

/**
 * The [Service] interface is the base interface for all Blackboard Learn REST
 * API service types.
 */
type Service interface {
  Host() string

  Token() oauth2.Token

  Request (
    endpoint string, method string, data map[string]interface{},
    options map[string]interface{}, useVersion int,
  ) (interface{}, error)
}

/**
 * The [_Service] type implements the Service interface.
 */
type _Service struct {
  host string

  token oauth2.Token

  roundTripper http.RoundTripper

  Service
}

/**
 * The [NewService] function creates a new Service instance.
 */
func NewService (
  host string, token oauth2.Token, roundTripper http.RoundTripper,
) Service {
  if "" == host || nil == token || "" == token.AccessToken() ||
     1 > token.ExpiresIn() {
    return nil
  }

  if nil == roundTripper {
    roundTripper = http.DefaultTransport
  }

  return &_Service {
    host: host,
    token: token,
    roundTripper: roundTripper,
  }
}

func (service *_Service) Host() string {
  return service.host
}

func (service *_Service) Token() oauth2.Token {
  return service.token
}

/**
 * The [Request] method makes the request to the REST API.
 */
func (service *_Service) Request (
  endpoint string, method string, data map[string]interface{},
  options map[string]interface{}, useVersion int,
) (interface{}, error) {
  if err := _verifyRequestConditions (endpoint, method); nil != err {
    return nil, err
  }

  return nil, nil
}

/**
 * The [_verifyRequestConditions] function verifies the information used to
 * create a service request is appropriate.
 */
func _verifyRequestConditions (endpoint, method string) error {
  if "" == endpoint {
    return errors.New ("missing service endpoint")
  }

  methods := []string {"GET", "POST", "PUT", "PATCH", "DELETE"}

  if !utils.StringInStrings (method, methods) {
    return errors.New ("inappropriate HTTP method")
  }

  return nil
}
