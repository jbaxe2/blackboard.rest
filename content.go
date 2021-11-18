package blackboard_rest

import (
  "github.com/jbaxe2/blackboard.rest/api"
  "github.com/jbaxe2/blackboard.rest/contents"
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

  CreateChild (courseId, contentId string, content contents.Content) error
}

/**
 * The [_Content] type implements the Content interface.
 */
type _Content struct {
  Content
}

/**
 * The [NewContent] function creates a new content service instance.
 */
func NewContent (service api.Service) Content {
  if nil == service {
    return nil
  }

  return new (_Content)
}
