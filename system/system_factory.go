package system

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
  return Version{
    Major: rawVersion["major"].(float64),
    Minor: rawVersion["minor"].(float64),
    Patch: rawVersion["patch"].(float64),
    Build: rawVersion["build"].(string),
  }
}

/**
 * The [NewPrivacyPolicy] function...
 */
func NewPrivacyPolicy (rawPrivacyPolicy map[string]interface{}) PrivacyPolicy {
  return PrivacyPolicy {
    Blackboard: rawPrivacyPolicy["blackboard"].(string),
    Institution: rawPrivacyPolicy["institution"].(string),
  }
}
