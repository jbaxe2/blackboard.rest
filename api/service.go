package api

import (
  "net/http"

  "github.com/jbaxe2/blackboard.rest/oauth2"
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
