package blackboard_rest

import (
  "encoding/json"
  "io/ioutil"
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
  SetClientIdAndSecret (string, string)

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
  if "" == host {
    return nil
  }

  if nil == roundTripper {
    roundTripper = http.DefaultTransport
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
 * can be used to generate an authorization token.  This method should be called
 * when using the REST API's via 3-legged OAuth.
 */
func (oAuth2 *_OAuth2) AuthorizationCode (
  request *http.Request, redirectUri *url.URL, clientId string, scope string,
) *http.Response {
  endpoint :=
    strings.Replace (api.Base, "{v}", "1", 1) + string (api.AuthorizationCode)

  if "" == scope {
    scope = "read"
  }

  authorizedUri := "https://" + oAuth2.host + endpoint + "?redirect_uri=" +
    url.QueryEscape (redirectUri.String()) + "&client_id=" + clientId +
    "&response_type=code&scope=" + scope

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
  if err := _verifyRequestTokenRequest (grantType, code); nil != err {
    return nil, err
  }

  tokenUri := _buildRequestTokenUri (oAuth2.host, grantType, code, redirectUri)
  request, err := http.NewRequest (http.MethodPost, tokenUri, nil)

  if nil != err {
    return nil, errors.NewOAuth2Error ("invalid_request", err.Error())
  }

  request.SetBasicAuth (oAuth2.clientId, oAuth2.secret)
  request.Header.Set ("Content-Type", "application/x-www-form-urlencoded")

  client := http.Client {Transport: oAuth2.roundTripper}
  response, _ := client.Do (request)

  return _parseTokenRequestResponse (response)
}

/**
 * The [_verifyRequestTokenRequest] function verifies the information used to
 * create a token request is as it should be; otherwise an OAuth2 error will be
 * returned.
 */
func _verifyRequestTokenRequest (grantType, code string) errors.OAuth2Error {
  if "authorization_code" == grantType && "" == code {
    return errors.NewOAuth2Error ("unauthorized_client", "Missing authorization code.")
  }

  if !("client_credentials" == grantType || "authorization_code" == grantType ||
       "refresh_token" == grantType) {
    return errors.NewOAuth2Error ("unauthorized_client", "Missing authorization code.")
  }

  return nil
}

/**
 * The [_buildRequestTokenUri] function builds the URI that will be used when
 * requesting an OAuth2 authorization token.
 */
func _buildRequestTokenUri (
  host, grantType, code string, redirectUri *url.URL,
) string {
  endpoint := strings.Replace (api.Base, "{v}", "1", 1) + string (api.RequestToken)

  tokenUri := "https://" + host + endpoint + "?grant_type=" + grantType

  if "authorization_code" == grantType {
    tokenUri += "&code=" + code
  }

  if nil != redirectUri {
    tokenUri += "&redirect_uri=" + url.QueryEscape (redirectUri.String())
  }

  return tokenUri
}

/**
 * The [_parseTokenRequestResponse] function parses the response provided by the
 * REST API call requesting an OAuth2 token.
 */
func _parseTokenRequestResponse (
  response *http.Response,
) (oauth2.Token, errors.OAuth2Error) {
  defer response.Body.Close()

  var rawResponse map[string]interface{}
  responseBytes, _ := ioutil.ReadAll (response.Body)

  if err := json.Unmarshal (responseBytes, &rawResponse); nil != err {
    return nil, errors.NewOAuth2Error ("invalid_response", err.Error())
  }

  if _, wasError := rawResponse["error"]; wasError {
    return nil, errors.NewOAuth2Error (
      rawResponse["error"].(string), rawResponse["error_description"].(string),
    )
  }

  if _, haveAccessToken := rawResponse["access_token"]; !haveAccessToken {
    return nil, errors.NewOAuth2Error ("missing_token", "Missing access token.")
  }

  userId := ""
  scope := ""
  refreshToken := ""

  if _, haveUserId := rawResponse["user_id"]; haveUserId {
    userId = rawResponse["user_id"].(string)
  }

  if _, haveScope := rawResponse["scope"]; haveScope {
    scope = rawResponse["scope"].(string)
  }

  if _, haveRefresh := rawResponse["refresh_token"]; haveRefresh {
    refreshToken = rawResponse["refresh_token"].(string)
  }

  return oauth2.NewToken (
    rawResponse["access_token"].(string), rawResponse["token_type"].(string),
    refreshToken, scope, userId, int32 (rawResponse["expires_in"].(float64)),
  ), nil
}
