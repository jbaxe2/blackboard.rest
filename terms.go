package blackboard_rest

import (
  "net/url"
  "strings"

  "github.com/jbaxe2/blackboard.rest/_scaffolding/config"
  "github.com/jbaxe2/blackboard.rest/oauth2"
  "github.com/jbaxe2/blackboard.rest/terms"
)

/**
 * The [Terms] interface...
 */
type Terms interface {
  GetTerms() ([]terms.Term, error)

  CreateTerm (term terms.Term) error

  GetTerm (termId string) (terms.Term, error)

  UpdateTerm (termId string, term terms.Term) error
}

/**
 * The [_BbRestTerms] type...
 */
type _BbRestTerms struct {
  _BlackboardRest

  Terms
}

/**
 * The [GetTermsInstance] function...
 */
func GetTermsInstance (host string, accessToken oauth2.AccessToken) Terms {
  hostUri, _ := url.Parse (host)

  termsService := new (_BbRestTerms)

  termsService.host = *hostUri
  termsService.accessToken = accessToken

  termsService.service.SetHost (host)
  termsService.service.SetAccessToken (accessToken)

  return termsService
}

/**
 * The [GetTerms] method...
 */
func (restTerms *_BbRestTerms) GetTerms() ([]terms.Term, error) {
  endpoint := config.TermsEndpoints["terms"]

  result, err := restTerms.service.Connector.SendBbRequest (
    endpoint, "GET", make (map[string]interface{}), 1,
  )

  if nil != err {
    return []terms.Term{}, err
  }

  rawTerms := result.(map[string]interface{})["results"]

  theTerms := terms.NewTerms(_normalizeRawTerms (rawTerms.([]interface{})))

  return theTerms, err
}

/**
 * The [GetTerm] method...
 */
func (restTerms *_BbRestTerms) GetTerm (termId string) (terms.Term, error) {
  endpoint := config.TermsEndpoints["term"]
  endpoint = strings.Replace (endpoint, "{termId}", termId, -1)

  result, err := restTerms.service.Connector.SendBbRequest (
    endpoint, "GET", make (map[string]interface{}), 1,
  )

  term := terms.NewTerm(result.(map[string]interface{}))

  return term, err
}

/**
 * The [_normalizeRawTerms] function...
 */
func _normalizeRawTerms (rawTerms []interface{}) []map[string]interface{} {
  mappedRawTerms := make ([]map[string]interface{}, len (rawTerms))

  for i, rawTerm := range rawTerms {
    mappedRawTerms[i] = rawTerm.(map[string]interface{})
  }

  return mappedRawTerms
}
