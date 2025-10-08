package auth

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func getJwtKey() []byte {
	secret := os.Getenv("JWT_SECRET")

	return []byte(secret)
}

func GenerateKey(userId int) (string, error) {

	jwtKey := getJwtKey()

	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(2 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateToken(c *gin.Context) (*jwt.Token, error) {
	jwtKey := getJwtKey()

	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "NULL TOKEN"})
		c.Abort()
		return nil, fmt.Errorf("NULL TOKEN")
	}

	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("INVALID SIGNATURE METHOD")
		}
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "INVALID TOKEN"})
		c.Abort()
		return nil, fmt.Errorf("INVALID TOKEN")
	}

	return token, nil
}
