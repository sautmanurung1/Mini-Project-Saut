package middleware

import (
	"github.com/golang-jwt/jwt"
	"time"
)

func CreateToken(id int, username, secret string) (string, error) {
	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add(1 * time.Hour).Unix()
	claims["id"] = id
	claims["username"] = username

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
