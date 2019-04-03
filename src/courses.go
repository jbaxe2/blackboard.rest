package blackboard_rest

import "github.com/jbaxe2/blackboard.rest.go/src/courses"

/**
 * The [Courses] interface...
 */
type Courses interface {
  GetCourses() []courses.Course

  CreateCourse (course courses.Course)

  GetCourse (courseId string) courses.Course

  UpdateCourse (courseId string, course courses.Course)
}
