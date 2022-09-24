package functions

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(id string) (string, error) {
	tokenSecret := GetEnv("TOKEN_SECRET", "qwerty")
	mySigningKey := []byte(tokenSecret)

	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)

	return ss, err
}
