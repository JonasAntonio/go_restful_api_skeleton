package handlers

import (
	"restful-api/api/v1/middlewares"
	"restful-api/internal/home"
	"restful-api/internal/login"
)

func (server *Server) Routes() {
	server.Router.GET("/v1", home.Home)
	server.Router.POST("/v1/login", login.Login)
	server.Router.POST("/v1/authenticated-test", middlewares.ValidateJWT, login.Login)
}
