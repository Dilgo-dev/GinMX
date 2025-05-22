package middleware

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)


func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		secretKey := os.Getenv("SECRET_KEY")
		if secretKey == "" {
			log.Fatal("SECRET_KEY is not set")
			c.Redirect(http.StatusFound, "/login")
		}

		token, err := c.Cookie("token")
		if err != nil {
			c.Redirect(http.StatusFound, "/login")
		}

		claims := jwt.MapClaims{}

		_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})
		if err != nil {
			c.Redirect(http.StatusFound, "/login")
		}

		c.Set("user", claims["username"])
		c.Next()
	}
}