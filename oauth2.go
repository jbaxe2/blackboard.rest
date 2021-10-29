package blackboard_rest

import (
  "net/http"
  "net/url"
  "strings"

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
    request *http.Request, redirectUri *url.URL, clientId string, scope string,
  ) *http.Response

  RequestToken (
    grantType string, code string, redirectUri *url.URL,
  ) (oauth2.Token, errors.OAuth2Error)

  GetTokenInfo (accessToken string) (oauth2.TokenInfo, errors.OAuth2Error)
}

/**
 * The [_OAuth2] type implements the OAuth2 interface.
 */
type _OAuth2 struct {
  host string

  clientId string

  secret string

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
 * The [SetClientIdAndSecret] method sets the client ID and password for the
 * basic authentication method for the HTTP request when requesting an OAuth2
 * authorization token.
 */
func (oAuth2 *_OAuth2) SetClientIdAndSecret (clientId, secret string) {
  oAuth2.clientId = clientId
  oAuth2.secret = secret
}

/**
 * The [AuthorizationCode] method is used to retrieve an authorization code that
 * can be used to generate an authorization token.
 */
func (oAuth2 *_OAuth2) AuthorizationCode (
  request *http.Request, redirectUri *url.URL, clientId string, scope string,
) *http.Response {
  endpoint :=
    strings.Replace (api.Base, "{v}", "1", 1) + string (api.AuthorizationCode)

  authorizedUri := "https://" + oAuth2.host + endpoint + "?redirect_uri=" +
    redirectUri.String() + "&client_id=" + clientId + "&response_type=code&scope=" +
    scope

  response, _ := oAuth2.roundTripper.RoundTrip (request)
  response.StatusCode = 200
  response.Header.Add ("Location", authorizedUri)

  return response
}

/**
 * The [RequestToken] method is used to retrieve an OAuth2 authorization token
 * that will be used by other services.
 */
func (oAuth2 *_OAuth2) RequestToken (
  grantType string, code string, redirectUri *url.URL,
) (oauth2.Token, errors.OAuth2Error) {

  client := http.Client {
    Transport: oAuth2.roundTripper,
  }

  tokenUri := oAuth2._buildRequestTokenUri (grantType, code, redirectUri)

  request, _ := http.NewRequest (http.MethodPost, tokenUri, nil)
  request.SetBasicAuth (oAuth2.clientId, oAuth2.secret)
  request.Header.Set ("Content-Type", "application/x-www-form-urlencoded")

  response, _ := client.Do (request)

  return oAuth2._parseTokenRequestResponse (response)
}

/**
 * The [_buildRequestTokenUri] method builds the URI that will be used when
 * requesting an OAuth2 authorization token.
 */
func (oAuth2 *_OAuth2) _buildRequestTokenUri (
  grantType string, code string, redirectUri *url.URL,
) string {
  endpoint := strings.Replace (api.Base, "{v}", "1", 1) + string (api.RequestToken)

  tokenUri := "https://" + oAuth2.host + endpoint + "?grant_type=" + grantType +
    "&code=" + code

  if "" != redirectUri.String() {
    tokenUri += "&redirect_uri=" + redirectUri.String()
  }

  return tokenUri
}

/**
 * The [_parseTokenRequestResponse] method parses the response provided by the
 * REST API call requesting an OAuth2 token.
 */
func (oAuth2 *_OAuth2) _parseTokenRequestResponse (
  response *http.Response,
) (oauth2.Token, errors.OAuth2Error) {
  return nil, nil
}
