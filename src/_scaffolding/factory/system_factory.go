package factory

import "github.com/jbaxe2/blackboard.rest.go/src/system"

/**
 * The [NewVersionInfo] function...
 */
func NewVersionInfo (rawVersionInfo map[string]interface{}) system.VersionInfo {
  return system.VersionInfo {
    Learn: _parseVersion (rawVersionInfo["learn"].(map[string]interface{})),
  }
}

/**
 * The [_parseVersion] function...
 */
func _parseVersion (rawVersion map[string]interface{}) system.Version {
  return system.Version {
    Major: rawVersion["major"].(float64),
    Minor: rawVersion["minor"].(float64),
    Patch: rawVersion["patch"].(float64),
    Build: rawVersion["build"].(string),
  }
}

/**
 * The [NewPrivacyPolicy] function...
 */
func NewPrivacyPolicy (
  rawPrivacyPolicy map[string]interface{},
) system.PrivacyPolicy {
  return system.PrivacyPolicy {
    Blackboard: rawPrivacyPolicy["blackboard"].(string),
    Institution: rawPrivacyPolicy["institution"].(string),
  }
}
