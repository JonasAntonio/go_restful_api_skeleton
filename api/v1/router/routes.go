package router

import (
	home "restful-api/internal/home"

	"github.com/gin-gonic/gin"
)

func SetRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/", home.Home)
	router.GET("/a", home.Home)

	return router
}
