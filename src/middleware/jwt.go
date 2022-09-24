package middlware

import (
	"fmt"
	functions "go-rest-api/src/functions"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func ValidateToken() gin.HandlerFunc {
	tokenSecret := functions.GetEnv("TOKEN_SECRET", "qwerty")

	type JwtClaims struct {
		jwt.RegisteredClaims
	}

	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")

		if auth == "" {
			c.String(http.StatusForbidden, "Forbidden")
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(auth, "Bearer ")
		if tokenString == auth {
			c.String(http.StatusForbidden, "Forbidden")
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(tokenSecret), nil
		})

		if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
			userId := claims.RegisteredClaims.Issuer
			c.Request.Header.Add("request-user-id", userId)
		} else {
			fmt.Printf("Error on JWT middleware: %v", err)
		}

		c.Next()
	}
}
