package contents

import (
  "net/url"
  "strconv"
)

/**
 * The [ContentHandler] interface provides the base type for all content handler
 * types.  Every handler has an identifier that specifies which content type to
 * handle.
 */
type ContentHandler interface {
  AsMap() map[string]interface{}
}

type DefaultHandler struct {
  ContentHandler
}

type Document struct {
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

func (handler *DefaultHandler) AsMap() map[string]interface{} {
  return map[string]interface{} {
    "id": "resource/x-bb-unknown",
  }
}

func (document *Document) AsMap() map[string]interface{} {
  return map[string]interface{} {
    "id": "resource/x-bb-document",
  }
}

func (externalLink *ExternalLink) AsMap() map[string]interface{} {
  return map[string]interface{} {
    "id": "resource/x-bb-externallink",
    "url": externalLink.Uri.String(),
  }
}

func (folder *Folder) AsMap() map[string]interface{} {
  return map[string]interface{} {
    "id": "resource/x-bb-folder",
    "isBbPage": strconv.FormatBool (folder.IsBbPage),
  }
}

func (courseLink *CourseLink) AsMap() map[string]interface{} {
  return map[string]interface{} {
    "id": "resource/x-bb-courselink",
    "targetId": courseLink.TargetId,
    "targetType": courseLink.TargetType,
  }
}

func (toolLink *ToolLink) AsMap() map[string]interface{} {
  return map[string]interface{} {
    "id": "resource/x-bb-toollink",
  }
}

func (ltiLink *LtiLink) AsMap() map[string]interface{} {
  return map[string]interface{} {
    "id": "resource/x-bb-blti-link",
    "url": ltiLink.Uri,
    "customParameters": ltiLink.CustomParameters,
  }
}
