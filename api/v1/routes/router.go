package routes

import (
	"restful-api/api/v1/handlers"
	"restful-api/api/v1/middlewares"

	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.RouterGroup) {
	router.GET("/", handlers.Home)
	router.POST("/login", handlers.Login)
	router.POST("/authenticated-test", middlewares.ValidateJWT, handlers.Login)
}
