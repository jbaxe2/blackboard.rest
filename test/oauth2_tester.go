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
}

/**
 * The [Run] method...
 */
func (tester *OAuth2Tester) Run() {
  print ("OAuth2:\n")

  _testGetOAuth2Instance (tester.t)
  _testBuildClientRestAuthorizer (tester.t)
  _testObtainAccessToken (tester.t)
}

/**
 * The [_testGetOAuth2Instance] function...
 */
func _testGetOAuth2Instance (t *testing.T) {
  print ("Obtain a valid OAuth2 instance.\n")

  host, err := url.Parse (config.Host)

  if nil != err {
    t.Error ("Parsing the host failed, unable to obtain valid instance")
  }

  blackboard_rest.GetOAuth2Instance (*host, config.ClientId, config.Secret)

  if nil != err {
    t.Error ("Obtaining a valid OAuth2 instance failed\n" + err.Error())
  }
}

/**
 * The [_testBuildClientRestAuthorizer] function...
 */
func _testBuildClientRestAuthorizer (t *testing.T) {
  print ("Build a client REST Authorizer.\n")

  host, err := url.Parse (config.Host)

  if nil != err {
    t.Error ("Parsing the host failed, so unable to build the authorizer")
  }

  restAuthorizer := (new (oauth2.AuthorizerFactory)).BuildAuthorizer (
    *host, config.ClientId, config.Secret, "",
  )

  if (nil == restAuthorizer) || (nil != err) {
    if nil == err {
      err = errors.New ("could not create the REST Authorizer")
    }

    t.Error ("Building a client REST Authorizer failed\n" + err.Error())
  }
}

/**
 * The [_testObtainAccessToken] function...
 */
func _testObtainAccessToken (t *testing.T) {
  print ("Obtain a valid REST access token.\n")

  host, err := url.Parse (config.Host)

  if nil != err {
    t.Error ("Parsing the host failed, so unable to obtain the access token")
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
