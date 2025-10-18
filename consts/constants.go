package consts

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	CorrelationID        = "correlation-id"
	ContextAuthJWTClaims = "auth_claims"
)

var (
	SigningMethodHS256     = jwt.SigningMethodHS256
	DefaultJWTValidityTime = 24 * time.Hour
)
