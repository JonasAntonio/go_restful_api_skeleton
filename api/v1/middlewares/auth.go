package middlewares

import (
	"fmt"
	"net/http"
	"restful-api/api/v1/helpers"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var secret = []byte("signatureKey")

func ValidateJWT(c *gin.Context) {
	tokenString := getAuthorization(c)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error in parsing")
		}
		return secret, nil
	})

	if err != nil {
		helpers.Respond(c, http.StatusForbidden, err.Error())
		c.Abort()
		return
	}

	// claims, ok := token.Claims.(jwt.MapClaims)
	_, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		helpers.Respond(c, http.StatusForbidden, "Access denied")
		c.Abort()
		return
	}

	if !token.Valid {
		helpers.Respond(c, http.StatusForbidden, "Invalid token")
		c.Abort()
		return
	}

	// TODO validate claimns

	c.Next()
}

func getAuthorization(c *gin.Context) string {
	authorization := c.Request.Header["Authorization"][0]
	authorization = strings.Replace("Bearer ", authorization, "", -1)
	return authorization
}
