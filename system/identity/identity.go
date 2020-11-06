package identity

import "github.com/casbin/casbin/v2"

// TokenType represent TokenType
type TokenType string

const (
	// TokenJWT jwt Token
	TokenJWT TokenType = "JWT"
	// TokenSimple simple Token
	TokenSimple TokenType = "Simple"
)

// TheIdentityType represent TheIdentityType
type TheIdentityType string

const (
	// DefaultIdentity represent Default Identity
	DefaultIdentity TheIdentityType = "default"
	// SystemIdentity represent System Identity
	SystemIdentity TheIdentityType = "system"
)

// Identity represent Identity
type Identity struct {
	identityType TheIdentityType

	IsLogin     bool
	IsAnonymous bool
	TokenType   TokenType
	Token       string
	Claims      *JWTCustomClaims

	enforcer casbin.IEnforcer
}
