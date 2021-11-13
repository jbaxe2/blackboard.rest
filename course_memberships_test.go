package blackboard_rest_test

import (
  "net/http"
  "testing"

  blackboardRest "github.com/jbaxe2/blackboard.rest"
  "github.com/jbaxe2/blackboard.rest/api"
)

/**
 * The [TestCreateNewCourseMemberships] function...
 */
func TestCreateNewCourseMemberships (t *testing.T) {
  println ("Create a new course memberships instance.")

  if nil == blackboardRest.NewCourseMemberships (mockCourseMembershipsService) {
    t.Error ("Creating new course memberships instance should not be nil reference.")
  }
}

/**
 * The [TestNewCourseMembershipsRequiresService] function...
 */
func TestNewCourseMembershipsRequiresService (t *testing.T) {
  println ("Creating new course memberships instance requires a service instance.")

  if nil != blackboardRest.NewCourseMemberships (nil) {
    t.Error ("Missing service instance should result in nil reference.")
  }
}

/**
 * Mocked types and instances to run the above tests with.
 */
var mockCourseMembershipsService =
  api.NewService ("localhost", mockToken, mockCourseMembershipsRoundTripper)

var mockCourseMembershipsRoundTripper = new (_MockCourseMembershipsRoundTripper)

type _MockCourseMembershipsRoundTripper struct {
  http.RoundTripper
}
