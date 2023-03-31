package auth

import (
	// "restful-api/api/v1/router"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

const EXPIRATION_TIME = 15

var secret = []byte("signatureKey")

func MakeJWT(sub string) (string, error) {
	now := time.Now()
	token := jwt.New(jwt.SigningMethodHS512)
	// router := router.GetInstance()
	claims := token.Claims.(jwt.MapClaims)

	claims["sub"] = sub
	// claims["iss"] = router.BasePath()
	// claims["aud"] = router.BasePath()
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
