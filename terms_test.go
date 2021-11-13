package blackboard_rest_test

import (
  "net/http"
  "testing"

  blackboardRest "github.com/jbaxe2/blackboard.rest"
  "github.com/jbaxe2/blackboard.rest/api"
)

/**
 * The [TestCreateNewTerms] function...
 */
func TestCreateNewTerms (t *testing.T) {
  println ("Create a new terms instance.")

  if nil == blackboardRest.NewTerms (mockTermsService) {
    t.Error ("Creating a new terms instance should not be a nil reference.")
  }
}

/**
 * The [TestNewTermsRequiresService] function...
 */
func TestNewTermsRequiresService (t *testing.T) {
  println ("Creating new terms instance requires a service instance.")

  if nil != blackboardRest.NewTerms (nil) {
    t.Error ("Missing service instance should result in nil reference.")
  }
}

/**
 * Mocked types and instances to run the above tests with.
 */
var mockTermsService =
  api.NewService ("localhost", mockToken, mockTermsRoundTripper)

var mockTermsRoundTripper = new (_MockTermsRoundTripper)

type _MockTermsRoundTripper struct {
  http.RoundTripper
}
