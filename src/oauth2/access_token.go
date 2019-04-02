package oauth2

/**
 * The [AccessToken] type...
 */
type AccessToken struct {
  access_token, token_type, refresh_token, scope, user_id string

  expires_in float64
}

func (accessToken *AccessToken) AccessToken() string {
  return accessToken.access_token
}

func (accessToken *AccessToken) TokenType() string {
  return accessToken.token_type
}

func (accessToken *AccessToken) RefreshToken() string {
  return accessToken.refresh_token
}

func (accessToken *AccessToken) Scope() string {
  return accessToken.scope
}

func (accessToken *AccessToken) UserId() string {
  return accessToken.user_id
}

func (accessToken *AccessToken) ExpiresIn() float64 {
  return accessToken.expires_in
}
