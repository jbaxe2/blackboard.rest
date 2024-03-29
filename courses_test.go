package blackboard_rest_test

import (
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
 * The [TestNewCoursesGetCourses] function...
 */
func TestNewCoursesGetCourses (t *testing.T) {
  println ("Retrieve multiple courses from the REST API.")

  courses := blackboardRest.NewCourses (mockCoursesService)
  newCourses, err := courses.GetCourses()

  if !(nil == err && 4 == len (newCourses)) {
    t.Error ("Retrieving courses should return the appropriate responses.")
  }
}

/**
 * The [TestNewCoursesGetCoursesByTerm] function...
 */
func TestNewCoursesGetCoursesByTerm (t *testing.T) {
  println ("Retrieve multiple courses from the REST API belonging to a same term.")

  courses := blackboardRest.NewCourses (mockCoursesService)
  termedCourses, err := courses.GetCoursesByTerm ("2021fall")

  if !(nil == err && 3 == len (termedCourses) && "2021fall" == termedCourses[0].TermId) {
    t.Error ("Retrieving termed courses should return the appropriate responses.")
  }
}

/**
 * The [TestNewCoursesGetCourse] function...
 */
func TestNewCoursesGetCourse (t *testing.T) {
  println ("Retrieve a course from the REST API.")

  courses := blackboardRest.NewCourses (mockCoursesService)
  externalId := "jaxenroth_sandbox_1"
  course, err := courses.GetCourse ("externalId:" + externalId)

  if !(nil == err && course.ExternalId == externalId) {
    t.Error ("Retrieving a course should return the appropriate response.")
  }
}

/**
 * Mocked types and instances to run the above tests with.
 */
var mockCoursesService =
  api.NewService ("localhost", mockToken, new (_MockCoursesRoundTripper))

type _MockCoursesRoundTripper struct {
  http.RoundTripper
}

func (roundTripper *_MockCoursesRoundTripper) RoundTrip (
  request *http.Request,
) (*http.Response, error) {
  conditions := []bool {
    "GET" == request.Method && request.URL.Query().Has ("termId"),
    "GET" == request.Method && strings.Contains (request.URL.Path, "/courses/"),
    "GET" == request.Method && strings.Contains (request.URL.Path, "/courses"),
  }

  responseBodies := []string {rawCourses, sandboxCourse, allCourses}

  builder := NewResponseBuilder (conditions, responseBodies)

  return builder.Build (request), nil
}

const allCourses =
  `{"results":[` + course1 + `,` + course2 + `,` + course3 + `,` + sandboxCourse + `]}`

const rawCourses = `{"results":[` + course1 + `,` + course2 + `,` + course3 + `]}`

const course1 = `{"id":"_1_1","courseId":"course_id_1","externalId":"course_id_1",` +
  `"uuid":"asdf1","name":"Course #1","dataSourceId":"data.source.courses.id","termId":` +
  `"2021fall","organization":false,"created":"2021-11-09T17:04:21.246Z"}`

const course2 = `{"id":"_2_1","courseId":"course_id_2","externalId":"course_id_2",` +
  `"uuid":"asdf2","name":"Course #2","dataSourceId":"data.source.courses.id","termId":` +
  `"2021fall","organization":false,"created":"2021-11-09T17:04:21.246Z"}`

const course3 = `{"id":"_3_1","courseId":"course_id_3","externalId":"course_id_3",` +
  `"uuid":"asdf3","name":"Course #3","dataSourceId":"data.source.courses.id","termId":` +
  `"2021fall","organization":false,"created":"2021-11-09T17:04:21.246Z"}`

const sandboxCourse = `{"id":"_121_1","courseId":"jaxenroth_sandbox_1",` +
  `"externalId":"jaxenroth_sandbox_1","uuid":"asdf","name":"Joseph Axenroth ` +
  `Sandbox #1","dataSourceId":"data.source.sandbox.id","termId":"sandboxes_term",` +
  `"organization":false,"created":"2021-11-09T17:04:21.246Z"}`
