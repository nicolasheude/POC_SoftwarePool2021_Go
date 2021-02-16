package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckPalindrome() gin.HandlerFunc {
	return func(c *gin.Context) {
		var temp []byte
		temp, err := ioutil.ReadAll(c.Request.Body)
		var str []string
		err = json.Unmarshal(temp, &str)
		if err != nil {
			c.String(http.StatusBadRequest, "Bad read\n")
			// c.AbortWithError(http.StatusBadRequest, fmt.Errorf("Bad read"))
			return
		}
		fmt.Println(str)
		if str == nil {
			c.String(http.StatusBadRequest, "Bad body")
			return
		}
		fmt.Println(str)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(temp))
		c.Next()
	}
}
