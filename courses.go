package blackboard_rest

import (
  "strings"

  "github.com/jbaxe2/blackboard.rest/api"
  "github.com/jbaxe2/blackboard.rest/courses"
  "github.com/jbaxe2/blackboard.rest/utils"
)

/**
 * The [Courses] interface provides the base type for interacting with the REST
 * API's courses endpoints.
 */
type Courses interface {
  GetCourses() ([]courses.Course, error)

  GetCoursesByTerm (termId string) ([]courses.Course, error)

  CreateCourse (course courses.Course) error

  GetCourse (courseId string) (courses.Course, error)

  UpdateCourse (courseId string, course courses.Course) error

  GetChildren (courseId string) ([]courses.CourseChild, error)

  GetChild (courseId string, childCourseId string) (courses.CourseChild, error)

  AddChildCourse (
    courseId string, childCourseId string, ignoreEnrollmentErrors bool,
  ) error

  CopyCourse (courseId string, newCourseId string) error

  GetCrossListSet (courseId string) ([]courses.CourseChild, error)

  GetTask (courseId string, taskId string) (courses.CourseTask, error)
}

/**
 * The [_Courses] type implements the Courses interface.
 */
type _Courses struct {
  service api.Service

  Courses
}

/**
 * The [NewCourses] function creates a new courses instance.
 */
func NewCourses (service api.Service) Courses {
  if nil == service {
    return nil
  }

  return &_Courses {
    service: service,
  }
}

/**
 * The [GetCourses] method retrieves a slice of courses.  Request options should
 * be used to slim down the groups of courses, such as by availability, name,
 * modified, etc.  Paging is currently not provided by this library directly, and
 * should also be performed via the request options.
 */
func (course *_Courses) GetCourses() ([]courses.Course, error) {
  rawCourses, err := course.service.Request (string (api.Courses), "GET", nil, 3)

  if nil != err {
    return nil, err
  }

  return courses.NewCourses (
    utils.NormalizeRawResponse (rawCourses["results"].([]interface{})),
  ), nil
}

/**
 * The [GetCourses] method retrieves a slice of courses for a particular term,
 * as provided by the term ID.  Note this is a convenience method for retrieving
 * group of courses based on a common criteria (some term), and does not reflect
 * a direct standalone endpoint in the courses REST API.
 */
func (course *_Courses) GetCoursesByTerm (termId string) ([]courses.Course, error) {
  course.service.SetRequestOption ("termId", termId)

  rawCourses, err := course.service.Request (string (api.Courses), "GET", nil, 3)

  if nil != err {
    return nil, err
  }

  return courses.NewCourses (
    utils.NormalizeRawResponse (rawCourses["results"].([]interface{})),
  ), nil
}

/**
 * The [GetCourse] method retrieves information about a single course based on
 * the provided course ID.
 */
func (course *_Courses) GetCourse (courseId string) (courses.Course, error) {
  endpoint := strings.Replace (string (api.Course), "{courseId}", courseId, 1)
  rawCourse, err := course.service.Request (endpoint, "GET", nil, 3)

  if nil != err {
    return courses.Course{}, err
  }

  return courses.NewCourse (rawCourse), nil
}
