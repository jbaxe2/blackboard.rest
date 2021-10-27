package oauth2

/**
 * The [TokenInfo] interface provides the base type for token info types.
 */
type TokenInfo interface {
  ApplicationId() string

  Scope() string
}

/**
 * The [_TokenInfo] type implements the Token Info interface.
 */
type _TokenInfo struct {
  applicationId string

  scope string

  TokenInfo
}

/**
 * The [NewTokenInfo] function creates a new Token Info instance.
 */
func NewTokenInfo (applicationId, scope string) TokenInfo {
  if "" == applicationId || "" == scope {
    return nil
  }

  return &_TokenInfo {
    applicationId: applicationId,
    scope: scope,
  }
}

func (tokenInfo *_TokenInfo) ApplicationId() string {
  return tokenInfo.applicationId
}

func (tokenInfo *_TokenInfo) Scope() string {
  return tokenInfo.scope
}
