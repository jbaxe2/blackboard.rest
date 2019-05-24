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
 * The [Authorizer] interface...
 */
type Authorizer interface {}

/**
 * The [RestAuthorizer] type...
 */
type RestAuthorizer struct {
  host url.URL

  clientId, secret string

  Authorizer
}

/**
 * The [RestUserAuthorizer] type...
 */
type RestUserAuthorizer struct {
  host url.URL

  clientId, secret string

  Authorizer
}

/**
 * The [NewRestAuthorizer] function...
 */
func NewRestAuthorizer (
  host url.URL, clientId string, secret string,
) RestAuthorizer {
  return RestAuthorizer {
    host: host, clientId: clientId, secret: secret,
  }
}

/**
 * The [NewRestUserAuthorizer] function...
 */
func NewRestUserAuthorizer (
  host url.URL, clientId string, secret string,
) RestUserAuthorizer {
  return RestUserAuthorizer {
    host: host, clientId: clientId, secret: secret,
  }
}

/**
 * The [RequestAuthorization] method...
 */
func (authorizer RestAuthorizer) RequestAuthorization() (AccessToken, error) {
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
 * The [RequestAuthorizationCode] method...
 */
func (authorizer *RestUserAuthorizer) RequestAuthorizationCode (
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

  return response.Body.Close()
}

/**
 * The [RequestUserAuthorization] method...
 */
func (authorizer *RestUserAuthorizer) RequestUserAuthorization (
  authCode string, redirectUri string,
) (AccessToken, error) {
  var accessToken AccessToken
  var err error
  var encodedRedirect string
  var parsedRedirect *url.URL

  if "" == redirectUri {
    encodedRedirect = ""
  } else {
    parsedRedirect, err = url.Parse (redirectUri)

    if nil != err {
      return accessToken, err
    }

    encodedRedirect = "&redirect_uri=" + parsedRedirect.String()
  }

  authCodeUriStr := authorizer.host.String() + config.Base +
    config.OAuth2Endpoints["authorization_code"] + "?code=" + authCode +
    encodedRedirect

  request := new (http.Request)
  request.SetBasicAuth (authorizer.clientId, authorizer.secret)
  request.URL, err = url.Parse (authCodeUriStr)

  response, err  := (new (http.Client)).Do (request)

  if nil != err {
    return accessToken, err
  }

  accessToken, err = _parseResponse (response)

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
