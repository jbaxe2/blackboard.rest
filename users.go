package blackboard_rest

import (
  "github.com/jbaxe2/blackboard.rest/users"
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
