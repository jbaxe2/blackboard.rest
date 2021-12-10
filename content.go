package blackboard_rest

import (
  "strings"

  "github.com/jbaxe2/blackboard.rest/api"
  "github.com/jbaxe2/blackboard.rest/contents"
  "github.com/jbaxe2/blackboard.rest/utils"
)

/**
 * The [Content] interface provides the base type for interacting with the REST
 * API's content endpoints.
 */
type Content interface {
  GetContents (courseId string) ([]contents.Content, error)

  CreateContent (courseId string, content contents.Content) error

  CreateAssignment (
    courseId string, assignment contents.Assignment,
  ) (contents.CreateAssignmentResult, error)

  GetContent (courseId, contentId string) (contents.Content, error)

  DeleteContent (courseId, contentId string) error

  UpdateContent (
    courseId, contentId string, content contents.Content,
  ) error

  GetContentChildren (courseId, contentId string) ([]contents.Content, error)

  CreateChild (
    courseId, contentId string, content contents.Content,
  ) (contents.Content, error)
}

/**
 * The [_Content] type implements the Content interface.
 */
type _Content struct {
  service api.Service

  Content
}

/**
 * The [NewContent] function creates a new content service instance.
 */
func NewContent (service api.Service) Content {
  if nil == service {
    return nil
  }

  return &_Content {
    service: service,
  }
}

/**
 * The [GetContents] method retrieves a slice of content instances for a particular
 * course based on the provided course ID.
 */
func (content *_Content) GetContents (courseId string) ([]contents.Content, error) {
  endpoint := strings.Replace (string (api.Contents), "{courseId}", courseId, 1)
  rawContents, err := content.service.Request (endpoint, "GET", nil, 1)

  if nil != err {
    return nil, err
  }

  return contents.NewContents (
    utils.NormalizeRawResponse (rawContents["results"].([]interface{})),
  ), nil
}

/**
 * The [GetContent] method retrieves an instance of some particular content based
 * on the provided course ID and content ID.
 */
func (content *_Content) GetContent (
  courseId, contentId string,
) (contents.Content, error) {
  endpoint := strings.Replace (string (api.Content), "{courseId}", courseId, 1)
  endpoint = strings.Replace (endpoint, "{contentId}", contentId, 1)

  rawContent, err := content.service.Request (endpoint, "GET", nil, 1)

  if nil != err {
    return contents.Content{}, err
  }

  return contents.NewContent (rawContent), nil
}

/**
 * The [GetContentChildren] method retrieves the children content items for some
 * particular content item, as specified by the course and content identifiers.
 */
func (content *_Content) GetContentChildren (
  courseId, contentId string,
) ([]contents.Content, error) {
  endpoint := strings.Replace (string (api.ContentChildren), "{courseId}", courseId, 1)
  endpoint = strings.Replace (endpoint, "{contentId}", contentId, 1)

  rawChildren, err := content.service.Request (endpoint, "GET", nil, 1)

  if nil != err {
    return nil, err
  }

  return contents.NewContents (
    utils.NormalizeRawResponse (rawChildren["results"].([]interface{})),
  ), nil
}

/**
 * The [CreateChild] method creates a new child content item to a parent content
 * item.  The identifier of the parent item is passed in through the content ID
 * parameter.
 */
func (content *_Content) CreateChild (
  courseId, contentId string, childContent contents.Content,
) (contents.Content, error) {
  endpoint := strings.Replace (string (api.ContentChildren), "{courseId}", courseId, 1)
  endpoint = strings.Replace (endpoint, "{contentId}", contentId, 1)

  newContent, err :=
    content.service.Request (endpoint, "POST", childContent.AsNewContentMap(), 1)

  if nil != err {
    return contents.Content{}, err
  }

  return contents.NewContent (newContent), nil
}
