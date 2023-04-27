package routes

import (
	"restful-api/api/v1/middlewares"
	"restful-api/internal/home"
	"restful-api/internal/login"

	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.RouterGroup) {
	router.GET("/", home.Home)
	router.POST("/login", login.Login)
	router.POST("/authenticated-test", middlewares.ValidateJWT, login.Login)
}
