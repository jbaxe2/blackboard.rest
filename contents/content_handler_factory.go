package contents

import "net/url"

/**
 * The [NewContentHandler] function creates a new Content Handler based on the
 * identifier and raw content handler information.
 */
func NewContentHandler (rawContentHandler map[string]interface{}) ContentHandler {
  var contentHandler ContentHandler

  switch rawContentHandler["id"].(string) {
    case "resource/x-bb-folder":
      contentHandler = new (Folder)
    case "resource/x-bb-document":
      contentHandler = new (Document)
    case "resource/x-bb-externallink":
      uri, _ := url.Parse (rawContentHandler["url"].(string))
      contentHandler = &ExternalLink {Uri: uri}
    case "resource/x-bb-toollink":
      contentHandler = new (ToolLink)
    case "resource/x-bb-blti-link":
      uri, _ := url.Parse (rawContentHandler["url"].(string))
      customParams, _ := rawContentHandler["customParameters"].(map[string]string)
      contentHandler = &LtiLink {Uri: uri, CustomParameters: customParams}
    default:
      contentHandler = new (DefaultHandler)
  }

  return contentHandler
}
