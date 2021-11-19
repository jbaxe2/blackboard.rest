package blackboard_rest_test

import (
  "io/ioutil"
  "net/http"
  "strings"
  "testing"

  blackboardRest "github.com/jbaxe2/blackboard.rest"
  "github.com/jbaxe2/blackboard.rest/api"
)

/**
 * The [TestCreateNewContent] function...
 */
func TestCreateNewContent (t *testing.T) {
  println ("Create a new content service instance.")

  if nil == blackboardRest.NewContent (mockContentService) {
    t.Error ("Creating new content service instance should not be nil reference.")
  }
}

/**
 * The [TestNewContentRequiresService] function...
 */
func TestNewContentRequiresService (t *testing.T) {
  println ("New content instance requires service reference.")

  if nil != blackboardRest.NewContent (nil) {
    t.Error ("Missing service should result in a nil reference.")
  }
}

/**
 * The [TestNewContentGetContents] function...
 */
func TestNewContentGetContents (t *testing.T) {
  println ("Retrieve multiple contents for a particular course from the REST API.")

  content := blackboardRest.NewContent (mockContentService)
  courseContents, err := content.GetContents ("_1_1")

  if !(nil == err && 2 == len (courseContents) &&
       "contentId1" == courseContents[0].Id && "contentId2" == courseContents[1].Id) {
    t.Error ("Retrieving course contents should return the appropriate response.")
  }
}

/**
 * The [TestNewContentGetContent] function...
 */
func TestNewContentGetContent (t *testing.T) {
  println ("Retrieve particular content from particular course from the REST API.")

  const contentId = "contentId1"

  content := blackboardRest.NewContent (mockContentService)
  newContent, err := content.GetContent ("_1_1", contentId)

  if !(nil == err && contentId == newContent.Id) {
    t.Error ("Retrieving course content should return the appropriate response.")
  }
}

/**
 * The [TestNewContentGetChildren] function...
 */
func TestNewContentGetChildren (t *testing.T) {
  println ("Retrieve children content from particular content from the REST API.")

  const parentId = "contentId1"

  content := blackboardRest.NewContent (mockContentService)
  children, err := content.GetContentChildren ("_1_1", parentId)

  if !(nil == err && 2 == len (children) &&
       parentId == children[0].ParentId && parentId == children[1].ParentId) {
    t.Error ("Retrieving children content should return the appropriate response.")
  }
}

/**
 * Mocked types and instances to run the above tests with.
 */
var mockContentService =
  api.NewService ("localhost", mockToken, new (_MockContentRoundTripper))

type _MockContentRoundTripper struct {
  http.RoundTripper
}

func (roundTripper *_MockContentRoundTripper) RoundTrip (
  request *http.Request,
) (*http.Response, error) {
  request.Response = &http.Response {
    Request: request,
    Header: make (http.Header),
  }

  request.Response.StatusCode = 200
  responseBody := ""

  switch true {
    case "GET" == request.Method && strings.Contains (request.URL.Path, "/children"):
      responseBody = childrenContents
    case "GET" == request.Method && strings.Contains (request.URL.Path, "/contents/"):
      responseBody = content1
    case "GET" == request.Method && strings.Contains (request.URL.Path, "/contents"):
      responseBody = rawContents
    default:
      request.Response.StatusCode = 404
      responseBody = improperRequest
  }

  request.Response.Body = ioutil.NopCloser (strings.NewReader (responseBody))

  return request.Response, nil
}

const rawContents = `{"results":[` + content1 + `,` + content2 + `]}`
const childrenContents = `{"results":[` + childContent1 + `,` + childContent2 + `]}`

const content1 = `{"id":"contentId1","parentId":"parentContentId1","title":` +
  `"Content Title 1","body":"Content Body 1","description":"","created":` +
  `"2021-11-19T16:54:41.273Z","modified":"2021-11-19T16:54:41.273Z","position":0,` +
  `"hasChildren":true,"hasGradebookColumns":true,"hasAssociatedGroups":true,` +
  `"launchInNewWindow":true,"reviewable":true,"availability":{"available":"Yes",` +
  `"allowGuests":true,"adaptiveRelease":{"start":"2021-11-19T16:54:41.273Z",` +
  `"end":"2021-11-19T16:54:41.273Z"}},"contentHandler":{"id":"handlerId"},` +
  `"links":[{"href": "string","rel": "string","title": "string","type": "string"}]}`

const content2 = `{"id":"contentId2","parentId":"parentContentId2","title":` +
  `"Content Title 2","body":"Content Body 2","description":"","created":` +
  `"2021-11-19T16:54:41.273Z","modified":"2021-11-19T16:54:41.273Z","position":0,` +
  `"hasChildren":false,"hasGradebookColumns":true,"hasAssociatedGroups":true,` +
  `"launchInNewWindow":true,"reviewable":true,"availability":{"available":"Yes",` +
  `"allowGuests":true,"adaptiveRelease":{"start":"2021-11-19T16:54:41.273Z",` +
  `"end":"2021-11-19T16:54:41.273Z"}},"contentHandler":{"id":"handlerId"},` +
  `"links":[{"href": "string","rel": "string","title": "string","type": "string"}]}`

const childContent1 = `{"id":"childContentId1","parentId":"contentId1","title":` +
  `"Child Content Title 1","body":"Child Content Body 1","description":"","created":` +
  `"2021-11-19T16:54:41.273Z","modified":"2021-11-19T16:54:41.273Z","position":0,` +
  `"hasChildren":true,"hasGradebookColumns":true,"hasAssociatedGroups":true,` +
  `"launchInNewWindow":true,"reviewable":true,"availability":{"available":"Yes",` +
  `"allowGuests":true,"adaptiveRelease":{"start":"2021-11-19T16:54:41.273Z",` +
  `"end":"2021-11-19T16:54:41.273Z"}},"contentHandler":{"id":"handlerId"},` +
  `"links":[{"href": "string","rel": "string","title": "string","type": "string"}]}`

const childContent2 = `{"id":"childContentId2","parentId":"contentId1","title":` +
  `"Child Content Title 2","body":"Child Content Body 2","description":"","created":` +
  `"2021-11-19T16:54:41.273Z","modified":"2021-11-19T16:54:41.273Z","position":0,` +
  `"hasChildren":true,"hasGradebookColumns":true,"hasAssociatedGroups":true,` +
  `"launchInNewWindow":true,"reviewable":true,"availability":{"available":"Yes",` +
  `"allowGuests":true,"adaptiveRelease":{"start":"2021-11-19T16:54:41.273Z",` +
  `"end":"2021-11-19T16:54:41.273Z"}},"contentHandler":{"id":"handlerId"},` +
  `"links":[{"href": "string","rel": "string","title": "string","type": "string"}]}`
