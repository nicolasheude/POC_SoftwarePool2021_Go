package server

import (
	"SofwareGoDay2/router"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Def *gin.Engine
}

func NewServer() Server {
	s := Server{Def: gin.Default()}
	router.ApplyRoutes(s.Def)
	return s
}
