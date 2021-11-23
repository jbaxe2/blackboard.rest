package blackboard_rest_test

import (
  "net/http"
  "strings"
  "testing"

  blackboardRest "github.com/jbaxe2/blackboard.rest"
  "github.com/jbaxe2/blackboard.rest/api"
)

/**
 * The [TestCreateNewTerms] function...
 */
func TestCreateNewTerms (t *testing.T) {
  println ("Create a new terms instance.")

  if nil == blackboardRest.NewTerms (mockTermsService) {
    t.Error ("Creating a new terms instance should not be a nil reference.")
  }
}

/**
 * The [TestNewTermsRequiresService] function...
 */
func TestNewTermsRequiresService (t *testing.T) {
  println ("Creating new terms instance requires a service instance.")

  if nil != blackboardRest.NewTerms (nil) {
    t.Error ("Missing service instance should result in nil reference.")
  }
}

/**
 * The [TestNewTermsGetTerms] function...
 */
func TestNewTermsGetTerms (t *testing.T) {
  println ("Retrieve multiple terms from the REST API.")

  terms := blackboardRest.NewTerms (mockTermsService)
  newTerms, err := terms.GetTerms()

  if !(nil == err && 2 == len (newTerms)) {
    t.Error ("Retrieving terms should result in the appropriate responses.")
  }
}

/**
 * Mocked types and instances to run the above tests with.
 */
var mockTermsService =
  api.NewService ("localhost", mockToken, mockTermsRoundTripper)

var mockTermsRoundTripper = new (_MockTermsRoundTripper)

type _MockTermsRoundTripper struct {
  http.RoundTripper
}

func (roundTripper *_MockTermsRoundTripper) RoundTrip (
  request *http.Request,
) (*http.Response, error) {
  conditions := []bool {
    "GET" == request.Method && strings.Contains (request.URL.Path, "/terms/"),
    "GET" == request.Method && strings.Contains (request.URL.Path, "/terms"),
  }

  responseBodies := []string {rawTerm, rawTerms}

  builder := NewResponseBuilder (conditions, responseBodies)

  return builder.Build (request), nil
}

const rawTerms = `{"results":[` + rawTerm + `,` + rawTerm2 + `]}`

const rawTerm = `{"id":"termId","externalId":"externalTermId","dataSourceId":` +
  `"term.data.source.id","name":"Term Name","description":"","availability":{` +
  `"available": "Yes","duration":{"type": "Continuous","start":` +
  `"2021-11-16T20:38:45.738Z","end":"2021-11-16T20:38:45.738Z","daysOfUse":0}}}`

const rawTerm2 = `{"id":"termId2","externalId":"externalTermId2","dataSourceId":` +
  `"term.data.source.id","name":"Term Name 2","description":"","availability":{` +
  `"available": "Yes","duration":{"type": "Continuous","start":` +
  `"2021-11-16T20:38:45.738Z","end":"2021-11-16T20:38:45.738Z","daysOfUse":0}}}`
