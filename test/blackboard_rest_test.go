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
  oauth2Tester := OAuth2Tester {t: tester.t}
  oauth2Tester.Run()

  systemTester := SystemTester {t: tester.t}
  systemTester.Run()

  usersTester := UsersTester {t: tester.t}
  usersTester.Run()

  coursesTester := CoursesTester {t: tester.t}
  coursesTester.Run()

  membershipsTester := CourseMembershipsTester {t: tester.t}
  membershipsTester.Run()

  termsTester := TermsTester {t: tester.t}
  termsTester.Run()

  courseGradesTester := CourseGradesTester {t: tester.t}
  courseGradesTester.Run()

  println()
}
