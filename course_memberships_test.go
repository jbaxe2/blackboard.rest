package blackboard_rest_test

import (
  "net/http"
  "strings"
  "testing"

  blackboardRest "github.com/jbaxe2/blackboard.rest"
  "github.com/jbaxe2/blackboard.rest/api"
)

/**
 * The [TestCreateNewCourseMemberships] function...
 */
func TestCreateNewCourseMemberships (t *testing.T) {
  println ("Create a new course memberships instance.")

  if nil == blackboardRest.NewCourseMemberships (mockCourseMembershipsService) {
    t.Error ("Creating new course memberships instance should not be nil reference.")
  }
}

/**
 * The [TestNewCourseMembershipsRequiresService] function...
 */
func TestNewCourseMembershipsRequiresService (t *testing.T) {
  println ("Creating new course memberships instance requires a service instance.")

  if nil != blackboardRest.NewCourseMemberships (nil) {
    t.Error ("Missing service instance should result in nil reference.")
  }
}

/**
 * The [TestNewCourseMembershipsGetMembershipsForCourse] function...
 */
func TestNewCourseMembershipsGetMembershipsForCourse (t *testing.T) {
  println ("Get the course memberships for a course.")

  memberships := blackboardRest.NewCourseMemberships (mockCourseMembershipsService)
  courseMemberships, err := memberships.GetMembershipsForCourse ("courseId1")

  if !(nil == err && 2 == len (courseMemberships)) {
    t.Error ("Retrieving memberships for a course should have appropriate responses.")
  }
}

/**
 * The [TestNewCourseMembershipsGetMembershipsForUser] function...
 */
func TestNewCourseMembershipsGetMembershipsForUser (t *testing.T) {
  println ("Get the course memberships for a users.")

  memberships := blackboardRest.NewCourseMemberships (mockCourseMembershipsService)
  userMemberships, err := memberships.GetMembershipsForUser ("userId1")

  if !(nil == err && 2 == len (userMemberships)) {
    t.Error ("Retrieving memberships for a user should have appropriate responses.")
  }
}

/**
 * Mocked types and instances to run the above tests with.
 */
var mockCourseMembershipsService =
  api.NewService ("localhost", mockToken, new (_MockCourseMembershipsRoundTripper))

type _MockCourseMembershipsRoundTripper struct {
  http.RoundTripper
}

func (roundTripper *_MockCourseMembershipsRoundTripper) RoundTrip (
  request *http.Request,
) (*http.Response, error) {
  coursesIndex := strings.Index (request.URL.Path, "courses")
  usersIndex := strings.Index (request.URL.Path, "users")

  conditions := []bool {
    "GET" == request.Method && (coursesIndex < usersIndex),
    "GET" == request.Method && (coursesIndex > usersIndex),
  }

  responseBodies := []string {rawCourse1Memberships, rawUser1Memberships}

  builder := NewResponseBuilder (conditions, responseBodies)

  return builder.Build (request), nil
}

const rawCourse1Memberships = `{"results":[` + rawCourse1Membership1 + `,` +
  rawCourse1Membership2 + `]}`

const rawUser1Memberships = `{"results":[` + rawCourse1Membership1 + `,` +
  rawCourse2Membership1 + `]}`

const rawCourse1Membership1 = `{"id":"membershipId1","userId":"userId1","user":` +
  `{"id":"userId1","uuid":"universally_unique_id_1","externalId":` +
  `"externalUserId1","dataSourceId":"data.source.id","userName":"username",` +
  `"studentId":"studentUserId1","created":"2021-11-16T18:58:19.500Z",` +
  `"modified":"2021-11-16T18:58:19.500Z","lastLogin":"2021-11-16T18:58:19.500Z",` +
  `"institutionRoleIds":["Student"],"name":{"given":"first","family": "last",` +
  `"middle":"","other":"","suffix":"","title":""},"contact":{"email":` +
  `"user1@school.edu"},"systemRoleIds":["NONE"],"availability":{"available":"Yes"}},` +
  `"courseId":"courseId1","childCourseId":"","dataSourceId":"data.source.id",` +
  `"created":"2022-01-20T18:18:03.323Z","modified":"2022-01-20T18:18:03.323Z",` +
  `"availability":{"available":"Yes"},"courseRoleId": "Instructor",` +
  `"bypassCourseAvailabilityUntil":"2022-01-20T18:18:03.323Z",` +
  `"lastAccessed":"2022-01-20T18:18:03.323Z"}`

const rawCourse1Membership2 = `{"id":"membershipId2","userId":"userId2","user":` +
  `{"id":"userId2","uuid":"universally_unique_id_2","externalId":` +
  `"externalUserId2","dataSourceId":"data.source.id","userName":"username",` +
  `"studentId":"studentUserId2","created":"2021-11-16T18:58:19.500Z",` +
  `"modified":"2021-11-16T18:58:19.500Z","lastLogin":"2021-11-16T18:58:19.500Z",` +
  `"institutionRoleIds":["Student"],"name":{"given":"first","family":"last",` +
  `"middle":"","other":"","suffix":"","title":""},"contact":{"email":` +
  `"user2@school.edu"},"systemRoleIds":["NONE"],"availability":{"available":"Yes"}},` +
  `"courseId":"courseId1","childCourseId":"","dataSourceId":"data.source.id",` +
  `"created":"2022-01-20T18:18:03.323Z","modified":"2022-01-20T18:18:03.323Z",` +
  `"availability":{"available":"Yes"},"courseRoleId": "Instructor",` +
  `"bypassCourseAvailabilityUntil":"2022-01-20T18:18:03.323Z",` +
  `"lastAccessed":"2022-01-20T18:18:03.323Z"}`

const rawCourse2Membership1 = `{"id":"membershipId3","userId":"userId1","user":` +
  `{"id":"userId1","uuid":"universally_unique_id_1","externalId":` +
  `"externalUserId1","dataSourceId":"data.source.id","userName":"username",` +
  `"studentId":"studentUserId1","created":"2021-11-16T18:58:19.500Z",` +
  `"modified":"2021-11-16T18:58:19.500Z","lastLogin":"2021-11-16T18:58:19.500Z",` +
  `"institutionRoleIds":["Student"],"name":{"given":"first","family": "last",` +
  `"middle":"","other":"","suffix":"","title":""},"contact":{"email":` +
  `"user1@school.edu"},"systemRoleIds":["NONE"],"availability":{"available":"Yes"}},` +
  `"courseId":"courseId2","childCourseId":"","dataSourceId":"data.source.id",` +
  `"created":"2022-01-20T18:18:03.323Z","modified":"2022-01-20T18:18:03.323Z",` +
  `"availability":{"available":"Yes"},"courseRoleId": "Instructor",` +
  `"bypassCourseAvailabilityUntil":"2022-01-20T18:18:03.323Z",` +
  `"lastAccessed":"2022-01-20T18:18:03.323Z"}`
