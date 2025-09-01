package utils

import (
	config "AuthInGo/config/env"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func CreateJWTToken(id int64, email string) (string, error) {
	secretKey := config.GetString("JWT_SECRET", "defaultsecretkey")
	fmt.Println("Secret key:", secretKey)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": email,
			"id":    id,
		})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
