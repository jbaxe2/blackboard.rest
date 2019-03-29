package config

const Base = "/learn/api/public/v1"

const BaseV2 = "/learn/api/public/v2"

/**
 * The [OAuth2Endpoints] function...
 */
func OAuth2Endpoints() map[string]string {
  endpoints := make(map[string]string)

  endpoints["authorization_code"] = "oauth2/authorizationcode"
  endpoints["request_token"] = "oauth2/token"

  return endpoints
}

/**
 * The [UserEndpoints] function...
 */
func UserEndpoints() map[string]string {
  endpoints := make(map[string]string)

  endpoints["users"] = "users"
  endpoints["user"] = "user/{userId}"

  return endpoints
}
