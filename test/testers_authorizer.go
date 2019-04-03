package test

import (
  "github.com/jbaxe2/blackboard.rest.go/src"
  "github.com/jbaxe2/blackboard.rest.go/src/_scaffolding/config"
  "github.com/jbaxe2/blackboard.rest.go/src/oauth2"
  "net/url"
)

/**
 * The [TestersAuthorizer] type...
 */
type TestersAuthorizer struct {
  accessToken oauth2.AccessToken
}

func (authorizer *TestersAuthorizer) AccessToken() oauth2.AccessToken {
  return authorizer.accessToken
}

/**
 * The [AuthorizeForTests] method...
 */
func (authorizer *TestersAuthorizer) AuthorizeForTests() error {
  var err error

  host, _ := url.Parse (config.Host)

  restOAuth2 := blackboard_rest.GetOAuth2Instance (
    *host, config.ClientId, config.Secret,
  )

  authorizer.accessToken, err = restOAuth2.RequestToken (
    "client_credentials", "", *host,
  )

  return err
}
