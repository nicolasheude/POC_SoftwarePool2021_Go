package server

import (
	"SofwareGoDay4/router"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Def *gin.Engine
}

func NewServer() Server {
	server := Server{Def: gin.Default()}
	router.ApplyRoutes(server.Def)
	return server
}
