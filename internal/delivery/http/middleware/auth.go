package middleware

import (
	"net/http"
	"strings"

	"github.com/faisallbhr/light-pos-be/pkg/httpx"
	"github.com/faisallbhr/light-pos-be/pkg/jwtx"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtManager *jwtx.JWTManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			httpx.ResponseError(c, "unauthorized", http.StatusUnauthorized, map[string]string{
				"token": "authorization header is missing",
			})
			c.Abort()
			return
		}

		tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
		if tokenString == "" {
			httpx.ResponseError(c, "unauthorized", http.StatusUnauthorized, map[string]string{
				"token": "invalid authorization header format",
			})
			c.Abort()
			return
		}

		claims, err := jwtManager.ValidateToken(tokenString, jwtx.AccessToken)
		if err != nil {
			httpx.ResponseError(c, "unauthorized", http.StatusUnauthorized, map[string]string{
				"token": "invalid token",
			})
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
