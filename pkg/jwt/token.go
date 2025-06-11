package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(id int64, username, email, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       id,
		"username": username,
		"email":    email,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	key := []byte(secretKey)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString, secretKey string) (int64, string, string, error) {
	key := []byte(secretKey)
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return 0, "", "", err
	}

	if !token.Valid {
		return 0, "", "", err
	}

	return int64(claims["id"].(float64)), claims["username"].(string), claims["email"].(string), nil
}
