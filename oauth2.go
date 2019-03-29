package blackboard_rest_go

import (
  "github.com/jbaxe2/blackboard.rest.go/oauth2"
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
 * The [BbRestOAuth2] type...
 */
type BbRestOAuth2 struct {
  host url.URL

  clientId, secret string

  authorizer oauth2.RestAuthorizer

  BlackboardRestOAuth2
}

func (restOAuth2 *BbRestOAuth2) Host() url.URL {
  return restOAuth2.host
}

func (restOAuth2 *BbRestOAuth2) ClientId() string {
  return restOAuth2.clientId
}

func (restOAuth2 *BbRestOAuth2) Secret() string {
  return restOAuth2.secret
}

/**
 * The [GetOAuth2Instance] function...
 */
func GetOAuth2Instance (
  host url.URL, clientId string, secret string,
) BlackboardRestOAuth2 {
  oauth2Instance :=  new (BbRestOAuth2)

  oauth2Instance.host = host
  oauth2Instance.clientId = clientId
  oauth2Instance.secret = secret

  return oauth2Instance
}

/**
 * The [GetAuthorizationCode] method...
 */
func (restOAuth2 *BbRestOAuth2) GetAuthorizationCode (
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
func (restOAuth2 *BbRestOAuth2) RequestToken (
  grantType string, code string, redirectUri url.URL,
) (oauth2.AccessToken, error) {
  var accessToken oauth2.AccessToken
  var err error

  return accessToken, err
}

/**
 * The [_createAuthorizer] method...
 */
func (restOAuth2 *BbRestOAuth2) _createAuthorizer (grantType string) {
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
