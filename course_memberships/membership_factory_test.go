package course_memberships_test

import (
  "testing"

  courseMemberships "github.com/jbaxe2/blackboard.rest/course_memberships"
)

/**
 * The [TestCreateNewMembership] function...
 */
func TestCreateNewMembership (t *testing.T) {
  println ("Create a new membership instance.")

  membership := courseMemberships.NewMembership (rawMembership1)

  if !(membership.Id == rawMembership1["id"] && membership.User.Id == "userId1") {
    t.Error ("Creating a new membership instance should have expected value.")
  }
}

/**
 * The [TestCreateNewMemberships] function...
 */
func TestCreateNewMemberships (t *testing.T) {
  println ("Create multiple new membership instances.")

  memberships := courseMemberships.NewMemberships (
    []map[string]interface{} {rawMembership1, rawMembership2},
  )

  if !(2 == len (memberships) && memberships[0].Id == rawMembership1["id"]) {
    t.Error ("Creating new membership instances should have expected values.")
  }
}

var rawMembership1 = map[string]interface{} {
  "id": "membershipId1",
  "userId": "userId1",
  "user": map[string]interface{} {
    "id": "userId1",
    "uuid": "some_user_uuid_1",
    "externalId": "externalUserId1",
    "dataSourceId": "data.source.user.id",
    "userName": "string",
    "studentId": "string",
    "educationLevel": "K8",
    "gender": "Female",
    "pronouns": "string",
    "created": "2021-11-26T17:19:32.114Z",
    "modified": "2021-11-26T17:19:32.114Z",
    "lastLogin": "2021-11-26T17:19:32.114Z",
    "institutionRoleIds": []string {
      "institutionRoleId1",
    },
    "systemRoleIds": []string {
      "SystemAdmin",
    },
    "availability": map[string]interface{} {
      "available": "Yes",
    },
    "name": map[string]interface{} {
      "given": "first",
      "family": "last",
      "middle": "",
      "other": "",
      "suffix": "",
      "title": "",
    },
    "job": map[string]interface{} {
      "title": "string",
      "department": "string",
      "company": "string",
    },
    "contact": map[string]interface{} {
      "homePhone": "string",
      "mobilePhone": "string",
      "businessPhone": "string",
      "businessFax": "string",
      "email": "string",
      "institutionEmail": "string",
      "webPage": "string",
    },
    "address": map[string]interface{} {
      "street1": "string","street2": "string","city": "string","state": "string",
      "zipCode": "string","country": "string"}, "locale": map[string]interface{} {
      "id": "string", "calendar": "Gregorian", "firstDayOfWeek": "Sunday"},
    "avatar": map[string]interface{} {
      "viewUrl": "string",
      "source": "Default",
      "uploadId": "string",
    },
  },
  "courseId": "string",
  "childCourseId": "string",
  "dataSourceId": "string",
  "created": "2021-11-26T17:19:32.114Z",
  "modified": "2021-11-26T17:19:32.114Z",
  "availability": map[string]interface{} {
    "available": "Yes",
  },
  "courseRoleId": "Instructor",
  "bypassCourseAvailabilityUntil": "2021-11-26T17:19:32.114Z",
  "lastAccessed": "2021-11-26T17:19:32.114Z",
}

var rawMembership2 = map[string]interface{} {
  "id": "membershipId2",
  "userId": "userId1",
  "user": map[string]interface{} {
    "id": "userId1",
    "uuid": "some_user_uuid_1",
    "externalId": "externalUserId1",
    "dataSourceId": "data.source.user.id",
    "userName": "string",
    "studentId": "string",
    "educationLevel": "K8",
    "gender": "Female",
    "pronouns": "string",
    "created": "2021-11-26T17:19:32.114Z",
    "modified": "2021-11-26T17:19:32.114Z",
    "lastLogin": "2021-11-26T17:19:32.114Z",
    "institutionRoleIds": []string {
      "institutionRoleId1",
    },
    "systemRoleIds": []string {
      "SystemAdmin",
    },
    "availability": map[string]interface{} {
      "available": "Yes",
    },
    "name": map[string]interface{} {
      "given": "first",
      "family": "last",
      "middle": "",
      "other": "",
      "suffix": "",
      "title": "",
    },
    "job": map[string]interface{} {
      "title": "string",
      "department": "string",
      "company": "string",
    },
    "contact": map[string]interface{} {
      "homePhone": "string",
      "mobilePhone": "string",
      "businessPhone": "string",
      "businessFax": "string",
      "email": "string",
      "institutionEmail": "string",
      "webPage": "string",
    },
    "address": map[string]interface{} {
      "street1": "string","street2": "string","city": "string","state": "string",
      "zipCode": "string","country": "string"}, "locale": map[string]interface{} {
      "id": "string", "calendar": "Gregorian", "firstDayOfWeek": "Sunday"},
    "avatar": map[string]interface{} {
      "viewUrl": "string",
      "source": "Default",
      "uploadId": "string",
    },
  },
  "courseId": "string",
  "childCourseId": "string",
  "dataSourceId": "string",
  "created": "2021-11-26T17:19:32.114Z",
  "modified": "2021-11-26T17:19:32.114Z",
  "availability": map[string]interface{} {
    "available": "Yes",
  },
  "courseRoleId": "Instructor",
  "bypassCourseAvailabilityUntil": "2021-11-26T17:19:32.114Z",
  "lastAccessed": "2021-11-26T17:19:32.114Z",
}
