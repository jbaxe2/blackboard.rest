package blackboard_rest_test

import (
  "testing"

  blackboardRest "github.com/jbaxe2/blackboard.rest"
)

/**
 * The [CreateNewOAuth2Instance] function...
 */
func CreateNewOAuth2Instance (t *testing.T) {
  println ("Create a new OAuth2 service instance.")

  if nil == blackboardRest.NewOAuth2() {
    t.Error ("New OAuth2 instance should not be a nil reference.")
    t.FailNow()
  }
}
