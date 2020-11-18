package oauth2

/**
 * The [AccessToken] type...
 */
type AccessToken struct {
  accessToken, tokenType, refreshToken, scope, userId string

  expiresIn float64
}

func (accessToken *AccessToken) AccessToken() string {
  return accessToken.accessToken
}

func (accessToken *AccessToken) TokenType() string {
  return accessToken.tokenType
}

func (accessToken *AccessToken) RefreshToken() string {
  return accessToken.refreshToken
}

func (accessToken *AccessToken) Scope() string {
  return accessToken.scope
}

func (accessToken *AccessToken) UserId() string {
  return accessToken.userId
}

func (accessToken *AccessToken) ExpiresIn() float64 {
  return accessToken.expiresIn
}
