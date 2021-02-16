package router

import (
	"github.com/gin-gonic/gin"
)

func world(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "world",
	})
}

func ApplyRoutes(r *gin.Engine) {
	r.GET("/hello", world)
}
