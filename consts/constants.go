package consts

import (
	jwt "github.com/dgrijalva/jwt-go"
)

const (
	CorrelationID        = "correlation-id"
	ContextAuthJWTClaims = "auth_claims"
)

var (
	SigningMethodHS256 = jwt.SigningMethodHS256
)
