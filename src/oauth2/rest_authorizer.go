package oauth2

import (
  "encoding/json"
  "github.com/jbaxe2/blackboard.rest.go/src/_scaffolding/config"
  "io/ioutil"
  "net/http"
  "net/url"
  "strings"
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

  restAuthorizer = &_RestAuthorizer {
    host: host, clientId: clientId, secret: secret,
  }

  if "user" == authType {
    restAuthorizer, _ = restAuthorizer.(_RestUserAuthorizer)
  }

  return restAuthorizer
}

/**
 * The [RequestAuthorization] method...
 */
func (authorizer _RestAuthorizer) RequestAuthorization() (AccessToken, error) {
  var accessToken AccessToken
  var err error
  var response *http.Response

  request := new (http.Request)

  request.URL, err = url.Parse (
    authorizer.host.String() + config.Base +
    config.OAuth2Endpoints["request_token"],
  )

  if nil != err {
    return accessToken, err
  }

  request.Header = make (http.Header)
  request.Header.Set ("Content-Type", "application/x-www-form-urlencoded")

  request.Method = "POST"
  request.SetBasicAuth (authorizer.clientId, authorizer.secret)

  request.Body = ioutil.NopCloser (
    strings.NewReader ("grant_type=client_credentials"),
  )

  response, err = (new (http.Client)).Do (request)

  if nil != err {
    return accessToken, err
  }

  accessToken, err = _parseResponse (response)

  err = response.Body.Close()

  return accessToken, err
}

/**
 * The [RequestAuthorization] method...
 */
func (authorizer _RestUserAuthorizer) RequestAuthorization() (AccessToken, error) {
  return authorizer._RestAuthorizer.RequestAuthorization()
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
    config.OAuth2Endpoints["authorization_code"] + "?redirect_uri=" +
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
  var encodedRedirect string
  var parsedRedirect *url.URL
println ("in request user authorization, with auth code")
  if "" == redirectUri {
    encodedRedirect = ""
  } else {
    parsedRedirect, err = url.Parse (redirectUri)

    if nil != err {
      return accessToken, err
    }

    encodedRedirect = "&redirect_uri=" + parsedRedirect.String()
  }
println ("before creating the auth code uri string")
  authCodeUriStr := authorizer.host.String() + config.Base +
    config.OAuth2Endpoints["authorization_code"] + "?code=" + authCode +
    encodedRedirect
println ("before creating the new request and setting the basic auth")
  request := new (http.Request)
  request.SetBasicAuth (authorizer.clientId, authorizer.secret)
  request.URL, err = url.Parse (authCodeUriStr)
println ("before creating the new client request")
  response, err  := (new (http.Client)).Do (request)

  if nil != err {
    return accessToken, err
  }
println ("before parsing the response")
  accessToken, err = _parseResponse (response)
println ("parsed response, closing the response body")
  err = response.Body.Close()

  return accessToken, err
}

/**
 * The [_parseResponse] function...
 */
func _parseResponse (response *http.Response) (AccessToken, error) {
  var accessToken AccessToken
  var err error
  var responseBytes []byte
  var parsedResponse map[string]interface{}

  responseBytes, err = ioutil.ReadAll (response.Body)

  if nil != err {
    return accessToken, err
  }

  err = json.Unmarshal (responseBytes, &parsedResponse)

  accessToken = AccessToken {
    accessToken: parsedResponse["access_token"].(string),
    tokenType:   parsedResponse["token_type"].(string),
    expiresIn:   parsedResponse["expires_in"].(float64),
  }

  if userId, ok := parsedResponse["user_id"]; ok {
    accessToken.userId = userId.(string)
  }

  if scope, ok := parsedResponse["scope"]; ok {
    accessToken.scope = scope.(string)
  }

  return accessToken, err
}
