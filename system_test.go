package blackboard_rest_test

import (
  "net/http"
  "strings"
  "testing"

  blackboardRest "github.com/jbaxe2/blackboard.rest"
  "github.com/jbaxe2/blackboard.rest/api"
)

/**
 * The [TestCreateNewSystem] function...
 */
func TestCreateNewSystem (t *testing.T) {
  println ("Create a new system instance.")

  if nil == blackboardRest.NewSystem (mockSystemService) {
    t.Error ("Creating a new system instance should not be a nil reference.")
  }
}

/**
 * The [TestNewSystemRequiresService] function...
 */
func TestNewSystemRequiresService (t *testing.T) {
  println ("Creating new system instance requires a service instance.")

  if nil != blackboardRest.NewSystem (nil) {
    t.Error ("Missing service instance should result in nil reference.")
  }
}

/**
 * The [TestNewSystemGetVersionInfo] function...
 */
func TestNewSystemGetVersionInfo (t *testing.T) {
  println ("Retrieve the version info from the REST API.")

  systemApi := blackboardRest.NewSystem (mockSystemService)
  versionInfo, err := systemApi.GetVersion()

  if !(nil == err && versionInfo.Learn.Build == "rel.11+732fc84") {
    if nil != err {
      t.Error (err.Error())
    }
    println (versionInfo.Learn.Build)
    t.Error ("Retrieving version info should have the appropriate response.")
  }
}

/**
 * Mocked types and instances to run the above tests with.
 */
var mockSystemService =
  api.NewService ("localhost", mockToken, new (_MockSystemRoundTripper))

type _MockSystemRoundTripper struct {
  http.RoundTripper
}

func (roundTripper *_MockSystemRoundTripper) RoundTrip (
  request *http.Request,
) (*http.Response, error) {
  conditions := []bool {
    "GET" == request.Method && strings.Contains (request.URL.Path, "system/version"),
  }

  responseBodies := []string {rawVersionInfo}

  builder := NewResponseBuilder (conditions, responseBodies)

  return builder.Build (request), nil
}

const rawVersionInfo =
  `{"learn":{"major":3900,"minor":28,"patch":0,"build":"rel.11+732fc84"}}`
