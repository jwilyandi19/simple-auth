package utils

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("simple_auth")

type JWTClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(time.Hour)

	claims := &JWTClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Println("Can't sign token: ", err.Error())
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(reqToken string) error {
	token, err := jwt.Parse(reqToken, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		log.Println("Error parsing: ", err.Error())
		return err
	}
	if !token.Valid {
		log.Println("Token not valid")
		return errors.New("token not valid")
	}
	return nil
}
