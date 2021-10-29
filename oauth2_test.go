package blackboard_rest_test

import (
  "encoding/json"
  "io"
  "net/http"
  "net/url"
  "strings"
  "testing"

  blackboardRest "github.com/jbaxe2/blackboard.rest"
)

/**
 * The [CreateNewOAuth2Instance] function...
 */
func TestCreateNewOAuth2Instance (t *testing.T) {
  println ("Create a new OAuth2 service instance.")

  if nil == blackboardRest.NewOAuth2 ("localhost", mockRoundTripper) {
    t.Error ("New OAuth2 instance should not be a nil reference.")
    t.FailNow()
  }
}

/**
 * The [TestNewOAuth2InstanceRequiresHost] function...
 */
func TestNewOAuth2InstanceRequiresHost (t *testing.T) {
  println ("New OAuth2 instance requires a host.")

  if nil != blackboardRest.NewOAuth2 ("", mockRoundTripper) {
    t.Error ("Missing host should result in nil reference.")
    t.FailNow()
  }
}

/**
 * The [TestNewOAuth2InstanceRequiresRoundTripper] function...
 */
func TestNewOAuth2InstanceRequiresRoundTripper (t *testing.T) {
  println ("New OAuth2 instance requires a RoundTripper instance.")

  if nil != blackboardRest.NewOAuth2 ("localhost", nil) {
    t.Error ("Missing round tripper instance should result in nil reference.")
    t.FailNow()
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
    t.FailNow()
  }
}

/**
 * The [TestNewOAuth2RequestTokenWithClientCredentials] function...
 */
func TestNewOAuth2RequestTokenWithClientCredentials (t *testing.T) {
  println ("Request a new OAuth2 authorization token with client credentials.")

  redirectUri, _ := url.Parse ("localhost")

  oAuth2 := blackboardRest.NewOAuth2 ("localhost", mockRequestTokenRoundTripper)
  token, err := oAuth2.RequestToken ("client_credentials", "authCode", redirectUri)

  if nil == token || nil != err {
    t.Error ("Obtaining an authorization token should complete successfully.")
    t.FailNow()
  }
}

/**
 * Mocked types and instances to run the above tests with.
 */
var mockRoundTripper = NewMockRoundTripper (NewMockEmptyWriter())
var mockRequestTokenRoundTripper = NewMockRoundTripper (NewMockTokenInfoWriter())

/**
 * The [_MockRoundTripper] type.
 */
type _MockRoundTripper struct {
  writer io.Writer

  http.RoundTripper
}

func NewMockRoundTripper (mockWriter io.Writer) http.RoundTripper {
  return &_MockRoundTripper {
    writer: mockWriter,
  }
}

func (roundTripper *_MockRoundTripper) RoundTrip (
  request *http.Request,
) (*http.Response, error) {
  request.Response = &http.Response {
    Request: request,
    Header: make (http.Header),
  }

  if strings.Contains (request.URL.Path, "oauth2/authorizationcode") {
    request.Response.Status = "200 OK"
    request.Response.StatusCode = 200
  }

  _ = request.Response.Write (roundTripper.writer)

  return request.Response, nil
}

/**
 * The [_MockWriter] type.
 */
type _MockWriter struct {
  data string

  io.Writer
}

func NewMockEmptyWriter() io.Writer {
  return &_MockWriter {
    data: "",
  }
}

func NewMockTokenInfoWriter() io.Writer {
  tokenInfo := map[string]interface{} {
    "access_token": "accessTokenValue",
    "token_type": "someTokenType",
    "expires_in": 3600,
    "refresh_token": "",
    "scope": "read",
    "user_id": "someUserId",
  }

  tokenBytes, _ := json.Marshal (tokenInfo)

  return &_MockWriter {
    data: string (tokenBytes),
  }
}

func (writer *_MockWriter) Write (bytes []byte) (int, error) {
  writer.data += string (bytes)

  return len (bytes), nil
}
