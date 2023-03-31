package login

import (
	"net/http"
	"restful-api/api/v1/auth"

	"github.com/gin-gonic/gin"
)

type Credentials struct {
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

func Login(c *gin.Context) {
	credentials := Credentials{}

	err := c.BindJSON(&credentials)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := auth.MakeJWT(credentials.ClientId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"accessToken": token})
}
