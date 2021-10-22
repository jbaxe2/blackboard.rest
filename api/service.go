package api

import "github.com/jbaxe2/blackboard.rest/oauth2"

/**
 * The [Service] interface is the base interface for all Blackboard Learn REST
 * API service types.
 */
type Service interface {
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

  Service
}

/**
 * The [NewService] function creates a new Service instance.
 */
func NewService (host string, token oauth2.Token) Service {
  if "" == host {
    return nil
  }

  return new (_Service)
}
