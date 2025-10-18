package middleware

import (
	"net/http"
	"pismo/consts"
	"pismo/utils"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	header = "authorization"
	scheme = "Bearer"
)

func (m *middleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := m.logger.WithContext(c)
		claims, err := m.authenticate(c)
		if err != nil {
			log.Errorf(err.Error())
			c.JSON(http.StatusUnauthorized, utils.Send(nil, err, utils.ErrCodeUnauthorized))
			c.Abort()
		}
		c.Set(consts.ContextAuthJWTClaims, claims)
		c.Next()
	}
}

func (m *middleware) authenticate(c *gin.Context) (claims jwt.MapClaims, err error) {
	userToken, err := parseToken(c.Request.Header.Get(header))
	if err == nil {
		claims, err = m.authenticateToken(userToken)
	}
	return
}

func (m *middleware) authenticateToken(token string) (jwt.MapClaims, error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) { return []byte(m.config.JWTSecretKey), nil })
	if err != nil {
		return nil, err
	}

	claims, _ := jwtToken.Claims.(jwt.MapClaims)
	return claims, err
}

func parseToken(authHeader string) (string, error) {
	if authHeader == "" {
		return "", utils.ErrEmptyToken
	}
	splits := strings.SplitN(authHeader, " ", 2)
	if len(splits) < 2 {
		return "", utils.ErrInvalidToken
	}
	if !strings.EqualFold(splits[0], scheme) {
		return "", utils.ErrInvalidScheme
	}

	return splits[1], nil
}
