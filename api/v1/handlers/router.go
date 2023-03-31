package handlers

import (
	"restful-api/internal/home"
	"restful-api/internal/login"
)

func (server *Server) Routes() {
	server.Router.GET("/v1", home.Home)
	server.Router.POST("/v1/login", login.Login)
}
