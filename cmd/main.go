package main

import (
	"log"
	"os"
	"restful-api/api/v1/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type App struct {
	basePath string
	Router   *gin.Engine
}

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := newApp()
	app.Router.Use(gin.Recovery())

	routes.SetRoutes(app.Router.Group("/v1"))

	SetCors(app.Router)

	app.Router.Run(os.Getenv("URL"))
}

func newApp() App {
	app := App{}
	app.Router = gin.Default()
	app.basePath = app.Router.BasePath()

	return app
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
