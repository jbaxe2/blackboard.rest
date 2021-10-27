package errors

import "github.com/jbaxe2/blackboard.rest/utils"

/**
 * The [OAuth2Error] interface provides the base type for OAuth2 error types.
 */
type OAuth2Error interface {
  RestError
}

/**
 * The [_OAuth2Error] type implements the OAuth2 Error interface.
 */
type _OAuth2Error struct {
  code string

  description string

  OAuth2Error
}

var _OAuth2ErrorCodes = []string {
  "invalid_request", "invalid_client", "invalid_grant", "unauthorized_client",
  "unsupported_grant_type", "invalid_scope", "unsupported_response_type",
  "server_error",
}

/**
 * The [NewOAuth2Error] function creates a new OAuth2 Error instance.
 */
func NewOAuth2Error (code, description string) OAuth2Error {
  if !utils.StringInStrings (code, _OAuth2ErrorCodes) {
    return nil
  }

  return &_OAuth2Error {
    code: code,
    description: description,
  }
}

func (oauth2Error *_OAuth2Error) Code() string {
  return oauth2Error.code
}

func (oauth2Error *_OAuth2Error) Description() string {
  return oauth2Error.description
}

func (oauth2Error *_OAuth2Error) Error() string {
  return "(" + oauth2Error.code + ") " + oauth2Error.description
}
