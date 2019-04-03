package blackboard_rest

import (
  "github.com/jbaxe2/blackboard.rest.go/src/_scaffolding/config"
  error2 "github.com/jbaxe2/blackboard.rest.go/src/_scaffolding/error"
  "github.com/jbaxe2/blackboard.rest.go/src/_scaffolding/factory"
  "github.com/jbaxe2/blackboard.rest.go/src/_scaffolding/services"
  "github.com/jbaxe2/blackboard.rest.go/src/oauth2"
  "github.com/jbaxe2/blackboard.rest.go/src/users"
  "net/url"
  "strings"
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

  service services.BlackboardRestServices

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

  usersService := &_BbRestUsers {host: *hostUri, accessToken: accessToken}
  usersService.service.SetAccessToken (accessToken)

  return usersService
}

/**
 * The [GetUser] method...
 */
func (restUsers *_BbRestUsers) GetUser (userId string) (users.User, error) {
  var user users.User
  var err error
  var result interface{}

  endpoint := config.UserEndpoints()["user"]
  endpoint = strings.Replace (endpoint, "{userId}", userId, -1)

  rawUser := make (map[string]interface{})

  result, err = restUsers.service.Connector.SendBbRequest (
    endpoint, "GET", make (map[string]interface{}), 1,
  )

  if (nil != err) && (error2.RestError{} != err) {
    return user, restUsers.HandleError (err.(error2.RestError))
  }

  rawUser = result.(map[string]interface{})

  user = (new (factory.UserFactory)).NewUser (rawUser)

  return user, err
}

/**
 * The [HandleError] method...
 */
func (restUsers *_BbRestUsers) HandleError (err error2.RestError) error2.UsersError {
  usersErr := error2.UsersError{}

  usersErr.SetStatus (err.Status())
  usersErr.SetCode (err.Code())
  usersErr.SetMessage (err.Message())
  usersErr.SetDeveloperMessage (err.DeveloperMessage())
  usersErr.SetExtraInfo (err.ExtraInfo())

  return usersErr
}
