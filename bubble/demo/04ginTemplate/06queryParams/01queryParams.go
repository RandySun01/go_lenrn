package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	// GET请求 url ?后面是querystring参数,key = value格式,多个key-value用&连接
	// eg: /queryParams?name=randySun&age=18
	r.GET("/queryParams", func(c *gin.Context) {
		// 获取浏览器那边发送请求携带的 query string 参数
		// 方式一
		//name := c.Query("name")
		//age := c.Query("age")
		// 方式二
		//name := c.DefaultQuery("name", "body")
		//age := c.DefaultQuery("age", "18")

		// 方式三:
		name, ok := c.GetQuery("name") // 取到值返回(值, true),否则(值, false)
		if !ok {
			name = "body"
		}
		age, ok := c.GetQuery("age")
		if !ok {
			age = "18"
		}
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})
	r.Run(":9999")
}
