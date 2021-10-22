package connector

import (
  "encoding/json"
  "errors"
  "io/ioutil"
  "net/http"
  "net/url"
  "strconv"
  "strings"

  "github.com/jbaxe2/blackboard.rest/_scaffolding/config"
  "github.com/jbaxe2/blackboard.rest/oauth2"
)

/**
 * The [BlackboardRestConnector] interface...
 */
type BlackboardRestConnector interface {
  SendBbRequest (
    endpoint string, method string, data map[string]interface{}, useVersion int,
  ) (interface{}, error)
}

/**
 * The [BbRestConnector] type...
 */
type BbRestConnector struct {
  host string

  accessToken oauth2.AccessToken

  BlackboardRestConnector
}

func (connector *BbRestConnector) SetHost (host string) {
  connector.host = host
}

func (connector *BbRestConnector) SetAccessToken (token oauth2.AccessToken) {
  connector.accessToken = token
}

/**
 * The [SendBbRequest] method...
 */
func (connector *BbRestConnector) SendBbRequest (
  endpoint string, method string, data map[string]interface{}, useVersion int,
) (interface{}, error) {
  var result interface{}
  var err error
  var endpointUri *url.URL
  var responseBytes []byte

  if 0 == len (endpoint) {
    return result, errors.New ("no endpoint to send a REST request to")
  }

  base := strings.Replace (config.Base, "{v}", strconv.Itoa (useVersion), 1)

  if endpointUri, err = url.Parse (connector.host + base + endpoint); nil != err {
    return result, err
  }

  headers := make (map[string]string)
  headers["Authorization"] = "Bearer " + connector.accessToken.AccessToken()

  var response *http.Response

  if "get" == strings.ToLower (method) {
    response, err = _handleGetRequest (endpointUri, headers, data)
  } else if "post" == strings.ToLower (method) {
    response, err = _handlePostRequest (endpointUri, headers, data)
  } else {
    return result, errors.New ("specified method is currently unsupported")
  }

  if nil != err {
    return result, err
  }

  responseBytes, err = ioutil.ReadAll (response.Body)

  err = json.Unmarshal (responseBytes, &result)
  err = response.Body.Close()

  if response.StatusCode >= 300 {
    return result, errors.New ("the returned response resulted in error")
  }

  return result, err
}

/**
 * The [_handleGetRequest] function...
 */
func _handleGetRequest (
  endpoint *url.URL, headers map[string]string, query map[string]interface{},
) (*http.Response, error) {
  if 0 < len (query) {
    queryString := ""

    for k, v := range query {
      queryString += k + "=" + v.(string) + "&"
    }

    endpoint.RawQuery = url.QueryEscape (queryString[:(len (queryString) - 1)])
  }

  request := new (http.Request)
  request.Header = make (http.Header)
  request.URL = endpoint

  for k, v := range headers {
    request.Header.Add (k, v)
  }

  return (new (http.Client)).Do (request)
}

/**
 * The [_handlePostRequest] function...
 */
func _handlePostRequest (
  endpoint *url.URL, headers map[string]string, body map[string]interface{},
) (*http.Response, error) {
  request := new (http.Request)
  request.Header = make (http.Header)
  request.URL = endpoint
  request.Method = "POST"

  for k, v := range headers {
    request.Header.Add (k, v)
  }

  for k, v := range body {
    request.Form.Add (k, v.(string))
  }

  return (new (http.Client)).Do (request)
}
