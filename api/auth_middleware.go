package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// AuthRequired ensures a valid JWT is present in cookie.
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr, err := c.Cookie("jwt")
		if err != nil || tokenStr == "" {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}
		secret := os.Getenv("JWT_SECRET")
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			// Ensure method is HMAC
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}
		c.Next()
	}
}
