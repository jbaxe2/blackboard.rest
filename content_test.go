package blackboard_rest_test

import (
  "net/http"
  "testing"

  blackboardRest "github.com/jbaxe2/blackboard.rest"
  "github.com/jbaxe2/blackboard.rest/api"
)

/**
 * The [TestCreateNewContent] function...
 */
func TestCreateNewContent (t *testing.T) {
  println ("Create a new content service instance.")

  if nil == blackboardRest.NewContent (mockContentService) {
    t.Error ("Creating new content service instance should not be nil reference.")
  }
}

/**
 * The [TestNewContentRequiresService] function...
 */
func TestNewContentRequiresService (t *testing.T) {
  println ("New content instance requires service reference.")

  if nil != blackboardRest.NewContent (nil) {
    t.Error ("Missing service should result in a nil reference.")
  }
}

/**
 * Mocked types and instances to run the above tests with.
 */
var mockContentService =
  api.NewService ("localhost", mockToken, new (_MockContentRoundTripper))

type _MockContentRoundTripper struct {
  http.RoundTripper
}

func (roundTripper *_MockContentRoundTripper) RoundTrip (
  request *http.Request,
) (*http.Response, error) {
  return nil, nil
}
