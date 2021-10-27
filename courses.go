package blackboard_rest

import (
  "github.com/jbaxe2/blackboard.rest/courses"
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
