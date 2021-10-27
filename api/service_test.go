package api_test

import (
  "testing"

  "github.com/jbaxe2/blackboard.rest/api"
  "github.com/jbaxe2/blackboard.rest/oauth2"
)

/**
 * The [TestCreateNewService] function...
 */
func TestCreateNewService (t *testing.T) {
  println ("Create a new service instance.")

  if nil == api.NewService ("localhost", mockToken) {
    t.Error ("Creating a new service instance should not result in nil reference.")
    t.FailNow()
  }
}

/**
 * The [TestNewServiceRequiresHost] function...
 */
func TestNewServiceRequiresHost (t *testing.T) {
  println ("Creating a new service instance requires a host.")

  if nil != api.NewService ("", mockToken) {
    t.Error ("Missing host should result in nil service reference.")
    t.FailNow()
  }
}

/**
 * The [TestNewServiceRequiresToken] function...
 */
func TestNewServiceRequiresToken (t *testing.T) {
  println ("Creating a new service instance requires a token.")

  if nil != api.NewService ("localhost", nil) {
    t.Error ("Missing token should result in nil service reference.")
    t.FailNow()
  }
}

/**
 * The [TestNewServiceHasPertinentInformation] function...
 */
func TestNewServiceHasPertinentInformation (t *testing.T) {
  println ("New service instance retains the information used to create it.")

  host := "localhost"
  service := api.NewService (host, mockToken)

  if !(service.Host() == host && service.Token() == mockToken) {
    t.Error ("New service instance should retain the info used to create it.")
    t.FailNow()
  }
}

/**
 * Mocked instances to run the above tests with.
 */
var mockToken = oauth2.NewToken (
  "access_token", "token_type", "refresh_token", "scope", "user_id", 3600,
)
