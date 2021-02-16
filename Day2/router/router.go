package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func world(c *gin.Context) {
	c.String(http.StatusOK, "Hello world")
}

func query(c *gin.Context) {
	message := c.Query("message")
	if message == "" {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.String(http.StatusOK, "%s", message)
	}
}

func param(c *gin.Context) {
	message := c.Param("message")
	if message == "" {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.String(http.StatusOK, "%s", message)
	}
}

func header(c *gin.Context) {
	message := c.GetHeader("message")
	if message == "" {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.String(http.StatusOK, "%s", message)
	}
}

func cookie(c *gin.Context) {
	cookie, err := c.Cookie("message")
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.String(http.StatusOK, "%s", cookie)
	}
}

func body(c *gin.Context) {
	temp := c.PostForm("message")
	if temp == "" {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.String(http.StatusOK, "%s", temp)
	}
}

func ApplyRoutes(r *gin.Engine) {
	r.GET("/hello", world)
	r.GET("/repeat-my-query", query)
	r.GET("/repeat-my-param/:message", param)
	r.POST("/repeat-my-body", body)
	r.GET("/repeat-my-header", header)
	r.GET("/repeat-my-cookie", cookie)
}
