package blackboard_rest_test

import (
  "net/http"
  "testing"

  blackboardRest "github.com/jbaxe2/blackboard.rest"
  "github.com/jbaxe2/blackboard.rest/api"
)

/**
 * The [TestCreateNewCourseGrades] function...
 */
func TestCreateNewCourseGrades (t *testing.T) {
  println ("Create a new course grades instance.")

  if nil == blackboardRest.NewCourseGrades (mockCourseGradeService) {
    t.Error ("Creating a new course grades instance should not be nil reference.")
  }
}

/**
 * The [TestNewCourseGradesRequiresService] function...
 */
func TestNewCourseGradesRequiresService (t *testing.T) {
  println ("Creating a new course grades instance requires service reference.")

  if nil != blackboardRest.NewCourseGrades (nil) {
    t.Error ("Missing service instance should result in nil reference.")
  }
}

/**
 * Mocked types and instances to run the above tests with.
 */
var mockCourseGradeService =
  api.NewService ("localhost", mockToken, mockCourseGradesRoundTripper)

var mockCourseGradesRoundTripper = new (_MockCourseGradesRoundTripper)

type _MockCourseGradesRoundTripper struct {
  http.RoundTripper
}
