package jwtUtils

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/waxer59/basic-go-fiber-api/config"
	"github.com/waxer59/basic-go-fiber-api/internal/helpers"
)

type JWTClaims struct {
	exp int64
	ID  uuid.UUID
	jwt.StandardClaims
}

var jwtKey = []byte(config.GetEnv("JWT_SECRET_KEY"))

const JWT_EXPIRATION_TIME = time.Hour * 24

func NewJwt(id uuid.UUID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
		exp: time.Now().Add(JWT_EXPIRATION_TIME).Unix(),
		ID:  id,
	})

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseJwt(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func GetAndParseJwt(c *fiber.Ctx) (*JWTClaims, error) {
	token, err := helpers.GetJwtToken(c)
	if err != nil {
		return nil, err
	}

	return ParseJwt(token)
}
