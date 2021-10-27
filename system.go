package blackboard_rest

import (
  "github.com/jbaxe2/blackboard.rest/system"
)

/**
 * The [System] interface...
 */
type System interface {
  GetPolicies() (system.PrivacyPolicy, error)

  GetVersion() (system.VersionInfo, error)
}
