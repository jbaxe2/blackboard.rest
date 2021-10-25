package blackboard_rest

import (
  "net/url"
  "strings"

  "github.com/jbaxe2/blackboard.rest/_scaffolding/config"
  "github.com/jbaxe2/blackboard.rest/courses"
  "github.com/jbaxe2/blackboard.rest/oauth2"
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
  _BlackboardRest

  Courses
}

/**
 * The [GetCoursesInstance] function...
 */
func GetCoursesInstance (host string, accessToken oauth2.AccessToken) Courses {
  hostUri, _ := url.Parse (host)

  coursesService := new (_BbRestCourses)

  coursesService.host = *hostUri
  coursesService.accessToken = accessToken

  coursesService.service.SetHost (host)
  coursesService.service.SetAccessToken (accessToken)

  return coursesService
}

/**
 * The [GetCourse] method...
 */
func (restCourses *_BbRestCourses) GetCourse (
  courseId string,
) (courses.Course, error) {
  var course courses.Course

  endpoint := config.CoursesEndpoints["course"]
  endpoint = strings.Replace (endpoint, "{courseId}", courseId, -1)

  result, err := restCourses.service.Connector.SendBbRequest (
    endpoint, "GET", make (map[string]interface{}), 1,
  )

  if nil != err {
    return course, err
  }

  course = courses.NewCourse(result.(map[string]interface{}))

  return course, err
}
