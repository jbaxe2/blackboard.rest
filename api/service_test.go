package api_test

import (
  "testing"

  "github.com/jbaxe2/blackboard.rest/api"
)

/**
 * The [TestCreateNewService] function...
 */
func TestCreateNewService (t *testing.T) {
  println ("Create a new service instance.")

  if nil == api.NewService ("localhost", nil) {
    t.Error ("Creating a new service instance should not result in nil reference.")
    t.FailNow()
  }
}

/**
 * The [TestNewServiceRequiresHost] function...
 */
func TestNewServiceRequiresHost (t *testing.T) {
  println ("Creating a new service instance requires a host.")

  if nil != api.NewService ("", nil) {
    t.Error ("Missing host should result in nil service reference.")
    t.FailNow()
  }
}
