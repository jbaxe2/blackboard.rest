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

/**
 * The [TestNewRestExceptionCanBeCreatedFromRawStructure] function...
 */
func TestNewRestExceptionCanBeCreatedFromRawStructure (t *testing.T) {
  println ("New REST exception can be created from raw structure.")

  var rawException = map[string]interface{} {
    "status": 401,
    "code": "401",
    "message": "Unauthorized",
    "developerMessage": "",
    "extraInfo": "",
  }

  if nil == errors.NewRestExceptionFromRaw (rawException) {
    t.Error ("New REST exception from raw structure should not be nil reference.")
  }
}

/**
 * The [TestNewRestExceptionFromNilRawStructureIsNil] function...
 */
func TestNewRestExceptionFromNilRawStructureIsNil (t *testing.T) {
  println ("New REST exception from nil raw structure results in nil reference.")

  if nil != errors.NewRestExceptionFromRaw (nil) {
    t.Error ("Nil raw structure should result in nil reference for REST exception.")
  }
}

/**
 * The [TestNewRestExceptionCanOmitNonRequiredFields] function...
 */
func TestNewRestExceptionCanOmitNonRequiredFields (t *testing.T) {
  println ("Omitted non-required fields should result in non-nil REST exception.")

  var rawException = map[string]interface{} {
    "status": 401,
    "message": "Unauthorized",
    "extraInfo": "",
  }

  if nil == errors.NewRestExceptionFromRaw (rawException) {
    t.Error ("Incomplete raw structure should not result in nil reference.")
  }
}

/**
 * The [TestNewRestExceptionFromRawHasPertinentInformation] function...
 */
func TestNewRestExceptionFromRawHasPertinentInformation (t *testing.T) {
  println ("New REST exception from raw structure retains info used to create it.")

  var rawException = map[string]interface{} {
    "status": 401,
    "code": "401",
    "message": "Unauthorized",
    "developerMessage": "",
    "extraInfo": "https://localhost/extra/info",
  }

  restException := errors.NewRestExceptionFromRaw (rawException)

  if !(restException.Status() == 401 &&
       restException.Code() == rawException["code"] &&
       restException.Message() == rawException["message"] &&
       restException.DeveloperMessage() == rawException["developerMessage"] &&
       restException.ExtraInfo().String() == rawException["extraInfo"]) {
    t.Error (
      "New REST exception from raw structure should retain info used to create it.",
    )
  }
}
