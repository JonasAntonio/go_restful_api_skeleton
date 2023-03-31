package handlers

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	basePath string
	Router   *gin.Engine
}

func New() Server {
	server := Server{}
	server.Router = gin.Default()
	server.basePath = server.Router.BasePath()

	return server
}

func Run() {
	server := New()
	server.Router.Use(gin.Recovery())
	server.Routes()

	SetCors(server.Router)

	server.Router.Run(os.Getenv("URL"))
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
