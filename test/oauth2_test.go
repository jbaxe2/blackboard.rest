package test

import (
  "errors"
  "github.com/jbaxe2/blackboard.rest.go/src"
  "github.com/jbaxe2/blackboard.rest.go/src/_scaffolding/config"
  "github.com/jbaxe2/blackboard.rest.go/src/oauth2"
  "net/url"
  "testing"
)

/**
 * The [OAuth2Tester] type...
 */
type OAuth2Tester struct {
  t *testing.T

  Testable
}

/**
 * The [Run] method...
 */
func (tester *OAuth2Tester) Run() {
  println ("\nOAuth2:")

  _testGetOAuth2Instance (tester.t)
  _testBuildClientRestAuthorizer (tester.t)
  _testObtainAccessToken (tester.t)
}

/**
 * The [_testGetOAuth2Instance] function...
 */
func _testGetOAuth2Instance (t *testing.T) {
  println ("Obtain a valid OAuth2 instance.")

  host, err := url.Parse (config.Host)

  if nil != err {
    t.Error ("Parsing the host failed, unable to obtain valid instance")

    return
  }

  oAuth2Service :=
    blackboard_rest.GetOAuth2Instance (*host, config.ClientId, config.Secret)

  if nil == oAuth2Service {
    t.Error ("Obtaining a valid OAuth2 instance failed\n")
  }
}

/**
 * The [_testBuildClientRestAuthorizer] function...
 */
func _testBuildClientRestAuthorizer (t *testing.T) {
  println ("Build a client REST Authorizer.")

  host, err := url.Parse (config.Host)

  if nil != err {
    t.Error ("Parsing the host failed, so unable to build the authorizer")

    return
  }

  restAuthorizer := (new (oauth2.AuthorizerFactory)).BuildAuthorizer (
    *host, config.ClientId, config.Secret, "",
  )

  if nil == restAuthorizer {
    err = errors.New ("could not create the REST Authorizer")

    t.Error ("Building a client REST Authorizer failed\n" + err.Error())
  }
}

/**
 * The [_testObtainAccessToken] function...
 */
func _testObtainAccessToken (t *testing.T) {
  println ("Obtain a valid REST access token.")

  host, err := url.Parse (config.Host)

  if nil != err {
    t.Error ("Parsing the host failed, so unable to obtain the access token")

    return
  }

  oauth2Service := blackboard_rest.GetOAuth2Instance (
    *host, config.ClientId, config.Secret,
  )

  accessToken, err := oauth2Service.RequestToken (
    "client_credentials", "", *host,
  )

  if ((oauth2.AccessToken{}) == accessToken) || (nil != err) {
    if nil == err {
      err = errors.New ("could not obtain a valid REST access token")
    }

    t.Error ("Obtaining a valid REST access token failed\n" + err.Error())
  }
}
