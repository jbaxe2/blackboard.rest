package config

const Base = "/learn/api/public/v1"

const BaseV2 = "/learn/api/public/v2"

func OAuth2Endpoints() map[string]string {
  endpoints := make(map[string]string)

  endpoints["authorization_code"] = "oauth2/authorizationcode"
  endpoints["request_token"] = "oauth2/token"

  return endpoints
}
