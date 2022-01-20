package system

import (
  "net/url"

  "github.com/jbaxe2/blackboard.rest/utils"
)

/**
 * The [NewVersionInfo] function...
 */
func NewVersionInfo (rawVersionInfo map[string]interface{}) VersionInfo {
  return VersionInfo {
    Learn: _parseVersion (rawVersionInfo["learn"].(map[string]interface{})),
  }
}

/**
 * The [_parseVersion] function...
 */
func _parseVersion (rawVersion map[string]interface{}) Version {
  build, _ := rawVersion["build"].(string)

  return Version {
    Major: utils.NormalizeNumeric (rawVersion["major"]),
    Minor: utils.NormalizeNumeric (rawVersion["minor"]),
    Patch: utils.NormalizeNumeric (rawVersion["patch"]),
    Build: build,
  }
}

/**
 * The [NewPrivacyPolicy] function...
 */
func NewPrivacyPolicy (rawPrivacyPolicy map[string]interface{}) PrivacyPolicy {
  blackboard, _ := url.Parse (rawPrivacyPolicy["blackboard"].(string))
  institution, _ := url.Parse (rawPrivacyPolicy["institution"].(string))

  return PrivacyPolicy {
    Blackboard: blackboard,
    Institution: institution,
  }
}
