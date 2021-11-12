package blackboard_rest_test

import (
  "net/http"
  "testing"

  blackboardRest "github.com/jbaxe2/blackboard.rest"
  "github.com/jbaxe2/blackboard.rest/api"
)

/**
 * The [TestCreateCourseGroupUsers] function...
 */
func TestCreateCourseGroupUsers (t *testing.T) {
  println ("Create a new course group users instance.")

  if nil == blackboardRest.NewCourseGroupUsers (mockCourseGroupUsersService) {
    t.Error ("New course group users instance should not be nil reference.")
  }
}

/**
 * The [TestNewCourseGroupUsersRequiresService] function...
 */
func TestNewCourseGroupUsersRequiresService (t *testing.T) {
  println ("Creating a new course group users instance requires service.")

  if nil == blackboardRest.NewCourseGroupUsers (nil) {
    t.Error ("Missing service instance should result in nil reference.")
  }
}

/**
 * Mocked types and instances to run the above tests with.
 */
var mockCourseGroupUsersService =
  api.NewService ("localhost", mockToken, mockCourseGroupUsersRoundTripper)

var mockCourseGroupUsersRoundTripper = new (_MockCourseGroupUsersRoundTripper)

type _MockCourseGroupUsersRoundTripper struct {
  http.RoundTripper
}
