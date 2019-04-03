package services

import (
  "github.com/jbaxe2/blackboard.rest.go/src/_scaffolding/connector"
  error2 "github.com/jbaxe2/blackboard.rest.go/src/_scaffolding/error"
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
