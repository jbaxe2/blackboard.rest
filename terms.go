package blackboard_rest

import (
  "github.com/jbaxe2/blackboard.rest/api"
  "github.com/jbaxe2/blackboard.rest/terms"
)

/**
 * The [Terms] interface provides the base interface for interacting with the
 * REST API's terms service.
 */
type Terms interface {
  GetTerms() ([]terms.Term, error)

  CreateTerm (term terms.Term) error

  GetTerm (termId string) (terms.Term, error)

  UpdateTerm (termId string, term terms.Term) error
}

/**
 * The [_Terms] type implements the Terms interface.
 */
type _Terms struct {
  service api.Service

  Terms
}

/**
 * The [NewTerms] function creates a new terms insance.
 */
func NewTerms (service api.Service) Terms {
  if nil == service {
    return nil
  }

  return &_Terms {
    service: service,
  }
}
