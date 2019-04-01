package test

import (
  "github.com/jbaxe2/blackboard.rest.go/src"
  "github.com/jbaxe2/blackboard.rest.go/src/config"
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
}

/**
 * The [_testGetOAuth2Instance] function...
 */
func _testGetOAuth2Instance (t *testing.T) {
  print ("Obtain a valid OAuth2 instance.\n")

  host, err := url.Parse (config.Host)

  blackboard_rest.GetOAuth2Instance (host, config.ClientId, config.Secret)

  if nil != err {
    t.Error ("Obtaining a valid OAuth2 instance failed\n" + err.Error())
  }
}
