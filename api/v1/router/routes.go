package router

import (
	home "restful-api/internal/home"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	router := gin.Default()

	router.GET("/", home.Home)

	return router
}
