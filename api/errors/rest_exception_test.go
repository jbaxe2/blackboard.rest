package errors_test

import (
  "testing"

  "github.com/jbaxe2/blackboard.rest/api/errors"
)

/**
 * The [TestCreateNewRestException] function...
 */
func TestCreateNewRestException (t *testing.T) {
  println ("Create a new REST exception instance.")

  if nil == errors.NewRestException ("", "", "", "", nil) {
    t.Error ("Creating a new REST exception should not result in nil reference.")
  }
}
