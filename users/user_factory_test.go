package users_test

import (
  "testing"

  "github.com/jbaxe2/blackboard.rest/users"
)

/**
 * The [TestCreateNewUser] function...
 */
func TestCreateNewUser (t *testing.T) {
  println ("Create a new user instance.")

  user := users.NewUser (rawUser)

  if !(user.Id == rawUser["id"] &&
       user.Name.Given == rawUser["name"].(map[string]interface{})["given"]) {
    t.Error ("Creating a new user instance should have expected value.")
  }
}

/**
 * The [TestCreateNewUsers] function...
 */
func TestCreateNewUsers (t *testing.T) {
  println ("Create multiple new user instances.")

  newUsers := users.NewUsers (rawUsers)

  if !(2 == len (newUsers) && newUsers[0].Id == rawUser["id"] &&
       newUsers[1].Id == rawUser2["id"]) {
    t.Error ("Creating multiple user instances should have expected values.")
  }
}

/**
 * Mocked instances to run the above tests with.
 */
var rawUsers = []map[string]interface{} {rawUser, rawUser2}

var rawUser = map[string]interface{} {
  "id": "userId",
  "uuid": "universally_unique_id",
  "externalId": "externalUserId",
  "dataSourceId": "data.source.id",
  "userName": "username",
  "studentId": "studentUserId",
  "created": "2021-11-16T18:58:19.500Z",
  "modified": "2021-11-16T18:58:19.500Z",
  "lastLogin": "2021-11-16T18:58:19.500Z",
  "institutionRoleIds": []interface{} {
    "Student",
  },
  "name": map[string]interface{} {
    "given": "first",
    "family": "last",
    "middle": "",
    "other": "",
    "suffix": "",
    "title": "",
  },
  "contact": map[string]interface{} {
    "email": "user@school.edu",
  },
  "systemRoleIds": []interface{} {
    "NONE",
  },
  "availability": map[string]interface{} {
    "available": "Yes",
  },
}

var rawUser2 = map[string]interface{} {
  "id": "userId2",
  "uuid": "universally_unique_id_2",
  "externalId": "externalUserId2",
  "dataSourceId": "data.source.id",
  "userName": "username2",
  "studentId": "studentUserId2",
  "created": "2021-11-16T18:58:19.500Z",
  "modified": "2021-11-16T18:58:19.500Z",
  "lastLogin": "2021-11-16T18:58:19.500Z",
  "institutionRoleIds": []interface{} {
    "Student",
  },
  "name": map[string]interface{} {
    "given": "first 2",
    "family": "last",
    "middle": "",
    "other": "",
    "suffix": "",
    "title": "",
  },
  "contact": map[string]interface{} {
    "email": "user2@school.edu",
  },
  "systemRoleIds": []interface{} {
    "NONE",
  },
  "availability": map[string]interface{} {
    "available": "Yes",
  },
}
