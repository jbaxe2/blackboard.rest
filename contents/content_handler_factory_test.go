package contents_test

import (
  "testing"

  "github.com/jbaxe2/blackboard.rest/contents"
)

/**
 * The [TestNewContentHandlerCreateFolder] function...
 */
func TestNewContentHandlerCreateFolder (t *testing.T) {
  println ("Create a new folder content handler.")

  folder := contents.NewContentHandler (
    map[string]interface{} {"id": "resource/x-bb-folder"},
  )

  if _, isFolder := folder.(*contents.Folder); !isFolder {
    t.Error ("The content handler was not created in the expected manner.")
  }
}

/**
 * The [TestNewContentHandlerCreateDocument] function...
 */
func TestNewContentHandlerCreateDocument (t *testing.T) {
  println ("Create a new document content handler.")

  document := contents.NewContentHandler (
    map[string]interface{} {"id": "resource/x-bb-document"},
  )

  if _, isDocument := document.(*contents.Document); !isDocument {
    t.Error ("The content handler was not created in the expected manner.")
  }
}

/**
 * The [TestNewContentHandlerCreateExternalLink] function...
 */
func TestNewContentHandlerCreateExternalLink (t *testing.T) {
  println ("Create a new external link content handler.")

  externalLink := contents.NewContentHandler (
    map[string]interface{} {"id": "resource/x-bb-externallink", "url": "localhost"},
  )

  if _, isExternalLink := externalLink.(*contents.ExternalLink); !isExternalLink {
    t.Error ("The content handler was not created in the expected manner.")
  }
}

/**
 * The [TestNewContentHandlerCreateToolLink] function...
 */
func TestNewContentHandlerCreateToolLink (t *testing.T) {
  println ("Create a new tool link content handler.")

  toolLink := contents.NewContentHandler (
    map[string]interface{} {"id": "resource/x-bb-toollink"},
  )

  if _, isToolLink := toolLink.(*contents.ToolLink); !isToolLink {
    t.Error ("The content handler was not created in the expected manner.")
  }
}

/**
 * The [TestNewContentHandlerCreateLtiLink] function...
 */
func TestNewContentHandlerCreateLtiLink (t *testing.T) {
  println ("Create a new LTI link content handler.")

  ltiLink := contents.NewContentHandler (
    map[string]interface{} {
      "id": "resource/x-bb-blti-link",
      "url": "localhost/launch/lti",
      "customParameters": map[string]string {},
    },
  )

  if _, isLtiLink := ltiLink.(*contents.LtiLink); !isLtiLink {
    t.Error ("The content handler was not created in the expected manner.")
  }
}

/**
 * The [TestNewContentHandlerCreatesDefaultWhenUnknownType] function...
 */
func TestNewContentHandlerCreatesDefaultWhenUnknownType (t *testing.T) {
  println ("Create a new default content handler when the type is unknown.")

  defaultHandler := contents.NewContentHandler (
    map[string]interface{} {"id": "resource/x-bb-something"},
  )

  if _, isDefault := defaultHandler.(*contents.DefaultHandler); !isDefault {
    t.Error ("The content handler was not created in the expected manner.")
  }
}
