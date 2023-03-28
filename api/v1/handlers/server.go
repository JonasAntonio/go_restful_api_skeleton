package handlers

import (
	routes "restful-api/api/v1/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run() {
	router := routes.Routes()

	SetCors(router)

	router.Run("localhost:8080")
}

func SetCors(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowMethods:           []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowCredentials:       true,
		AllowBrowserExtensions: false,
		AllowHeaders:           []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowOrigins:           []string{"https://*", "http://*"},
		MaxAge:                 300,
	}))
}
