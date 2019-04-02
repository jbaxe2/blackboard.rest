package test

import "testing"

/**
 * The [UsersTester] type...
 */
type UsersTester struct {
  t *testing.T
}

/**
 * The [Run] method...
 */
func (tester *UsersTester) Run() {
  print ("Users:\n")

  _testGetUsersInstance (tester.t)
}

/**
 * The [_testGetUsersInstance] function...
 */
func _testGetUsersInstance (t *testing.T) {
  print ("Obtain a valid Users service instance.\n")
}
