package blackboard_rest

import (
  "github.com/jbaxe2/blackboard.rest/terms"
)

/**
 * The [Terms] interface...
 */
type Terms interface {
  GetTerms() ([]terms.Term, error)

  CreateTerm (term terms.Term) error

  GetTerm (termId string) (terms.Term, error)

  UpdateTerm (termId string, term terms.Term) error
}
