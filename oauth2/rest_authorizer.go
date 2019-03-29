package oauth2

import (
  "encoding/json"
  "github.com/jbaxe2/blackboard.rest.go/config"
  "net/http"
  "net/url"
  "strconv"
)

/**
 * The [RestAuthorizer] interface...
 */
type RestAuthorizer interface {
  RequestAuthorization() (AccessToken, error)
}

/**
 * The [RestUserAuthorizer] interface...
 */
type RestUserAuthorizer interface {
  RequestAuthorizationCode (redirectUri string, response http.Response)

  RequestUserAuthorization (authCode string, redirectUri string) (AccessToken, error)

  RestAuthorizer
}

/**
 * The [AuthorizerFactory] type...
 */
type AuthorizerFactory struct {}

/**
 * The [_RestAuthorizer] type...
 */
type _RestAuthorizer struct {
  host url.URL

  clientId, secret string

  RestAuthorizer
}

/**
 * The [_RestUserAuthorizer] type...
 */
type _RestUserAuthorizer struct {
  _RestAuthorizer

  RestUserAuthorizer
}

/**
 * The [BuildAuthorizer] method...
 */
func (*AuthorizerFactory) BuildAuthorizer (
  host url.URL, clientId string, secret string, authType string,
) RestAuthorizer {
  var restAuthorizer RestAuthorizer

  if "user" == authType {
    restAuthorizer = new (_RestAuthorizer)
    restAuthorizer = restAuthorizer.(RestUserAuthorizer)
  } else {
    restAuthorizer = new (_RestAuthorizer)
  }

  return restAuthorizer
}

/**
 * The [RequestAuthorization] method...
 */
func (authorizer *_RestAuthorizer) RequestAuthorization() (AccessToken, error) {
  request := new (http.Request)
  request.SetBasicAuth (authorizer.clientId, authorizer.secret)

  response, err := (new (http.Client)).Do (request)
  accessToken, err := _parseResponse (response)

  err = response.Body.Close()

  return accessToken, err
}

/**
 * The [RequestAuthorizationCode] method...
 */
func (authorizer *_RestUserAuthorizer) RequestAuthorizationCode (
  redirectUri string, response *http.Response,
) error {
  var err error
  var encoded *url.URL

  encoded, err = url.Parse (redirectUri)

  if nil != err {
    return err
  }

  authorizeUriStr := authorizer.host.String() + config.Base +
    config.OAuth2Endpoints()["authorization_code"] + "?redirect_uri=" +
    encoded.String() + "&client_id=" + authorizer.clientId +
    "&response_type=code&scope=read"

  response.Header.Add ("Location", authorizeUriStr)
  err = response.Body.Close()

  return err
}

/**
 * The [RequestUserAuthorization] method...
 */
func (authorizer *_RestUserAuthorizer) RequestUserAuthorization (
  authCode string, redirectUri string,
) (AccessToken, error) {
  var accessToken AccessToken
  var err error

  return accessToken, err
}

/**
 * The [_parseResponse] function...
 */
func _parseResponse (response *http.Response) (AccessToken, error) {
  var accessToken AccessToken
  var err error
  var responseMap = make (map[string]string)
  var expires int

  err = json.NewDecoder (response.Body).Decode (responseMap)
  expires, err = strconv.Atoi (responseMap["expires_in"])

  accessToken = AccessToken {
    responseMap["access_token"], responseMap["token_type"],
    responseMap["refresh_token"], responseMap["scope"],
    responseMap["user_id"], expires,
  }

  return accessToken, err
}
