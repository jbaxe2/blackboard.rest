package blackboard_rest

import (
  "github.com/jbaxe2/blackboard.rest/api"
  "github.com/jbaxe2/blackboard.rest/users"
)

/**
 * The [Users] interface provides the base interface for interacting with the
 * REST API's users endpoints.
 */
type Users interface {
  GetUsers() ([]users.User, error)

  CreateUser (user users.User) error

  GetUser (userId string) (users.User, error)

  UpdateUser (userId string, user users.User) error
}

/**
 * The [_Users] type implements the Users interface.
 */
type _Users struct {
  service api.Service

  Users
}

/**
 * The [NewUsers] function creates a new Users instance.
 */
func NewUsers (service api.Service) Users {
  if nil == service {
    return nil
  }

  return &_Users {
    service: service,
  }
}
