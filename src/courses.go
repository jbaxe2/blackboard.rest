package blackboard_rest

import (
  "github.com/jbaxe2/blackboard.rest.go/src/_scaffolding/config"
  error2 "github.com/jbaxe2/blackboard.rest.go/src/_scaffolding/error"
  "github.com/jbaxe2/blackboard.rest.go/src/_scaffolding/factory"
  "github.com/jbaxe2/blackboard.rest.go/src/_scaffolding/services"
  "github.com/jbaxe2/blackboard.rest.go/src/courses"
  "github.com/jbaxe2/blackboard.rest.go/src/oauth2"
  "net/url"
  "strings"
)

/**
 * The [Courses] interface...
 */
type Courses interface {
  GetCourses() []courses.Course

  CreateCourse (course courses.Course)

  GetCourse (courseId string) (courses.Course, error)

  UpdateCourse (courseId string, course courses.Course)

  GetChildren (courseId string) []courses.CourseChild

  GetChild (courseId string, childCourseId string) courses.CourseChild

  AddChildCourse (
    courseId string, childCourseId string, ignoreEnrollmentErrors bool,
  )

  CopyCourse (courseId string, newCourseId string)

  GetCrossListSet (courseId string) []courses.CourseChild

  GetTask (courseId string, taskId string) courses.CourseTask
}

/**
 * The [_BbRestCourses] type...
 */
type _BbRestCourses struct {
  host url.URL

  accessToken oauth2.AccessToken

  service services.BlackboardRestService

  Courses
}

func (restCourses *_BbRestCourses) Host() url.URL {
  return restCourses.host
}

func (restCourses *_BbRestCourses) AccessToken() oauth2.AccessToken {
  return restCourses.accessToken
}

/**
 * The [GetCoursesInstance] function...
 */
func GetCoursesInstance (host string, accessToken oauth2.AccessToken) Courses {
  hostUri, _ := url.Parse (host)

  coursesService := &_BbRestCourses {host: *hostUri, accessToken: accessToken}
  coursesService.service.SetAccessToken (accessToken)

  return coursesService
}

/**
 * The [GetCourse] method...
 */
func (restCourses *_BbRestCourses) GetCourse (courseId string) (courses.Course, error) {
  var course courses.Course
  var err error
  var result interface{}

  endpoint := config.CoursesEndpoints["course"]
  endpoint = strings.Replace (endpoint, "{courseId}", courseId, -1)

  result, err = restCourses.service.Connector.SendBbRequest (
    endpoint, "GET", make (map[string]interface{}), 1,
  )

  if (nil != err) && (error2.RestError{} != err) {
    return course, err.(error2.CoursesError)
  }

  course = factory.NewCourse (result.(map[string]interface{}))

  return course, err
}
