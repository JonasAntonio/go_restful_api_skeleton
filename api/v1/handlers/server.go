package handlers

import routes "restful-api/api/v1/router"

func Run() {
	router := routes.SetRoutes()

	router.Run("localhost:8080")
}
