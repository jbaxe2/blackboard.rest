package blackboard_rest_test

import (
  "net/http"
  "testing"

  blackboardRest "github.com/jbaxe2/blackboard.rest"
  "github.com/jbaxe2/blackboard.rest/api"
)

/**
 * The [TestCreateNewUsers] function...
 */
func TestCreateNewUsers (t *testing.T) {
  println ("Create a new users service instance.")

  if nil == blackboardRest.NewUsers (mockUsersService) {
    t.Error ("Creating a new users service instance should not be nil reference.")
  }
}

/**
 * The [TestNewUsersRequiresService] function...
 */
func TestNewUsersRequiresService (t *testing.T) {
  println ("Creating a new users instance requires a service instance.")

  if nil != blackboardRest.NewUsers (nil) {
    t.Error ("Missing service instance should result in nil reference.")
  }
}

/**
 * Mocked types and instances to run the above tests with.
 */
var mockUsersService = api.NewService ("localhost", mockToken, mockUsersRoundTripper)

var mockUsersRoundTripper = new (_MockUsersRoundTripper)

type _MockUsersRoundTripper struct {
  http.RoundTripper
}
