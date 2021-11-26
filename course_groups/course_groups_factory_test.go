package course_groups_test

import (
  "testing"

  courseGroups "github.com/jbaxe2/blackboard.rest/course_groups"
)

/**
 * The [TestCreateNewCourseGroups] function...
 */
func TestCreateNewCourseGroup (t *testing.T) {
  println ("Create a new course group instance.")

  courseGroup := courseGroups.NewCourseGroup (rawCourseGroup1)

  if courseGroup.Id != rawCourseGroup1["id"] {
    t.Error ("Creating the course group instance should have expected value.")
  }
}

/**
 * The [TestCreateMultipleNewCourseGroups] function...
 */
func TestCreateMultipleNewCourseGroups (t *testing.T) {
  println ("Create multiple new course group instances.")

  newCourseGroups := courseGroups.NewCourseGroups (
    []map[string]interface{} {rawCourseGroup1, rawCourseGroup2},
  )

  if !(2 == len (newCourseGroups) &&
       newCourseGroups[0].Id == rawCourseGroup1["id"] &&
       newCourseGroups[1].Id == rawCourseGroup2["id"]) {
    t.Error("Creating multiple group instances should have expected values.")
  }
}

/**
 * Mocked instances to run the above tests with.
 */
var rawCourseGroup1 = map[string]interface{} {
  "id": "courseGroupId1",
  "externalId": "externalCourseGroupId1",
  "groupSetId": "groupSetId1",
  "name": "Course Group 1",
  "description": "",
  "availability": map[string]interface{} {
    "available": "Yes",
  },
  "enrollment": map[string]interface{} {
    "type": "InstructorOnly",
    "limit": 0,
    "signupSheet": map[string]interface{} {
      "name": "Course Group Signup",
      "description": "",
      "showMembers": true,
    },
  },
  "uuid": "course_group_uuid_1",
  "created": "2021-11-26T18:18:03.270Z",
  "modified": "2021-11-26T18:18:03.270Z",
}

var rawCourseGroup2 = map[string]interface{} {
  "id": "courseGroupId2",
  "externalId": "externalCourseGroupId2",
  "groupSetId": "groupSetId2",
  "name": "Course Group 2",
  "description": "",
  "availability": map[string]interface{} {
    "available": "Yes",
  },
  "enrollment": map[string]interface{} {
    "type": "InstructorOnly",
    "limit": 0,
    "signupSheet": map[string]interface{} {
      "name": "Course Group Signup",
      "description": "",
      "showMembers": true,
    },
  },
  "uuid": "course_group_uuid_2",
  "created": "2021-11-26T18:18:03.270Z",
  "modified": "2021-11-26T18:18:03.270Z",
}
