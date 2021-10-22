package oauth2

import "strings"

/**
 * The [Token] interface provides the base type for access tokens.
 */
type Token interface {
  AccessToken() string

  TokenType() string

  RefreshToken() string

  Scope() string

  UserId() string

  ExpiresIn() float64
}

/**
 * The [_Token] type implements the Token interface.
 */
type _Token struct {
  accessToken, tokenType, refreshToken, scope, userId string

  expiresIn float64

  Token
}

/**
 * The [NewToken] function creates a new Token instance.
 */
func NewToken (
  accessToken, tokenType, refreshToken, scope, userId string, expiresIn float64,
) Token {
  if !_verifyTokenConditions (
    accessToken, tokenType, refreshToken, scope, userId, expiresIn,
  ) {
    return nil
  }

  return new (_Token)
}

func (token *_Token) AccessToken() string {
  return token.accessToken
}

func (token *_Token) TokenType() string {
  return token.tokenType
}

func (token *_Token) RefreshToken() string {
  return token.refreshToken
}

func (token *_Token) Scope() string {
  return token.scope
}

func (token *_Token) UserId() string {
  return token.userId
}

func (token *_Token) ExpiresIn() float64 {
  return token.expiresIn
}

/**
 * The [_verifyTokenConditions] function verifies the conditions used to create
 * an OAuth2 token are as they should be.
 */
func _verifyTokenConditions (
  accessToken, tokenType, refreshToken, scope, userId string, expiresIn float64,
) bool {
  if "" == accessToken || "" == tokenType {
    return false
  }

  if strings.Contains (scope, "offline") && "" == refreshToken {
    return false
  }

  return true
}
