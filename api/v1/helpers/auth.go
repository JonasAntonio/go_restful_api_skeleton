package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

const EXPIRATION_TIME = 15

var secret = []byte("signatureKey")

func MakeJWT(sub string) (string, error) {
	now := time.Now()
	token := jwt.New(jwt.SigningMethodHS512)
	claims := token.Claims.(jwt.MapClaims)
	baseURL := os.Getenv("URL")

	claims["sub"] = sub
	claims["iss"] = baseURL
	claims["aud"] = baseURL
	claims["jti"] = uuid.New()
	claims["iat"] = now
	claims["nbf"] = now
	claims["exp"] = now.Add(EXPIRATION_TIME * time.Minute)

	tokenString, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
