package services

import (
  "github.com/jbaxe2/blackboard.rest.go/src/_scaffolding/connector"
  "github.com/jbaxe2/blackboard.rest.go/src/oauth2"
)

/**
 * The [BbRestService] interface...
 */
type BbRestService interface {
  HandleError (err error)
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
func (service *BlackboardRestService) SetHost (host string) {
  service.Connector.SetHost (host)
}

/**
 * The [SetAccessToken] method...
 */
func (service *BlackboardRestService) SetAccessToken (token oauth2.AccessToken) {
  service.Connector.SetAccessToken (token)
}
