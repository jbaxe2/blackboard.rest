package blackboard_rest_test

import (
  "io/ioutil"
  "net/http"
  "strings"
  "testing"

  blackboardRest "github.com/jbaxe2/blackboard.rest"
  "github.com/jbaxe2/blackboard.rest/api"
)

/**
 * The [TestCreateNewCourses] function...
 */
func TestCreateNewCourses (t *testing.T) {
  println ("Create a new Courses service instance.")

  if nil == blackboardRest.NewCourses (mockCoursesService) {
    t.Error ("Creating a new Courses instance should not result in nil reference.")
  }
}

/**
 * The [TestNewCoursesRequiresService] function...
 */
func TestNewCoursesRequiresService (t *testing.T) {
  println ("New courses service requires a service instance.")

  if nil != blackboardRest.NewCourses (nil) {
    t.Error ("Missing service instance should result in nil reference.")
  }
}

/**
 * The [TestNewCoursesGetCourse] function...
 */
func TestNewCoursesGetCourse (t *testing.T) {
  println ("Retrieve a course from the REST API.")

  courses := blackboardRest.NewCourses (mockCoursesService)
  externalId := "wsu_jaxenroth_sandbox_1"
  course, err := courses.GetCourse ("externalId:" + externalId)

  if !(nil == err && course.ExternalId == externalId) {
    t.Error ("Retrieving a course should return the appropriate response.")
  }
}

/**
 * Mocked types and instances to run the above tests with.
 */
var mockCoursesService =
  api.NewService ("localhost", mockToken, mockCoursesRoundTripper)

var mockCoursesRoundTripper = new (_MockCoursesRoundTripper)

type _MockCoursesRoundTripper struct {
  http.RoundTripper
}

func (roundTripper *_MockCoursesRoundTripper) RoundTrip (
  request *http.Request,
) (*http.Response, error) {
  request.Response = &http.Response {
    Request: request,
    Header: make (http.Header),
  }

  if "GET" == request.Method && strings.Contains (request.URL.Path, "/courses/") {
    request.Response.StatusCode = 200

    request.Response.Body = ioutil.NopCloser (strings.NewReader (
      `{"id":"_121_1","courseId":"wsu_jaxenroth_sandbox_1","externalId":` +
      `"wsu_jaxenroth_sandbox_1","uuid":"asdf","name":"Joseph Axenroth Sandbox ` +
      `#1","dataSourceId":"plato.sis.courses","termId":"sandboxes_term",` +
      `"organization":false,"created":"2021-11-09T17:04:21.246Z"}`,
    ))
  } else {
    request.Response.StatusCode = 404

    request.Response.Body = ioutil.NopCloser (strings.NewReader (
      `{"status":404,"message":"Improper request"}`,
    ))
  }

  return request.Response, nil
}
