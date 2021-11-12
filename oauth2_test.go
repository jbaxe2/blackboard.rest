package blackboard_rest_test

import (
  "bytes"
  "encoding/json"
  "io/ioutil"
  "net/http"
  "net/url"
  "strings"
  "testing"

  blackboardRest "github.com/jbaxe2/blackboard.rest"
  "github.com/jbaxe2/blackboard.rest/oauth2"
)

/**
 * The [CreateNewOAuth2Instance] function...
 */
func TestCreateNewOAuth2Instance (t *testing.T) {
  println ("Create a new OAuth2 service instance.")

  if nil == blackboardRest.NewOAuth2 ("localhost", mockRoundTripper) {
    t.Error ("New OAuth2 instance should not be a nil reference.")
  }
}

/**
 * The [TestNewOAuth2InstanceRequiresHost] function...
 */
func TestNewOAuth2InstanceRequiresHost (t *testing.T) {
  println ("New OAuth2 instance requires a host.")

  if nil != blackboardRest.NewOAuth2 ("", mockRoundTripper) {
    t.Error ("Missing host should result in nil reference.")
  }
}

/**
 * The [TestNewOAuth2InstanceRequiresRoundTripper] function...
 */
func TestNewOAuth2InstanceCanHaveNilRoundTripper (t *testing.T) {
  println ("New OAuth2 instance can have a nil round tripper instance.")

  if nil == blackboardRest.NewOAuth2 ("localhost", nil) {
    t.Error ("Nil round tripper instance should not result in nil reference.")
  }
}

/**
 * The [TestNewOAuth2ObtainAuthorizationCode] function...
 */
func TestNewOAuth2ObtainAuthorizationCode (t *testing.T) {
  println ("Obtain an authorization code for the OAuth2 instance.")

  requestUrl := "https://localhost/learn/api/public/v1/oauth2/authorizecode"
  request, _ := http.NewRequest ("GET", requestUrl, strings.NewReader ("{}"))
  redirectUri, _ := url.Parse ("localhost")

  oAuth2 := blackboardRest.NewOAuth2 ("localhost", mockRoundTripper)
  response := oAuth2.AuthorizationCode (request, redirectUri, "clientId", "read")

  location, err := response.Location()

  result := nil != response && nil == err && 200 == response.StatusCode &&
    strings.Contains (location.String(), "&response_type=code")

  if !result {
    t.Error ("The authorization code response was not properly established.")
  }
}

/**
 * The [TestNewOAuth2RequestTokenWithClientCredentials] function...
 */
func TestNewOAuth2RequestTokenWithClientCredentials (t *testing.T) {
  println ("Request a new OAuth2 authorization token with client credentials.")

  redirectUri, _ := url.Parse ("localhost")

  oAuth2 := blackboardRest.NewOAuth2 ("localhost", mockRoundTripper)
  oAuth2.SetClientIdAndSecret ("clientId", "secret")
  token, err := oAuth2.RequestToken ("client_credentials", "authCode", redirectUri)

  if nil == token || nil != err {
    t.Error ("Obtaining an authorization token should complete successfully.")
  }
}

/**
 * The [TestNewOAuth2RequestNoClientIdAndSecret] function...
 */
func TestNewOAuth2RequestNoClientIdAndSecretProperError (t *testing.T) {
  println ("Missing client ID and secret results in proper error response.")

  redirectUri, _ := url.Parse ("localhost")

  oAuth2 := blackboardRest.NewOAuth2 ("localhost", mockRoundTripper)
  token, err := oAuth2.RequestToken ("client_credentials", "authCode", redirectUri)

  if nil != token || nil == err {
    t.Error ("Missing client ID and secret should result in proper error.")
  }
}

/**
 * The [TestNewOAuth2TokenRequestRequiresAppropriateGrantType] function...
 */
func TestNewOAuth2TokenRequestRequiresAppropriateGrantType (t *testing.T) {
  println ("The grant type must be of an acceptable value for a token request.")

  redirectUri, _ := url.Parse ("localhost")

  oAuth2 := blackboardRest.NewOAuth2 ("localhost", mockRoundTripper)
  oAuth2.SetClientIdAndSecret ("clientId", "secret")
  token, err := oAuth2.RequestToken ("improper_grant", "authCode", redirectUri)

  if nil != token || nil == err {
    t.Error ("Improper grant type should result in proper error.")
  }
}

/**
 * The [TestNewOAuth2RequestTokenNoAuthorizationCodeProperError] function...
 */
func TestNewOAuth2RequestTokenNoAuthorizationCodeProperError (t *testing.T) {
  println (
    "Missing authorization code when requesting token results in proper error response.",
  )

  redirectUri, _ := url.Parse ("localhost")

  oAuth2 := blackboardRest.NewOAuth2 ("localhost", mockRoundTripper)
  oAuth2.SetClientIdAndSecret ("clientId", "secret")
  token, err := oAuth2.RequestToken ("authorization_code", "", redirectUri)

  if nil != token || nil == err {
    t.Error ("Missing authorization code should result in proper error.")
  }
}

/**
 * Mocked types and instances to run the above tests with.
 */
var mockRoundTripper = NewMockOAuth2RoundTripper()

var mockToken = oauth2.NewToken (
  "access_token", "token_type", "refresh_token", "scope", "user_id", 3600,
)

/**
 * The [_MockOAuth2RoundTripper] type.
 */
type _MockOAuth2RoundTripper struct {
  http.RoundTripper
}

func NewMockOAuth2RoundTripper() http.RoundTripper {
  return new (_MockOAuth2RoundTripper)
}

/**
 * The [RoundTrip] method mocks out the transactional HTTP requests.
 */
func (_ *_MockOAuth2RoundTripper) RoundTrip (
  request *http.Request,
) (*http.Response, error) {
  request.Response = &http.Response {
    Request: request,
    Header: make (http.Header),
  }

  switch true {
    case strings.Contains (request.URL.Path, "oauth2/authorizationcode"):
      request.Response.Status = "200 OK"
      request.Response.StatusCode = 200
    case strings.Contains (request.URL.Path, "oauth2/token"):
      if user, pass, _ := request.BasicAuth(); "" == user || "" == pass {
        request.Response.StatusCode = 401
        request.Response.Body = ioutil.NopCloser (strings.NewReader (
          `{"error":"invalid_client","error_description":"Invalid client ` +
          `credentials, or no access granted to this Learn server."}`,
        ))
      } else {
        request.Response.Body =
          ioutil.NopCloser (bytes.NewReader (_mockTokenBytes()))
      }
  }

  return request.Response, nil
}

/**
 * The [_mockTokenBytes] function creates a mock token, marshalled in JSON and
 * encoded as a sequence of bytes.
 */
func _mockTokenBytes() []byte {
  tokenInfo := map[string]interface{} {
    "access_token": "accessTokenValue",
    "token_type": "someTokenType",
    "expires_in": int32 (3600),
    "refresh_token": "",
    "scope": "read",
    "user_id": "someUserId",
  }

  tokenBytes, _ := json.Marshal (tokenInfo)

  return tokenBytes
}
