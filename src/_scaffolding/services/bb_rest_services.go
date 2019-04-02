package services

import "github.com/jbaxe2/blackboard.rest.go/src/_scaffolding/connector"

/**
 * The [BbRestServices] interface...
 */
type BbRestServices interface {}

/**
 * The [BlackboardRestServices] type...
 */
type BlackboardRestServices struct {
  connector connector.BbRestConnector

  BbRestServices
}

func (services *BlackboardRestServices) Connector() connector.BbRestConnector {
  return services.connector
}
