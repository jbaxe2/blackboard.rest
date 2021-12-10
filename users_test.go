package blackboard_rest_test

import (
  "net/http"
  "strings"
  "testing"

  blackboardRest "github.com/jbaxe2/blackboard.rest"
  "github.com/jbaxe2/blackboard.rest/api"
)

/**
 * The [TestCreateNewUsers] function...
 */
func TestCreateNewUsers (t *testing.T) {
  println ("Create a new users service instance.")

  if nil == blackboardRest.NewUsers (mockUsersService) {
    t.Error ("Creating a new users service instance should not be nil reference.")
  }
}

/**
 * The [TestNewUsersRequiresService] function...
 */
func TestNewUsersRequiresService (t *testing.T) {
  println ("Creating a new users instance requires a service instance.")

  if nil != blackboardRest.NewUsers (nil) {
    t.Error ("Missing service instance should result in nil reference.")
  }
}

/**
 * The [TestNewUsersGetUser] function...
 */
func TestNewUsersGetUser (t *testing.T) {
  println ("Retrieve a user from the REST API, based on the user's external ID.")

  users := blackboardRest.NewUsers (mockUsersService)
  externalId := "externalUserId"
  user, err := users.GetUser ("externalId:" + externalId)

  if !(nil == err && user.ExternalId == externalId) {
    t.Error ("Retrieving a user should return the appropriate response.")
  }
}

/**
 * Mocked types and instances to run the above tests with.
 */
var mockUsersService =
  api.NewService ("localhost", mockToken, new (_MockUsersRoundTripper))

type _MockUsersRoundTripper struct {
  http.RoundTripper
}

func (roundTripper *_MockUsersRoundTripper) RoundTrip (
  request *http.Request,
) (*http.Response, error) {
  conditions := []bool {
    "GET" == request.Method && strings.Contains (request.URL.Path, "/users/"),
  }

  responseBodies := []string {rawUser}

  builder := NewResponseBuilder (conditions, responseBodies)

  return builder.Build (request), nil
}

const rawUser = `{"id":"userId","uuid":"universally_unique_id","externalId":` +
  `"externalUserId","dataSourceId":"data.source.id","userName":"username",` +
  `"studentId":"studentUserId","created":"2021-11-16T18:58:19.500Z",` +
  `"modified":"2021-11-16T18:58:19.500Z","lastLogin":"2021-11-16T18:58:19.500Z",` +
  `"institutionRoleIds":["Student"],"name":{"given":"first","family": "last",` +
  `"middle":"","other":"","suffix":"","title":""},"contact":{"email":` +
  `"user@school.edu"},"systemRoleIds":["NONE"],"availability":{"available":"Yes"}}`
