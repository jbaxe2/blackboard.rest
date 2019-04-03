package test

import "testing"

/**
 * The [TestBlackboardRest] function...
 */
func TestBlackboardRest (t *testing.T) {
  blackboardRestTester := BlackboardRestTester{t: t}
  blackboardRestTester.Run()
}

/**
 * The [Testable] interface...
 */
type Testable interface {
  Run()
}

/**
 * The [BlackboardRestTester] type...
 */
type BlackboardRestTester struct {
  t *testing.T

  Testable
}

/**
 * The [Run] method...
 */
func (tester *BlackboardRestTester) Run() {
  oauth2Tester := OAuth2Tester{t: tester.t}
  oauth2Tester.Run()

  usersTester := UsersTester{t: tester.t}
  usersTester.Run()
}
