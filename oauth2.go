package blackboard_rest_go

import (
  "github.com/jbaxe2/blackboard.rest.go/oauth2"
  "net/http"
  "net/url"
)

/// The [BlackboardRestOAuth2] interface...
type BlackboardRestOAuth2 interface {
  GetAuthorizationCode (
    redirectUri url.URL, responseType string, response http.Response,
  ) error

  RequestToken (
    grantType string, code string, redirectUri url.URL,
  ) (oauth2.AccessToken, error)
}

/// The [BbRestOAuth2] type...
type BbRestOAuth2 struct {
  host url.URL

  clientId, secret string

  BlackboardRestOAuth2
}

func (oauth2 *BbRestOAuth2) Host() url.URL {
  return oauth2.host
}

func (oauth2 *BbRestOAuth2) ClientId() string {
  return oauth2.clientId
}

func (oauth2 *BbRestOAuth2) Secret() string {
  return oauth2.secret
}

func (oauth2 *BbRestOAuth2) GetAuthorizationCode (
  redirectUri url.URL, responseType string, response http.Response,
) error {
  var err error

  return err
}

func (oauth2 *BbRestOAuth2) RequestToken (
    grantType string, code string, redirectUri url.URL,
) (oauth2.AccessToken, error) {
  var accessToken oauth2.AccessToken
  var err error

  return accessToken, err
}
