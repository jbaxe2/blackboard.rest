package test

import (
  "github.com/jbaxe2/blackboard.rest.go/src"
  "github.com/jbaxe2/blackboard.rest.go/src/_scaffolding/config"
  "testing"
)

/**
 * The [UsersTester] type...
 */
type UsersTester struct {
  t *testing.T
}

/**
 * The [Run] method...
 */
func (tester *UsersTester) Run() {
  print ("Users:\n")

  _testGetUsersInstance (tester.t)
  _testGetUserByPrimaryId (tester.t)
}

/**
 * The [_getUsersInstance] function...
 */
func _getUsersInstance() blackboard_rest.Users {
  authorizer := TestersAuthorizer{}
  _ = authorizer.AuthorizeForTests()

  return blackboard_rest.GetUsersInstance (
    config.Host, authorizer.accessToken,
  )
}

/**
 * The [_testGetUsersInstance] function...
 */
func _testGetUsersInstance (t *testing.T) {
  print ("Obtain a valid Users service instance.\n")

  authorizer := TestersAuthorizer{}
  _ = authorizer.AuthorizeForTests()

  usersService := blackboard_rest.GetUsersInstance (
    config.Host, authorizer.accessToken,
  )

  if nil == usersService {
    t.Error ("Obtaining a valid Users service instance failed.\n")
  }
}

/**
 * The [_testGetUserByPrimaryId] function...
 */
func _testGetUserByPrimaryId (t *testing.T) {
  print ("Get a user by his or her primary ID.\n")

  usersService := _getUsersInstance()
  user, err := usersService.GetUser ("_27_1")

  if nil != err {
    t.Error ("Error while retrieving the user: " + err.Error())

    return
  }

  if "_27_1" != user.Id {
    t.Error ("The retrieved user does not match the one selected.")
  }
}
