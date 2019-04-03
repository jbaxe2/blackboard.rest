package services

import (
  "github.com/jbaxe2/blackboard.rest.go/src/_scaffolding/connector"
  error2 "github.com/jbaxe2/blackboard.rest.go/src/_scaffolding/error"
  "github.com/jbaxe2/blackboard.rest.go/src/oauth2"
)

/**
 * The [BbRestServices] interface...
 */
type BbRestServices interface {
  HandleError (err error) error2.RestError
}

/**
 * The [BlackboardRestServices] type...
 */
type BlackboardRestServices struct {
  Connector connector.BbRestConnector

  BbRestServices
}

/**
 * The [SetAccessToken] method...
 */
func (services *BlackboardRestServices) SetAccessToken (token oauth2.AccessToken) {
  services.Connector.SetAccessToken (token)
}
