package blackboard_rest

import (
  "github.com/jbaxe2/blackboard.rest/api"
  "github.com/jbaxe2/blackboard.rest/system"
)

/**
 * The [System] interface provides the base interface for interacting with the
 * REST API's system service.
 */
type System interface {
  GetPolicies() (system.PrivacyPolicy, error)

  GetVersion() (system.VersionInfo, error)
}

/**
 * The [_System] type implements the System interface.
 */
type _System struct {
  service api.Service

  System
}

/**
 * The [NewSystem] function creates a new system instance.
 */
func NewSystem (service api.Service) System {
  if nil == service {
    return nil
  }

  return &_System {
    service: service,
  }
}

/**
 * The [GetVersion] method retrieves the version information about the Learn
 * installation for which the REST API is interacting with.
 */
func (systems *_System) GetVersion() (system.VersionInfo, error) {
  rawVersion, err := systems.service.Request (string (api.Version), "GET", nil, 1)

  if nil != err {
    return system.VersionInfo{}, err
  }

  return system.NewVersionInfo (rawVersion), nil
}
