package errors_test

import (
  "testing"

  "github.com/jbaxe2/blackboard.rest/api/errors"
)

/**
 * The [TestCreateNewOAuth2Error] function...
 */
func TestCreateNewOAuth2Error (t *testing.T) {
  println ("Create a new OAuth2 Error instance.")

  if nil == errors.NewOAuth2Error ("invalid_client", "description") {
    t.Error (
      "Creating a new OAuth2 Error instance should not result in nil reference.",
    )
  }
}

/**
 * The [TestNewOAuth2ErrorRequiresCode] function...
 */
func TestNewOAuth2ErrorRequiresCode (t *testing.T) {
  println ("Creating a new OAuth2 Error instance requires a code.")

  if nil != errors.NewOAuth2Error ("", "description") {
    t.Error ("Missing code should result in nil reference.")
  }
}

/**
 * The [TestNewOAuth2ErrorDescriptionCanBeEmpty] function...
 */
func TestNewOAuth2ErrorDescriptionCanBeEmpty (t *testing.T) {
  println ("Creating a new OAuth2 Error can have empty description.")

  if nil == errors.NewOAuth2Error ("invalid_client", "") {
    t.Error ("New OAuth2 Error can have an empty description.")
  }
}

/**
 * The [TestNewOAuth2ErrorDescriptionCanBeNonEmpty] function...
 */
func TestNewOAuth2ErrorDescriptionCanBeNonEmpty (t *testing.T) {
  println ("Creating a new OAuth2 Error can have non-empty description.")

  if nil == errors.NewOAuth2Error ("invalid_client", "description") {
    t.Error ("New OAuth2 Error can have a non-empty description.")
  }
}

/**
 * The [TestNewOAuth2ErrorCanHaveValidCode] function...
 */
func TestNewOAuth2ErrorCanHaveValidCode (t *testing.T) {
  println ("Creating a new OAuth2 Error can have a valid code.")

  if nil == errors.NewOAuth2Error ("invalid_client", "") {
    t.Error ("New OAuth2 Error with valid code should not be nil reference.")
  }
}

/**
 * The [TestNewOAuth2ErrorMustHaveValidCode] function...
 */
func TestNewOAuth2ErrorMustHaveValidCode (t *testing.T) {
  println ("Creating a new OAuth2 Error must have a valid code.")

  if nil != errors.NewOAuth2Error ("code", "description") {
    t.Error ("New OAuth2 Error with valid code should not be nil reference.")
  }
}

/**
 * The [TestNewOAuth2ErrorHasPertinentInformation] function...
 */
func TestNewOAuth2ErrorHasPertinentInformation (t *testing.T) {
  println ("New OAuth2 Error retains the information used to create it.")

  code := "invalid_client"
  description := "error description"

  oauth2Error := errors.NewOAuth2Error (code, description)

  if !(oauth2Error.Code() == code && oauth2Error.Description() == description) {
    t.Error ("The OAuth2 Error should retain the information used to create it.")
  }
}
