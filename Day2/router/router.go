package router

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func world(c *gin.Context) {
	temp := os.Getenv("HELLO_MESSAGE")
	if temp == "" {
		c.String(http.StatusNotFound, "no message defined")
	}
	c.String(http.StatusOK, "%s", temp)
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

func health(c *gin.Context) {
	c.AbortWithStatus(http.StatusOK)
}

func ApplyRoutes(r *gin.Engine) {
	r.GET("/hello", world)
	r.GET("/health", health)
	r.GET("/repeat-my-query", query)
	r.GET("/repeat-my-param/:message", param)
	r.POST("/repeat-my-body", body)
	r.GET("/repeat-my-header", header)
	r.GET("/repeat-my-cookie", cookie)
}
