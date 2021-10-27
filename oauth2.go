package blackboard_rest

import (
  "net/url"

  "github.com/jbaxe2/blackboard.rest/api/errors"
  "github.com/jbaxe2/blackboard.rest/oauth2"
)

/**
 * The [OAuth2] interface provides the base type for interacting with the REST
 * API's oauth service.
 */
type OAuth2 interface {
  AuthorizationCode (redirectUri url.URL, clientId string, scope string) string

  RequestToken (
    grantType string, code string, redirectUri url.URL,
  ) (oauth2.Token, errors.OAuth2Error)

  GetTokenInfo (accessToken string) (oauth2.TokenInfo, errors.OAuth2Error)
}

/**
 * The [_OAuth2] type implements the OAuth2 interface.
 */
type _OAuth2 struct {
  OAuth2
}

/**
 * The [NewOAuth2] function creates a new OAuth2 service instance.
 */
func NewOAuth2() OAuth2 {
  return new (_OAuth2)
}
