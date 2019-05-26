package services

import (
  "github.com/jbaxe2/blackboard.rest.go/src/_scaffolding/connector"
  error2 "github.com/jbaxe2/blackboard.rest.go/src/_scaffolding/error"
  "github.com/jbaxe2/blackboard.rest.go/src/oauth2"
)

/**
 * The [BbRestService] interface...
 */
type BbRestService interface {
  HandleError (err error) error2.RestError
}

/**
 * The [BlackboardRestService] type...
 */
type BlackboardRestService struct {
  Connector connector.BbRestConnector

  BbRestService
}

/**
 * The [SetHost] method...
 */
func (services *BlackboardRestService) SetHost (host string) {
  services.Connector.SetHost (host)
}

/**
 * The [SetAccessToken] method...
 */
func (services *BlackboardRestService) SetAccessToken (token oauth2.AccessToken) {
  services.Connector.SetAccessToken (token)
}
