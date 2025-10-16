package middleware

import (
	"net/http"
	"strings"
	"xm/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func (m *middleware) Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.Request.Header.Get("Authorization")
		if len(accessToken) < 1 {
			m.logger.WithContext(c).Warnf("authorization token not found")
			c.JSON(http.StatusUnauthorized, utils.Fail(http.StatusUnauthorized,
				"auth token not found"))
			c.Abort()
			return
		}

		if strings.HasPrefix(strings.ToUpper(accessToken), "BEARER") {
			accessToken = accessToken[len("BEARER "):]
		}

		token, err := jwt.ParseWithClaims(accessToken, &JWTClaims{}, m.GetPublicKey)
		if err != nil {
			m.logger.WithContext(c).Warnf("invalid token while authorizing request", "error", err.Error())
			c.JSON(http.StatusUnauthorized, utils.Fail(http.StatusUnauthorized,
				"invalid auth token found"))
			c.Abort()
			return
		}

		claims := token.Claims.(*JWTClaims)
		if !claims.VerifyIssuer(TokenIssuer, true) {
			m.logger.WithContext(c).Warnf("invalid token issuer while authorizing request")
			c.JSON(http.StatusUnauthorized, utils.Fail(http.StatusUnauthorized,
				"invalid auth token issuer found"))
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserId)
	}
}
