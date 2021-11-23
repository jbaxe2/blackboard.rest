package contents

import (
  "net/url"
)

/**
 * The [ContentHandler] interface provides the base type for all content handler
 * types.  Every handler has an identifier that specifies which content type to
 * handle.
 */
type ContentHandler interface {
  Id() string
}

type DefaultHandler struct {
  ContentHandler
}

type Document struct {
  Title string

  ContentHandler
}

type ExternalLink struct {
  Uri *url.URL

  ContentHandler
}

type Folder struct {
  IsBbPage bool

  ContentHandler
}

type CourseLink struct {
  TargetId, TargetType string

  ContentHandler
}

type ToolLink struct {
  ContentHandler
}

type LtiLink struct {
  Uri *url.URL

  CustomParameters map[string]string

  ContentHandler
}

func (handler *DefaultHandler) Id() string { return "resource/x-bb-unknown" }
func (document *Document) Id() string { return "resource/x-bb-document" }
func (externalLink *ExternalLink) Id() string { return "resource/x-bb-externallink" }
func (folder *Folder) Id() string { return "resource/x-bb-folder" }
func (courseLink *CourseLink) Id() string { return "resource/x-bb-courselink" }
func (toolLink *ToolLink) Id() string { return "resource/x-bb-toollink" }
func (ltiLink *LtiLink) Id() string { return "resource/x-bb-blti-link" }
