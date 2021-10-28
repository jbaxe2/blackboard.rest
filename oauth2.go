package blackboard_rest

import (
  "net/http"
  "net/url"

  "github.com/jbaxe2/blackboard.rest/api"
  "github.com/jbaxe2/blackboard.rest/api/errors"
  "github.com/jbaxe2/blackboard.rest/oauth2"
)

/**
 * The [OAuth2] interface provides the base type for interacting with the REST
 * API's oauth service.
 */
type OAuth2 interface {
  AuthorizationCode (
    request *http.Request, redirectUri url.URL, clientId string, scope string,
  ) *http.Response

  RequestToken (
    grantType string, code string, redirectUri url.URL,
  ) (oauth2.Token, errors.OAuth2Error)

  GetTokenInfo (accessToken string) (oauth2.TokenInfo, errors.OAuth2Error)
}

/**
 * The [_OAuth2] type implements the OAuth2 interface.
 */
type _OAuth2 struct {
  host string

  roundTripper http.RoundTripper

  OAuth2
}

/**
 * The [NewOAuth2] function creates a new OAuth2 service instance.
 */
func NewOAuth2 (host string, roundTripper http.RoundTripper) OAuth2 {
  if "" == host || nil == roundTripper {
    return nil
  }

  return &_OAuth2 {
    host: host,
    roundTripper: roundTripper,
  }
}

/**
 * The [AuthorizationCode] method is used to retrieve an authorization code that
 * can be used to generate an authorization token.
 */
func (oAuth2 *_OAuth2) AuthorizationCode (
  request *http.Request, redirectUri url.URL, clientId string, scope string,
) *http.Response {
  authorizedUri := request.URL.Scheme + "://" + oAuth2.host + "/" +
    api.OAuth2Endpoints["authorization_code"] + "?redirect_uri=" +
    redirectUri.String() + "&client_id=" + clientId + "&response_type=code&scope=" +
    scope

  response, _ := oAuth2.roundTripper.RoundTrip (request)
  response.Header.Add ("Location", authorizedUri)

  return response
}
