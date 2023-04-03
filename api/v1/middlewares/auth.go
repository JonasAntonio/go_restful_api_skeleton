package middlewares

import (
	"net/http"
	"restful-api/api/v1/helpers"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var secret = []byte("signatureKey")

func ValidateJWT(c *gin.Context) {
	claims := &helpers.Claims{}
	token, err := jwt.ParseWithClaims(getAuthorization(c), claims, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{
				"error":   true,
				"message": "Access denied",
			})
		return
	}

	if !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{
				"error":   true,
				"message": "Invalid token",
			})
		return
	}

	validateClaims(c, claims)

	// tokenClaims, ok := token.Claims.(jwt.MapClaims)
	// if !ok {
	// 	c.AbortWithStatusJSON(http.StatusUnauthorized,
	// 		gin.H{
	// 			"error":   true,
	// 			"message": "Access denied 1",
	// 		})
	// 	return
	// }

	// TODO validate claimns

	// c.Next()
}

func getAuthorization(c *gin.Context) string {
	authorization := c.Request.Header["Authorization"][0]
	authorization = strings.Replace(authorization, "Bearer ", "", -1)
	return authorization
}

func validateClaims(c *gin.Context, claims *helpers.Claims) {
	if !claims.VerifyExpiresAt(time.Now(), false) {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{
				"error":   true,
				"message": "Token is expired",
			})
		return
	}
}
