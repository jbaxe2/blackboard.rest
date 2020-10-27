package blackboard_rest

import (
  "net/url"
  "strings"

  "github.com/jbaxe2/blackboard.rest.go/src/_scaffolding/config"
  "github.com/jbaxe2/blackboard.rest.go/src/_scaffolding/factory"
  "github.com/jbaxe2/blackboard.rest.go/src/oauth2"
  "github.com/jbaxe2/blackboard.rest.go/src/users"
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
  _BlackboardRest

  Users
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

  usersService := new (_BbRestUsers)

  usersService.host = *hostUri
  usersService.accessToken = accessToken

  usersService.service.SetHost (host)
  usersService.service.SetAccessToken (accessToken)

  return usersService
}

/**
 * The [GetUser] method...
 */
func (restUsers *_BbRestUsers) GetUser (userId string) (users.User, error) {
  var user users.User

  endpoint := config.UsersEndpoints["user"]
  endpoint = strings.Replace (endpoint, "{userId}", userId, -1)

  result, err := restUsers.service.Connector.SendBbRequest (
    endpoint, "GET", make (map[string]interface{}), 1,
  )

  if nil != err {
    return user, err
  }

  user = factory.NewUser (result.(map[string]interface{}))

  return user, err
}
