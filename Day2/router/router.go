package router

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type Queries struct {
	Key   string
	Value string
}

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

func extSepStr(line string) ([]string, error) {
	mylist := strings.Split(line, "=")
	if len(mylist) != 2 {
		return nil, errors.New("Bad format")
	}
	return mylist, nil
}

func extAllQueries(c *gin.Context) ([]Queries, error) {
	temp := c.Request.URL.Query().Encode()
	mylist := strings.Split(string(temp), "&")
	lenContent := len(mylist)
	q := make([]Queries, lenContent)
	for i := range mylist {
		myQuery, err := extSepStr(mylist[i])
		if err != nil {
			return nil, errors.New("Bad format")
		}
		q[i].Key = myQuery[0]
		q[i].Value = myQuery[1]
	}
	return q, nil
}

func queries(c *gin.Context) {
	q, err := extAllQueries(c)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(http.StatusOK, q)
	}
}

// func fn0(c *gin.Context) []Queries {
// 	q := c.Request.URL.Query()
// 	lenContent := len(q)
// 	t := make([]Queries, lenContent)
// 	var i int
// 	for key, value := range q {
// 		t[i].Key = key
// 		t[i].Value = value[0]
// 		i++
// 	}
// 	return q
// }

func ApplyRoutes(r *gin.Engine) {
	r.GET("/hello", world)
	r.GET("/health", health)
	r.GET("/repeat-my-query", query)
	r.GET("/repeat-my-param/:message", param)
	r.POST("/repeat-my-body", body)
	r.GET("/repeat-my-header", header)
	r.GET("/repeat-my-cookie", cookie)
	r.GET("/repeat-all-my-queries", queries)
}
