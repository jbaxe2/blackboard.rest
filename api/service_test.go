package api_test

import (
  "net/http"
  "testing"

  "github.com/jbaxe2/blackboard.rest/api"
  "github.com/jbaxe2/blackboard.rest/oauth2"
)

/**
 * The [TestCreateNewService] function...
 */
func TestCreateNewService (t *testing.T) {
  println ("Create a new service instance.")

  if nil == api.NewService ("localhost", mockToken, mockRoundTripper) {
    t.Error ("Creating a new service instance should not result in nil reference.")
  }
}

/**
 * The [TestNewServiceRequiresHost] function...
 */
func TestNewServiceRequiresHost (t *testing.T) {
  println ("Creating a new service instance requires a host.")

  if nil != api.NewService ("", mockToken, mockRoundTripper) {
    t.Error ("Missing host should result in nil service reference.")
  }
}

/**
 * The [TestNewServiceRequiresToken] function...
 */
func TestNewServiceRequiresToken (t *testing.T) {
  println ("Creating a new service instance requires a token.")

  if nil != api.NewService ("localhost", nil, mockRoundTripper) {
    t.Error ("Missing token should result in nil service reference.")
  }
}

/**
 * The [TestNewServiceCanHaveNilRoundTripper] function...
 */
func TestNewServiceCanHaveNilRoundTripper (t *testing.T) {
  println ("Creating a new service instance can be done with nil round tripper.")

  if nil == api.NewService ("localhost", mockToken, nil) {
    t.Error ("Nil round tripper should not result in nil service reference.")
  }
}

/**
 * The [TestNewServiceHasPertinentInformation] function...
 */
func TestNewServiceHasPertinentInformation (t *testing.T) {
  println ("New service instance retains the information used to create it.")

  host := "localhost"
  service := api.NewService (host, mockToken, mockRoundTripper)

  if !(service.Host() == host && service.Token() == mockToken) {
    t.Error ("New service instance should retain the info used to create it.")
  }
}

/**
 * The [TestNewServiceRequiresTokenHasAccessCode] function...
 */
func TestNewServiceRequiresTokenHasAccessCode (t *testing.T) {
  println ("OAuth2 token for service has an access code.")

  service := api.NewService (
    "localhost", new (_MockImproperAccessCodeToken), mockRoundTripper,
  )

  if nil != service {
    t.Error ("New service requires the OAuth2 token has an access code.")
  }
}

/**
 * The [TestNewServiceRequiresTokenHasGreaterThanZeroExpiresIn] function...
 */
func TestNewServiceRequiresTokenHasGreaterThanZeroExpiresIn (t *testing.T) {
  println ("OAuth2 token for service has expires in value greater than 0.")

  service := api.NewService (
    "localhost", new (_MockImproperExpiresInToken), mockRoundTripper,
  )

  if nil != service {
    t.Error ("New service requires the OAuth2 token has expires in value > 0.")
  }
}

/**
 * Mocked types and instances to run the above tests with.
 */
var mockToken = oauth2.NewToken (
  "access_token", "token_type", "refresh_token", "scope", "user_id", 3600,
)

/**
 * The [_MockImproperAccessCodeToken] type.
 */
type _MockImproperAccessCodeToken struct {
  oauth2.Token
}

func (_ *_MockImproperAccessCodeToken) AccessToken() string {
  return ""
}

func (_ *_MockImproperAccessCodeToken) ExpiresIn() int32 {
  return 360
}

/**
 * The [_MockImproperExpiresInToken] type.
 */
type _MockImproperExpiresInToken struct {
  oauth2.Token
}

func (_ *_MockImproperExpiresInToken) AccessToken() string {
  return "accessTokenValue"
}

func (_ *_MockImproperExpiresInToken) ExpiresIn() int32 {
  return 0
}

var mockRoundTripper = NewMockServiceRoundTripper()

/**
 * The [_MockServiceRoundTripper] type.
 */
type _MockServiceRoundTripper struct {
  http.RoundTripper
}

func NewMockServiceRoundTripper() http.RoundTripper {
  return new (_MockServiceRoundTripper)
}

func (roundTripper *_MockServiceRoundTripper) RoundTrip (
  request *http.Request,
) (*http.Response, error) {
  return request.Response, nil
}
