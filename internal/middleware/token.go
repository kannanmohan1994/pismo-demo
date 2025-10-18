package middleware

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

type TokenFunc func(int) (string, error)

func NewUserToken(method jwt.SigningMethod, secret string, expiration time.Duration) TokenFunc {
	return func(userId int) (string, error) {
		tNow := time.Now()
		t := jwt.NewWithClaims(method, jwt.StandardClaims{
			ExpiresAt: tNow.Add(expiration).Unix(),
			IssuedAt:  tNow.Unix(),
			NotBefore: tNow.Unix(),
			Issuer:    "service",
			Subject:   fmt.Sprint(userId),
		})

		token, err := t.SignedString([]byte(secret))
		if err != nil {
			return "", errors.Wrap(err, "token signing failed")
		}
		return token, nil
	}
}
