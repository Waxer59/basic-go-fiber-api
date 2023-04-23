package jwtUtils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/waxer59/basic-go-fiber-api/config"
)

var jwtKey = []byte(config.GetEnv("JWT_SECRET_KEY"))

func NewJwt(data map[string]interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(10 * time.Minute),
	})

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return "", errors.New("could not parse claims")
	}

	for key, value := range data {
		claims[key] = value
	}

	fmt.Println(token.Claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseJwt(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return &claims, nil
}
