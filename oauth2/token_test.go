package oauth2_test

import (
  "testing"

  "github.com/jbaxe2/blackboard.rest/oauth2"
)

/**
 * The [TestCreateNewOAuth2Token] function...
 */
func TestCreateNewOAuth2Token (t *testing.T) {
  println ("Create a new OAuth2 token instance.")

  if nil == oauth2.NewToken (
    "access_token", "token_type", "refresh_token", "scope", "user_id", 3600,
  ) {
    t.Error ("Creating a new OAuth2 token instance should not be nil.")
    t.FailNow()
  }
}

/**
 * The [TestNewOAuth2TokenRequiresAccessToken] function...
 */
func TestNewOAuth2TokenRequiresAccessToken (t *testing.T) {
  println ("Creating OAuth2 token requires access token.")

  if nil != oauth2.NewToken ("", "token_type", "refresh_token", "", "", 0) {
    t.Error ("Missing access token should result in a nil reference.")
    t.FailNow()
  }
}

/**
 * The [TestNewOAuth2TokenRequiresTokenType] function...
 */
func TestNewOAuth2TokenRequiresTokenType (t *testing.T) {
  println ("Creating OAuth2 token requires token type.")

  if nil != oauth2.NewToken ("access_token", "", "refresh_token", "", "", 0) {
    t.Error ("Missing token type should result in a nil reference.")
    t.FailNow()
  }
}

/**
 * The [TestNewOAuth2TokenCanHaveEmptyRefreshTokenScopeNotOffline] function...
 */
func TestNewOAuth2TokenCanHaveEmptyRefreshTokenScopeNotOffline (t *testing.T) {
  println (
    "Creating OAuth2 token can have empty refresh token if scope is not offline.",
  )

  if nil == oauth2.NewToken ("access_token", "token_type", "", "", "", 0) {
    t.Error ("Missing refresh token should not result in a nil reference.")
    t.FailNow()
  }
}
