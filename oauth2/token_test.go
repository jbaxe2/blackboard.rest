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
  }
}

/**
 * The [TestNewOAuth2TokenRequiresAccessToken] function...
 */
func TestNewOAuth2TokenRequiresAccessToken (t *testing.T) {
  println ("Creating OAuth2 token requires access token.")

  if nil != oauth2.NewToken ("", "token_type", "refresh_token", "", "user_id", 3600) {
    t.Error ("Missing access token should result in a nil reference.")
  }
}

/**
 * The [TestNewOAuth2TokenRequiresTokenType] function...
 */
func TestNewOAuth2TokenRequiresTokenType (t *testing.T) {
  println ("Creating OAuth2 token requires token type.")

  if nil != oauth2.NewToken ("access_token", "", "refresh_token", "", "user_id", 3600) {
    t.Error ("Missing token type should result in a nil reference.")
  }
}

/**
 * The [TestNewOAuth2TokenCanHaveEmptyRefreshTokenScopeNotOffline] function...
 */
func TestNewOAuth2TokenCanHaveEmptyRefreshTokenScopeNotOffline (t *testing.T) {
  println (
    "Creating OAuth2 token can have empty refresh token if scope is not offline.",
  )

  if nil == oauth2.NewToken ("access_token", "token_type", "", "", "user_id", 3600) {
    t.Error ("Missing refresh token should not result in a nil reference.")
  }
}

/**
 * The [TestNewOAuth2TokenRefreshTokenNotEmptyIfOfflineScope] function...
 */
func TestNewOAuth2TokenRefreshTokenNotEmptyIfOfflineScope (t *testing.T) {
  println (
    "Creating OAuth2 token cannot have empty refresh token if scope is offline.",
  )

  if nil != oauth2.NewToken ("access_token", "token_type", "", "offline", "user_id", 3600) {
    t.Error (
      "Missing refresh token for offline scope should result in nil reference.",
    )
  }
}

/**
 * The [TestNewOAuth2TokenRequiresUserId] function...
 */
func TestNewOAuth2TokenRequiresUserId (t *testing.T) {
  println ("Creating OAuth2 token requires a user ID.")

  if nil != oauth2.NewToken ("access_token", "token_type", "", "", "", 3600) {
    t.Error ("Missing user ID should result in nil reference.")
  }
}

/**
 * The [TestNewOAuth2TokenRequiresExpiresInGreaterThanZero] function...
 */
func TestNewOAuth2TokenRequiresExpiresInGreaterThanZero (t *testing.T) {
  println ("Creating OAuth2 token requires expires in value greater than 0.")

  if nil != oauth2.NewToken (
    "access_token", "token_type", "", "read", "user_id", 0,
  ) {
    t.Error ("Expires in value less than 1 should result in nil reference.")
  }
}

/**
 * The [TestNewOAuth2TokenHasPertinentInformation] function...
 */
func TestNewOAuth2TokenHasPertinentInformation (t *testing.T) {
  println ("New OAuth2 token retains the information used to create it.")

  accessToken := "access_token"
  tokenType := "token_type"
  refreshToken := "refresh_token"
  scope := "offline"
  userId := "user_id"
  expiresIn := int32 (3600)

  token :=
    oauth2.NewToken (accessToken, tokenType, refreshToken, scope, userId, expiresIn)

  if !(token.AccessToken() == accessToken && token.TokenType() == tokenType &&
       token.RefreshToken() == refreshToken && token.Scope() == scope &&
       token.UserId() == userId && token.ExpiresIn() == expiresIn) {
    t.Error ("New token should retain the information used to create it.")
    t.FailNow()
  }
}
