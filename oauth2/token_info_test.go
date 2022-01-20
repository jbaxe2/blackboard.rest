package oauth2_test

import (
  "testing"

  "github.com/jbaxe2/blackboard.rest/oauth2"
)

/**
 * The [TestCreateNewTokenInfo] function...
 */
func TestCreateNewTokenInfo (t *testing.T) {
  println ("Create a new token info instance.")

  if nil == oauth2.NewTokenInfo ("applicationId", "read") {
    t.Error ("Creating a new token info instance should not be nil.")
  }
}

/**
 * The [TestNewTokenInfoRequiresApplicationId] function...
 */
func TestNewTokenInfoRequiresApplicationId (t *testing.T) {
  println ("New token info instance requires an application ID.")

  if nil != oauth2.NewTokenInfo ("", "read") {
    t.Error ("Missing application ID should result in nil reference.")
  }
}

/**
 * The [TestNewTokenInfoRequiresScope] function...
 */
func TestNewTokenInfoRequiresScope (t *testing.T) {
  println ("New token info instance requires a scope.")

  if nil != oauth2.NewTokenInfo ("applicationId", "") {
    t.Error ("Missing scope should result in nil reference.")
  }
}

/**
 * The [TestNewTokenInfoHasPertinentInformation] function...
 */
func TestNewTokenInfoHasPertinentInformation (t *testing.T) {
  println ("New token info retains the information used to create it.")

  applicationId := "applicationId"
  scope := "read"

  tokenInfo := oauth2.NewTokenInfo (applicationId, scope)

  if !(tokenInfo.ApplicationId() == applicationId && tokenInfo.Scope() == scope) {
    t.Error ("New token info should retain the info used to create it.")
  }
}
