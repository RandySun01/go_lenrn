package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/jsonMap", func(c *gin.Context) {
		// 方法1: 使用map
		data := map[string]interface{}{
			"name":     "randysun",
			"age":      18,
			"messages": "jsonMap",
		}

		c.JSON(http.StatusOK, data)

	})

	r.GET("/jsonStruct", func(c *gin.Context) {
		// 方法2：通过结构体方式
		type userInfo struct {
			Name     string `json:"name"`
			Age      int
			Messages string
			low      string
		}
		user := userInfo{
			"randySun",
			18,
			"json struct",
			"小写访问不到",
		}
		c.JSON(http.StatusOK, user)
	})
	r.Run(":9999")
}
