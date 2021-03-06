package test

import (
  "testing"

  blackboardRest "github.com/jbaxe2/blackboard.rest"
  "github.com/jbaxe2/blackboard.rest/_scaffolding/config"
  "github.com/jbaxe2/blackboard.rest/terms"
)

/**
 * The [TermsTester] type...
 */
type TermsTester struct {
  t *testing.T

  Testable
}

/**
 * The [Run] method...
 */
func (tester *TermsTester) Run() {
  println ("\nTerms:")

  _testObtainValidTermsInstance (tester.t)
  _testGetAllTerms (tester.t)
  _testGetTermByPrimaryId (tester.t)
}

/**
 * The [_getTermsInstance] function...
 */
func _getTermsInstance() blackboardRest.Terms {
  authorizer := TestersAuthorizer{}
  _ = authorizer.AuthorizeForTests()

  return blackboardRest.GetTermsInstance (
    config.Host, authorizer.accessToken,
  )
}

/**
 * The [_testObtainValidTermsInstance] function...
 */
func _testObtainValidTermsInstance (t *testing.T) {
  println ("Obtain a valid Terms service instance.")

  termsService := _getTermsInstance()

  if nil == termsService {
    t.Error ("Obtaining a valid Terms service instance failed.\n")
    t.FailNow()
  }
}

/**
 * The [_testGetAllTerms] function...
 */
func _testGetAllTerms (t *testing.T) {
  println ("Get all of the terms.")

  termsService := _getTermsInstance()

  theTerms, err := termsService.GetTerms()

  if (nil == theTerms) || (nil != err) {
    t.Error ("Failed to retrieve the list of terms.\n")
    t.FailNow()
  }

  if 0 == len (theTerms) {
    t.Error ("Retrieved an empty list of terms that should not be empty.\n")
    t.FailNow()
  }
}

/**
 * The [_testGetTerm] function...
 */
func _testGetTermByPrimaryId (t *testing.T) {
  println ("Get a term by its primary (internal) ID.")

  termsService := _getTermsInstance()

  term, err := termsService.GetTerm ("_380_1")

  if (terms.Term{} == term) || (nil != err) {
    t.Error ("Failed to obtain the specified term.\n")
    t.Fail()
  }

  if "_380_1" != term.Id {
    t.Error ("The retrieved term does not match what was specified.")
    t.Fail()
  }
}
