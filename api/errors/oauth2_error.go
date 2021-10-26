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

  return new (_OAuth2Error)
}
