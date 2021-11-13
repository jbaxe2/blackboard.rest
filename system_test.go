package blackboard_rest_test

import (
  "net/http"
  "testing"

  blackboardRest "github.com/jbaxe2/blackboard.rest"
  "github.com/jbaxe2/blackboard.rest/api"
)

/**
 * The [TestCreateNewSystem] function...
 */
func TestCreateNewSystem (t *testing.T) {
  println ("Create a new system instance.")

  if nil == blackboardRest.NewSystem (mockSystemService) {
    t.Error ("Creating a new system instance should not be a nil reference.")
  }
}

/**
 * The [TestNewSystemRequiresService] function...
 */
func TestNewSystemRequiresService (t *testing.T) {
  println ("Creating new system instance requires a service instance.")

  if nil != blackboardRest.NewSystem (nil) {
    t.Error ("Missing service instance should result in nil reference.")
  }
}

/**
 * Mocked types and instances to run the above tests with.
 */
var mockSystemService =
  api.NewService ("localhost", mockToken, mockSystemRoundTripper)

var mockSystemRoundTripper = new (_MockSystemRoundTripper)

type _MockSystemRoundTripper struct {
  http.RoundTripper
}
