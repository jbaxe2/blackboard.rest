package blackboard_rest

import (
  "github.com/jbaxe2/blackboard.rest.go/src/oauth2"
  "net/http"
  "net/url"
)

/**
 * The [BlackboardRestOAuth2] interface...
 */
type BlackboardRestOAuth2 interface {
  GetAuthorizationCode (
    redirectUri url.URL, responseType string, response http.Response,
  ) error

  RequestToken (
    grantType string, code string, redirectUri url.URL,
  ) (oauth2.AccessToken, error)
}

/**
 * The [_BbRestOAuth2] type...
 */
type _BbRestOAuth2 struct {
  host url.URL

  clientId, secret string

  authorizer oauth2.RestAuthorizer

  BlackboardRestOAuth2
}

func (restOAuth2 *_BbRestOAuth2) Host() url.URL {
  return restOAuth2.host
}

func (restOAuth2 *_BbRestOAuth2) ClientId() string {
  return restOAuth2.clientId
}

func (restOAuth2 *_BbRestOAuth2) Secret() string {
  return restOAuth2.secret
}

/**
 * The [GetOAuth2Instance] function...
 */
func GetOAuth2Instance (
  host url.URL, clientId string, secret string,
) BlackboardRestOAuth2 {
  return &_BbRestOAuth2 {
    host: host, clientId: clientId, secret: secret,
  }
}

/**
 * The [GetAuthorizationCode] method...
 */
func (restOAuth2 *_BbRestOAuth2) GetAuthorizationCode (
  redirectUri url.URL, responseType string, response http.Response,
) error {
  var err error

  restOAuth2._createAuthorizer ("authorization_code")

  restOAuth2.authorizer.(oauth2.RestUserAuthorizer).RequestAuthorizationCode (
    redirectUri.String(), response,
  )

  return err
}

/**
 * The [RequestToken] method...
 */
func (restOAuth2 *_BbRestOAuth2) RequestToken (
  grantType string, code string, redirectUri url.URL,
) (oauth2.AccessToken, error) {
  var accessToken oauth2.AccessToken
  var err error
println ("before creating authorizer")
  restOAuth2._createAuthorizer (grantType)
println ("after creating authorizer")
  if "client_credentials" == grantType {
    accessToken, err = restOAuth2.authorizer.RequestAuthorization()
  } else if "authorization_code" == grantType {
    println ("before casting rest user authorizer")
    userAuthorizer, _ := restOAuth2.authorizer.(oauth2.RestUserAuthorizer)
println ("after casting rest authorizer")
    println ("before requesting user authorization")
    accessToken, err =
      userAuthorizer.RequestUserAuthorization (code, redirectUri.String())
  }
println ("after creating access token, or encountered error")
  return accessToken, err
}

/**
 * The [_createAuthorizer] method...
 */
func (restOAuth2 *_BbRestOAuth2) _createAuthorizer (grantType string) {
  factory := new (oauth2.AuthorizerFactory)

  if "authorization_code" == grantType {
    restOAuth2.authorizer = factory.BuildAuthorizer (
      restOAuth2.host, restOAuth2.clientId, restOAuth2.secret, "user",
    )
  } else {
    restOAuth2.authorizer = factory.BuildAuthorizer (
      restOAuth2.host, restOAuth2.clientId, restOAuth2.secret, "",
    )
  }
}
