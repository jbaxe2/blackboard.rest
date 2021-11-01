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

  if nil == errors.NewRestException (400, "", "", "", nil) {
    t.Error ("Creating a new REST exception should not result in nil reference.")
  }
}

/**
 * The [TestNewRestExceptionRequiresAppropriateStatus] function...
 */
func TestNewRestExceptionRequiresAppropriateStatus (t *testing.T) {
  println ("New REST exception instance requires an appropriate status.")

  if nil != errors.NewRestException (0, "", "", "", nil) {
    t.Error ("Inappropriate status code value should result in nil reference.")
  }
}
