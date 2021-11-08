package api

import (
  "encoding/json"
  "errors"
  "io/ioutil"
  "net/http"
  "strconv"
  "strings"

  restErrors "github.com/jbaxe2/blackboard.rest/api/errors"
  "github.com/jbaxe2/blackboard.rest/oauth2"
  "github.com/jbaxe2/blackboard.rest/utils"
)

/**
 * The [Service] interface is the base interface for all Blackboard Learn REST
 * API service types.
 */
type Service interface {
  Host() string

  Token() oauth2.Token

  Request (
    endpoint string, method string, data map[string]interface{},
    options map[string]string, useVersion int,
  ) (map[string]interface{}, error)
}

/**
 * The [_Service] type implements the Service interface.
 */
type _Service struct {
  host string

  token oauth2.Token

  roundTripper http.RoundTripper

  Service
}

/**
 * The [NewService] function creates a new Service instance.
 */
func NewService (
  host string, token oauth2.Token, roundTripper http.RoundTripper,
) Service {
  if "" == host || nil == token || "" == token.AccessToken() ||
     1 > token.ExpiresIn() {
    return nil
  }

  if nil == roundTripper {
    roundTripper = http.DefaultTransport
  }

  return &_Service {
    host: host,
    token: token,
    roundTripper: roundTripper,
  }
}

func (service *_Service) Host() string {
  return service.host
}

func (service *_Service) Token() oauth2.Token {
  return service.token
}

/**
 * The [Request] method makes the request to the REST API.
 */
func (service *_Service) Request (
  endpoint string, method string, data map[string]interface{},
  options map[string]string, useVersion int,
) (map[string]interface{}, error) {
  if err := _verifyRequestConditions (endpoint, method); nil != err {
    return nil, err
  }

  requestUri := _buildRequestUri (service.host, endpoint, useVersion, options)

  request, _ := http.NewRequest (method, requestUri, nil)
  request.Header.Set ("Authorization", "Bearer " + service.token.AccessToken())

  client := http.Client {Transport: service.roundTripper}

  response, _ := client.Do (request)

  return _parseResponse (response)
}

/**
 * The [_verifyRequestConditions] function verifies the information used to
 * create a service request is appropriate.
 */
func _verifyRequestConditions (endpoint, method string) error {
  if "" == endpoint {
    return errors.New ("missing service endpoint")
  }

  methods := []string {"GET", "POST", "PUT", "PATCH", "DELETE"}

  if !utils.StringInStrings (method, methods) {
    return errors.New ("inappropriate HTTP method")
  }

  return nil
}

/**
 * The [_buildRequestUri] function builds the URI that will be used for making
 * some REST API request.
 */
func _buildRequestUri (
  host, endpoint string, useVersion int, options map[string]string,
) string {
  endpoint = strings.Replace (endpoint, "{v}", strconv.Itoa (useVersion), 1)
  uri := "https://" + host + endpoint

  if 0 < len (options) {
    uri += "?"

    for k, v := range options {
      uri += k + "=" + v + "&"
    }

    uri = uri[0:len (uri) - 2]
  }

  return uri
}

/**
 * The [_parseResponse] function parses the response from a REST API request,
 * converting the response to either a raw map with string-based keys or an error
 * of some sort.
 */
func _parseResponse (response *http.Response) (map[string]interface{}, error) {
  defer response.Body.Close()

  var rawResponse map[string]interface{}
  responseBytes, _ := ioutil.ReadAll (response.Body)

  if err := json.Unmarshal (responseBytes, &rawResponse); nil != err {
    return nil, errors.New ("response from the REST server is not unreadable")
  }

  if _, wasError := rawResponse["status"]; wasError {
    return nil, restErrors.NewRestExceptionFromRaw (rawResponse)
  }

  return rawResponse, nil
}
