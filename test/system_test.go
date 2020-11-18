package test

import (
  "testing"

  blackboardRest "github.com/jbaxe2/blackboard.rest"
  "github.com/jbaxe2/blackboard.rest/_scaffolding/config"
  "github.com/jbaxe2/blackboard.rest/system"
)

/**
 * The [SystemTester] type...
 */
type SystemTester struct {
  t *testing.T

  Testable
}

/**
 * The [Run] method...
 */
func (tester *SystemTester) Run() {
  println ("\nSystem:")

  _testGetValidSystemInstance (tester.t)
  _testGetVersion (tester.t)
}

/**
 * The [_getSystemInstance] function...
 */
func _getSystemInstance() blackboardRest.System {
  authorizer := TestersAuthorizer{}
  _ = authorizer.AuthorizeForTests()

  return blackboardRest.GetSystemInstance (
    config.Host, authorizer.accessToken,
  )
}

/**
 * The [_testGetValidSystemInstance] function...
 */
func _testGetValidSystemInstance (t *testing.T) {
  println ("Obtain a valid System service instance.")

  systemService := _getSystemInstance()

  if nil == systemService {
    t.Error ("Failed to obtain a valid System service instance.")
    t.FailNow()
  }
}

/**
 * The [_testGetVersion] function...
 */
func _testGetVersion (t *testing.T) {
  println ("Get the Blackboard Learn version information.")

  systemService := _getSystemInstance()

  versionInfo, err := systemService.GetVersion()

  if (system.VersionInfo{} == versionInfo) || (nil != err) {
    t.Error ("Failed to obtain the Blackboard Learn version information.")
    t.FailNow()
  }
}
