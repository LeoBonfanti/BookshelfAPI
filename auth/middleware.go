package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		token, err := ValidateToken(c)

		if err != nil {

			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if uid, exists := claims["user_id"]; exists {
				c.Set("user_id", uid)
			}
		}

		c.Next()
	}
}
