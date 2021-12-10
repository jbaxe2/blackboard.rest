package system_test

import (
  "testing"

  "github.com/jbaxe2/blackboard.rest/system"
)

/**
 * The [TestCreateNewVersionInfo] function...
 */
func TestCreateNewVersionInfo (t *testing.T) {
  println ("Create a new version info instance.")

  versionInfo := system.NewVersionInfo (rawVersionInfo)

  if !(versionInfo.Learn.Build == "rel.11+732fc84" &&
       versionInfo.Learn.Major == 3900) {
    t.Error ("New version info instance should have expected values.")
  }
}

/**
 * The [TestCreateNewPrivacyPolicy] function...
 */
func TestCreateNewPrivacyPolicy (t *testing.T) {
  println ("Create a new privacy policy instance.")

  privacyPolicy := system.NewPrivacyPolicy (rawPrivacyPolicy)

  if !(privacyPolicy.Blackboard.String() == rawPrivacyPolicy["blackboard"] &&
       privacyPolicy.Institution.String() == rawPrivacyPolicy["institution"]) {
    t.Error ("New privacy policy instance should have expected values.")
  }
}

/**
 * Mocked instances to run the above tests with.
 */
var rawVersionInfo = map[string]interface{} {
  "learn": map[string]interface{} {
    "major": 3900,
    "minor": 28,
    "patch": 0,
    "build": "rel.11+732fc84",
  },
}

var rawPrivacyPolicy = map[string]interface{} {
  "blackboard": "www.blackboard.com/privacy",
  "institution": "localhost",
}
