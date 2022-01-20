package courses_test

import (
  "testing"

  "github.com/jbaxe2/blackboard.rest/courses"
)

/**
 * The [TestCreateNewCourseViaFactory] function...
 */
func TestCreateNewCourseViaFactory (t *testing.T) {
  println ("Create a new course instance via the factory.")

  course := courses.NewCourse (rawCourse1)

  if course.CourseId != rawCourse1["courseId"] {
    t.Error ("Creating a new course should have expected results.")
  }
}

/**
 * The [TestCreateNewCoursesViaFactory] function...
 */
func TestCreateNewCoursesViaFactory (t *testing.T) {
  println ("Create multiple course instances via the factory.")

  newCourses := courses.NewCourses (rawCourses)

  if !(3 == len (newCourses) && newCourses[0].CourseId == rawCourse1["courseId"]) {
    t.Error ("Creating the newCourses should have expected results.")
  }
}

/**
 * Mocked instances to run the above tests with.
 */
var rawCourses = []map[string]interface{} {rawCourse1, rawCourse2, rawCourse3}

var rawCourse1 = map[string]interface{} {
  "id":"_1_1", "courseId":"wsu_course_1", "externalId":"wsu_course_1",
  "uuid":"asdf1", "name":"Course #1", "dataSourceId":"plato.sis.courses",
  "termId":"2021fall", "organization":false, "created":"2021-11-09T17:04:21.246Z",
}

var rawCourse2 = map[string]interface{} {
  "id":"_2_1", "courseId":"wsu_course_2", "externalId":"wsu_course_2",
  "uuid":"asdf2", "name":"Course #2", "dataSourceId":"plato.sis.courses",
  "termId":"2021fall", "organization":false, "created":"2021-11-09T17:04:21.246Z",
}

var rawCourse3 = map[string]interface{} {
  "id":"_3_1", "courseId":"wsu_course_3", "externalId":"wsu_course_3",
  "uuid":"asdf3", "name":"Course #3", "dataSourceId":"plato.sis.courses",
  "termId":"2021fall", "organization":false, "created":"2021-11-09T17:04:21.246Z",
}
