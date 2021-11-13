package blackboard_rest_test

import (
  "net/http"
  "testing"

  blackboardRest "github.com/jbaxe2/blackboard.rest"
  "github.com/jbaxe2/blackboard.rest/api"
)

/**
 * The [TestCreateNewCourseGroups] function...
 */
func TestCreateNewCourseGroups (t *testing.T) {
  println ("Create a new course groups instance.")

  if nil == blackboardRest.NewCourseGroups (mockCourseGroupsService) {
    t.Error ("Creating new course groups instance should not be nil reference.")
  }
}

/**
 * The [TestNewCourseGroupsRequiresService] function...
 */
func TestNewCourseGroupsRequiresService (t *testing.T) {
  println ("Creating new course groups instance requires a service instance.")

  if nil != blackboardRest.NewCourseGroups (nil) {
    t.Error ("Missing service instance should result in nil reference.")
  }
}

/**
 * Mocked types and instances to run the above tests with.
 */
var mockCourseGroupsService =
  api.NewService ("localhost", mockToken, mockCourseGroupsRoundTripper)

var mockCourseGroupsRoundTripper = new (_MockCourseGroupsRoundTripper)

type _MockCourseGroupsRoundTripper struct {
  http.RoundTripper
}
