package connector

import (
  "encoding/json"
  "errors"
  "io/ioutil"
  "net/http"
  "net/url"
  "strings"

  "github.com/jbaxe2/blackboard.rest.go/src/_scaffolding/config"
  "github.com/jbaxe2/blackboard.rest.go/src/oauth2"
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

  base := config.Base

  if 2 == useVersion {
    base = config.BaseV2
  }

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
  var response *http.Response
  var err error

  if 0 < len (query) {
    queryString := ""

    for k, v := range query {
      queryString += k + "=" + v.(string) + "&"
    }

    endpoint.RawQuery = queryString[:(len (queryString) - 1)]
  }

  request := new (http.Request)
  request.URL = endpoint

  request.Header = make (http.Header)

  for k, v := range headers {
    request.Header.Set (k, v)
  }

  response, err = (new (http.Client)).Do (request)

  return response, err
}

/**
 * The [_handlePostRequest] function...
 */
func _handlePostRequest (
  endpoint *url.URL, headers map[string]string, body map[string]interface{},
) (*http.Response, error) {
  var response *http.Response
  var err error

  return response, err
}
