package handlers

import (
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("your-secret-key")

func CreateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{}

	claims["user_id"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
