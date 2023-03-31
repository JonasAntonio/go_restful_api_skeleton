package router

import (
	home "restful-api/internal/home"
	login "restful-api/internal/login"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine = nil

func GetInstance() *gin.Engine {
	if router == nil {
		router = gin.Default()
	}

	return router
}

func Routes() *gin.Engine {
	router := GetInstance()

	router.GET("/v1", home.Home)
	router.POST("/v1/login", login.Login)

	return router
}
