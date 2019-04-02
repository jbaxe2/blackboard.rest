package blackboard_rest

import (
  "github.com/jbaxe2/blackboard.rest.go/src/_scaffolding/services"
  "github.com/jbaxe2/blackboard.rest.go/src/oauth2"
  "github.com/jbaxe2/blackboard.rest.go/src/users"
  "net/url"
)

/**
 * The [Users] interface...
 */
type Users interface {
  GetUsers() ([]users.User, error)

  CreateUser (user users.User) error

  GetUser (userId string) (users.User, error)

  UpdateUser (userId string, user users.User) error
}

/**
 * The [_BbRestUsers] type...
 */
type _BbRestUsers struct {
  host url.URL

  accessToken oauth2.AccessToken

  Users

  services.BbRestServices
}

func (restUsers *_BbRestUsers) Host() url.URL {
  return restUsers.host
}

func (restUsers *_BbRestUsers) AccessToken() oauth2.AccessToken {
  return restUsers.accessToken
}

/**
 * The [GetUsersInstance] function...
 */
func GetUsersInstance (host string, accessToken oauth2.AccessToken) Users {
  hostUri, _ := url.Parse (host)

  return &_BbRestUsers {host: *hostUri, accessToken: accessToken}
}

/**
 * The [GetUser] method...
 */
func (restUsers *_BbRestUsers) GetUser (userId string) (users.User, error) {
  var user users.User
  var err error

  return user, err
}
