package identity

import (
	"github.com/casbin/casbin/v2"
	"github.com/d3ta-go/system/system/handler"
)

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

	ctx      interface{}
	handler  *handler.Handler
	enforcer casbin.IEnforcer
}

// NewIdentity new Identity
func NewIdentity(iType TheIdentityType, tokenType TokenType, token string, claims *JWTCustomClaims, ctx interface{}, h *handler.Handler) (Identity, error) {

	i := Identity{
		identityType: iType,
		handler:      h,
		IsLogin:      false,
		IsAnonymous:  false,
		TokenType:    tokenType,
		Token:        token,
		Claims:       claims,
		ctx:          ctx,
	}

	return i, nil
}
