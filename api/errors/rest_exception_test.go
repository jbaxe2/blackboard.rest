package errors_test

import (
  "net/url"
  "strconv"
  "testing"

  "github.com/jbaxe2/blackboard.rest/api/errors"
)

/**
 * The [TestCreateNewRestException] function...
 */
func TestCreateNewRestException (t *testing.T) {
  println ("Create a new REST exception instance.")

  if nil == errors.NewRestException (400, "", "error message", "", nil) {
    t.Error ("Creating a new REST exception should not result in nil reference.")
  }
}

/**
 * The [TestNewRestExceptionRequiresAppropriateStatus] function...
 */
func TestNewRestExceptionRequiresAppropriateStatus (t *testing.T) {
  println ("New REST exception instance requires an appropriate status.")

  if nil != errors.NewRestException (0, "", "error message", "", nil) {
    t.Error ("Inappropriate status code value should result in nil reference.")
  }
}

/**
 * The [TestNewRestExceptionRequiresMessage] function...
 */
func TestNewRestExceptionRequiresMessage (t *testing.T) {
  println ("New REST exception instance requires an error message.")

  if nil != errors.NewRestException (400, "", "", "", nil) {
    t.Error ("Missing error message should result in nil reference.")
  }
}

/**
 * The [TestNewRestExceptionHasPertinentInformation] function...
 */
func TestNewRestExceptionHasPertinentInformation (t *testing.T) {
  println ("New REST exception retains the information used to create it.")

  status := 401
  code := "401"
  message := "Unauthorized"
  developerMessage := "Developer message information"
  extraInfo, _ := url.Parse ("localhost")

  exception :=
    errors.NewRestException (status, code, message, developerMessage, extraInfo)

  if !(exception.Status() == status && exception.Code() == code &&
       exception.Message() == message &&
       exception.DeveloperMessage() == developerMessage &&
       exception.ExtraInfo() == extraInfo) {
    t.Error ("New REST exception should retain the information used to create it.")
  }
}

/**
 * The [TestNewRestExceptionProvidesAppropriateErrorMessage] function...
 */
func TestNewRestExceptionProvidesAppropriateErrorMessage (t *testing.T) {
  println ("New REST exception has the appropriate error message.")

  status := 401
  message := "Unauthorized"
  errorMsg := "(" + strconv.Itoa (status) + ") " + message
  exception := errors.NewRestException (status, "", message, "", nil)

  if exception.Error() != errorMsg {
    t.Error ("The provided error message is not what was expected.")
  }
}
