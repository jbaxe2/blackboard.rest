package blackboard_rest

import (
  "github.com/jbaxe2/blackboard.rest.go/src/_scaffolding/services"
  "github.com/jbaxe2/blackboard.rest.go/src/oauth2"
  "github.com/jbaxe2/blackboard.rest.go/src/system"
  "net/url"
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
  host url.URL

  accessToken oauth2.AccessToken

  service services.BlackboardRestService

  System
}

/**
 * The [GetSystemInstance] function...
 */
func GetSystemInstance (
  host string, accessToken oauth2.AccessToken,
) System {
  hostUri, _ := url.Parse (host)

  systemService := &_BbRestSystem {host: *hostUri, accessToken: accessToken}
  systemService.service.SetAccessToken (accessToken)

  return systemService
}

/**
 * The [GetVersion] method...
 */
func (restSystem *_BbRestSystem) GetVersion() (system.VersionInfo, error) {
  var version system.VersionInfo
  var err error
  var result interface{}

  return version, err
}
