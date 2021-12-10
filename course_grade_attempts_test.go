package blackboard_rest_test

import (
  "net/http"
  "testing"

  blackboardRest "github.com/jbaxe2/blackboard.rest"
  "github.com/jbaxe2/blackboard.rest/api"
)

/**
 * The [TestCreateNewCourseGradeAttempts] function...
 */
func TestCreateNewCourseGradeAttempts (t *testing.T) {
  println ("Create a new course grade attempts instance.")

  if nil == blackboardRest.NewCourseGradeAttempts (mockCourseGradeAttemptsService) {
    t.Error ("Creating a new course grades attempts should not be nil reference.")
  }
}

/**
 * The [TestNewCourseGradeAttemptsRequiresService] function...
 */
func TestNewCourseGradeAttemptsRequiresService (t *testing.T) {
  println ("Creating a new course grade attempts requires service instance.")

  if nil != blackboardRest.NewCourseGradeAttempts (nil) {
    t.Error ("Missing course grades attempts should result in nil reference.")
  }
}

/**
 * Mocked types and instances to run the above tests with.
 */
var mockCourseGradeAttemptsService =
  api.NewService ("localhost", mockToken, new (_MockCourseGradeAttemptsRoundTripper))

type _MockCourseGradeAttemptsRoundTripper struct {
  http.RoundTripper
}
