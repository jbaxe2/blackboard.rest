package blackboard_rest

import (
  "net/url"

  "github.com/jbaxe2/blackboard.rest/_scaffolding/config"
  "github.com/jbaxe2/blackboard.rest/_scaffolding/factory"
  "github.com/jbaxe2/blackboard.rest/oauth2"
  "github.com/jbaxe2/blackboard.rest/system"
)

/**
 * The [System] interface...
 */
type System interface {
  GetPolicies() (system.PrivacyPolicy, error)

  GetVersion() (system.VersionInfo, error)
}

/**
 * The [_BbRestSystem] type...
 */
type _BbRestSystem struct {
  _BlackboardRest

  System
}

/**
 * The [GetSystemInstance] function...
 */
func GetSystemInstance (
  host string, accessToken oauth2.AccessToken,
) System {
  hostUri, _ := url.Parse (host)

  systemService := new (_BbRestSystem)

  systemService.host = *hostUri
  systemService.accessToken = accessToken

  systemService.service.SetHost (host)
  systemService.service.SetAccessToken (accessToken)

  return systemService
}

/**
 * The [GetVersion] method...
 */
func (restSystem *_BbRestSystem) GetVersion() (system.VersionInfo, error) {
  var version system.VersionInfo

  result, err := restSystem.service.Connector.SendBbRequest (
    config.SystemEndpoints["version"], "GET", make (map[string]interface{}), 1,
  )

  if nil != err {
    return version, err
  }

  version = factory.NewVersionInfo (result.(map[string]interface{}))

  return version, err
}
