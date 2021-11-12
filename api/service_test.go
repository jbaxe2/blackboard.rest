package api_test

import (
  "io/ioutil"
  "net/http"
  "strings"
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
 * The [TestNewServiceRequest] function...
 */
func TestNewServiceRequest (t *testing.T) {
  println ("Create a new service request.")

  service := api.NewService ("localhost", mockToken, mockRoundTripper)

  if _, err := service.Request ("/endpoint/string", "GET", nil, 1); nil != err {
    t.Error ("Performing a service request should not result in error.")
  }
}

/**
 * The [TestNewServiceRequestRequiresEndpoint] function...
 */
func TestNewServiceRequestRequiresEndpoint (t *testing.T) {
  println ("New service request requires endpoint.")

  service := api.NewService ("localhost", mockToken, mockRoundTripper)

  if _, err := service.Request ("", "GET", nil, 1); nil == err {
    t.Error ("Missing endpoint should result in error.")
  }
}

/**
 * The [TestNewServiceRequestRequiresHttpMethod] function...
 */
func TestNewServiceRequestRequiresHttpMethod (t *testing.T) {
  println ("New service request requires an HTTP method.")

  service := api.NewService ("localhost", mockToken, mockRoundTripper)

  if _, err := service.Request ("/endpoint/string", "", nil, 1);
      nil == err {
    t.Error ("Inappropriate HTTP method should result in error.")
  }
}

/**
 * The [TestNewServiceRequiresRequiresAppropriateMethod] function...
 */
func TestNewServiceRequiresRequiresAppropriateMethod (t *testing.T) {
  println ("New service request requires an appropriate HTTP method.")

  service := api.NewService ("localhost", mockToken, mockRoundTripper)

  if _, err := service.Request ("/endpoint/string", "METHOD", nil, 1);
      nil == err {
    t.Error ("Inappropriate HTTP method should result in error.")
  }
}

/**
 * The [TestNewServiceRequestsRequireValidAccessCode] function...
 */
func TestNewServiceRequestsRequireValidAccessCode (t *testing.T) {
  println ("A valid access code from an authorized token is required for requests.")

  service :=
    api.NewService ("localhost", new (_MockInvalidAccessCodeToken), mockRoundTripper)

  _, err := service.Request ("/access/invalid/token", "GET", nil, 1)

  if nil == err {
    t.Error ("An invalid access code should result in a returned error.")
  }
}

/**
 * The [TestNewServiceSetOptionKeyAndValue] function...
 */
func TestNewServiceSetOptionKeyAndValue (t *testing.T) {
  println ("New service appropriately set and use request option key and value.")

  service := api.NewService ("localhost", mockToken, mockRoundTripper)
  service.SetRequestOption ("key", "value")

  response, _ := service.Request ("/set/request/option", "GET", nil, 1)

  if "value" != response["result"].(string) {
    t.Error ("The request option value was not set or used appropriately.")
  }
}

/**
 * The [TestNewServiceOptionsClearedAfterRequest] function...
 */
func TestNewServiceOptionsClearedAfterRequest (t *testing.T) {
  println ("Service request options are cleared after a request.")

  service := api.NewService ("localhost", mockToken, mockRoundTripper)
  service.SetRequestOption ("key", "value")

  _, _ = service.Request ("/set/request/option", "GET", nil, 1)
  response, _ := service.Request ("/set/request/option", "GET", nil, 1)

  if "" != response["result"].(string) {
    t.Error ("The request option value was not cleared appropriately.")
  }
}

/**
 * The [TestNewServiceClearingOptionsClearsOptions] function...
 */
func TestNewServiceClearingOptionsClearsOptions (t *testing.T) {
  println ("Clearing service request options clears the request options.")

  service := api.NewService ("localhost", mockToken, mockRoundTripper)
  service.SetRequestOption ("key", "value")
  service.ClearRequestOptions()

  _, _ = service.Request ("/set/request/option", "GET", nil, 1)
  response, _ := service.Request ("/set/request/option", "GET", nil, 1)

  if "" != response["result"].(string) {
    t.Error ("The request option value was not cleared appropriately.")
  }
}

/**
 * The [TestNewServiceAddRequestOptions] function...
 */
func TestNewServiceAddRequestOptions (t *testing.T) {
  println ("New service can have multiple request options added and used.")

  options := map[string]string {
    "key1": "value1",
    "key2": "value2",
  }

  service := api.NewService ("localhost", mockToken, mockRoundTripper)
  service.AddRequestOptions (options)

  response, _ := service.Request ("/add/request/options", "GET", nil, 1)
  value1 := response["key1"].(string)
  value2 := response["key2"].(string)

  if !("value1" == value1 && "value2" == value2) {
    t.Error ("The request options were not added and used appropriately.")
  }
}

/**
 * The [TestNewServiceOptionNotSetIfKeyIsEmpty] function...
 */
func TestNewServiceOptionNotSetIfKeyIsEmpty (t *testing.T) {
  println ("Service request option is not set if key is empty.")

  service := api.NewService ("localhost", mockToken, mockRoundTripper)
  service.SetRequestOption ("", "value")

  response, _ := service.Request ("/set/empty/option", "GET", nil, 1)

  if "ok" != response["result"].(string) {
    t.Error ("The request option value was not set or used appropriately.")
  }
}

/**
 * The [TestNewServiceSetOptionKeyAndValue] function...
 */
func TestNewServiceSetOptionKeyAndReplacedValue (t *testing.T) {
  println ("Replaced value for service option is appropriately set and used.")

  service := api.NewService ("localhost", mockToken, mockRoundTripper)
  service.SetRequestOption ("key", "value")
  service.SetRequestOption ("key", "newValue")

  response, _ := service.Request ("/set/request/option", "GET", nil, 1)

  if "newValue" != response["result"].(string) {
    t.Error ("The request option new value was not set or used appropriately.")
  }
}

/**
 * Mocked types and instances to run the above tests with.
 */
var mockToken = oauth2.NewToken (
  "access_token", "token_type", "refresh_token", "scope", "user_id", 3600,
)

var mockRoundTripper = NewMockServiceRoundTripper()

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

/**
 * The [_MockInvalidAccessCodeToken] type.
 */
type _MockInvalidAccessCodeToken struct {
  oauth2.Token
}

func (_ *_MockInvalidAccessCodeToken) AccessToken() string {
  return "invalid_access_code"
}

func (_ *_MockInvalidAccessCodeToken) ExpiresIn() int32 {
  return 360
}

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
  request.Response = &http.Response {
    Request: request,
    Header: make (http.Header),
  }

  responseBody := ""

  switch true {
    case strings.Contains (request.URL.Path, "invalid/token"):
      responseBody = `{"status":"400","message":"invalid_client"}`
    case strings.Contains (request.URL.Path, "set/request/option"):
      result := request.URL.Query().Get ("key")

      responseBody = `{"result":"` + result + `"}`
    case strings.Contains (request.URL.Path, "set/empty/option"):
      if "" != request.URL.RawQuery {
        responseBody = `{"result":"empty_key"}`
      } else {
        responseBody = `{"result":"ok"}`
      }
    case strings.Contains (request.URL.Path, "add/request/options"):
      value1 := request.URL.Query().Get ("key1")
      value2 := request.URL.Query().Get ("key2")

      responseBody = `{"key1":"` + value1 + `", "key2":"` + value2 + `"}`
    default:
      responseBody = `{"result":"ok"}`
  }

  request.Response.Body = ioutil.NopCloser (strings.NewReader (responseBody))

  return request.Response, nil
}
