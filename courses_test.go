package blackboard_rest_test

import (
  "testing"

  blackboardRest "github.com/jbaxe2/blackboard.rest"
)

/**
 * The [TestCreateNewCourses] function...
 */
func TestCreateNewCourses (t *testing.T) {
  println ("Create a new Courses service instance.")

  if nil == blackboardRest.NewCourses (nil) {
    t.Error ("Creating a new Courses instance should not result in nil reference.")
  }
}
