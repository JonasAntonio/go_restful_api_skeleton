package handlers

import (
	"net/http"
	"restful-api/api/v1/helpers"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	helpers.Respond(c, http.StatusOK, "Home sweet home")
}
