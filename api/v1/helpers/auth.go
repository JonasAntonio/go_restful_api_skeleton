package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

const EXPIRATION_TIME = 1

var secret = []byte("signatureKey")

type Claims struct {
	jwt.RegisteredClaims
}

func MakeJWT(sub string) (string, error) {
	now := time.Now()
	baseURL := os.Getenv("URL")
	claims := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid.NewString(),
			Subject:   sub,
			Issuer:    baseURL,
			Audience:  jwt.ClaimStrings{baseURL},
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(EXPIRATION_TIME * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	tokenString, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
