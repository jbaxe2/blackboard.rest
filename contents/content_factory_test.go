package contents_test

import (
  "testing"

  "github.com/jbaxe2/blackboard.rest/contents"
)

/**
 * The [TestCreateNewContent] function...
 */
func TestCreateNewContent (t *testing.T) {
  println ("Create a new content instance.")

  content := contents.NewContent (rawContent)

  if !(content.Id == rawContent["id"] &&
       content.Title == rawContent["title"].(string)) {
    t.Error ("Creating content instance should have the expected value.")
  }
}

/**
 * The [TestCreateNewContents] function...
 */
func TestCreateNewContents (t *testing.T) {
  println ("Create multiple new content instances.")

  newContents := contents.NewContents (rawContents)

  if !(2 == len (newContents) && newContents[0].Id == rawContent["id"]) {
    t.Error ("Creating multiple content instances should have expected values.")
  }
}

/**
 * Mocked instances to run the above tests with.
 */
var rawContents = []map[string]interface{} {rawContent, rawContent2}

var rawContent = map[string]interface{} {
  "id": "contentId",
  "parentId": "parentContentId",
  "title": "Content Title",
  "body": "Content body, that may contain some (Bb|HT)ML.",
  "created": "2021-11-18T16:53:40.693Z",
  "modified": "2021-11-18T16:53:40.693Z",
  "position": 0,
  "hasChildren": false,
  "hasGradebookColumns": false,
  "hasAssociatedGroups": false,
  "launchInNewWindow": false,
  "reviewable": false,
  "availability": map[string]interface{} {
    "available": "Yes",
    "allowGuests": false,
    "adaptiveRelease": map[string]string {
      "start": "2021-11-18T16:53:40.693Z",
      "end": "2021-11-18T16:53:40.693Z",
    },
  },
  "contentHandler": map[string]string {
    "id": "resource/x-bb-document",
  },
  "links": map[string]string {
    "url": "localhost",
    "rel": "alternate",
    "title": "some:iri:reference:title",
    "type": "linkType",
  },
}

var rawContent2 = map[string]interface{} {
  "id": "contentId2",
  "parentId": "parentContentId",
  "title": "Content Title #2",
  "body": "Content body, that may contain some (Bb|HT)ML.",
  "created": "2021-11-18T16:53:40.693Z",
  "modified": "2021-11-18T16:53:40.693Z",
  "position": 0,
  "hasChildren": false,
  "hasGradebookColumns": false,
  "hasAssociatedGroups": false,
  "launchInNewWindow": false,
  "reviewable": false,
  "availability": map[string]interface{} {
    "available": "Yes",
    "allowGuests": false,
    "adaptiveRelease": map[string]string {
      "start": "2021-11-18T16:53:40.693Z",
      "end": "2021-11-18T16:53:40.693Z",
    },
  },
  "contentHandler": map[string]string {
    "id": "resource/x-bb-document",
  },
  "links": map[string]string {
    "url": "localhost",
    "rel": "alternate",
    "title": "some:iri:reference:title",
    "type": "linkType",
  },
}
