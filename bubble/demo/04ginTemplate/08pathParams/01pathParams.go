package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()

	// 获取请求path URL 参数
	// 注意URL的匹配不要冲突
	r.GET("/user/:name/:age", func(c *gin.Context) {
		// 获取路径参数
		name := c.Param("name")
		age := c.Param("age")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	r.GET("/blog/:year/:month/:day", func(c *gin.Context) {
		// 获取路径参数
		year := c.Param("year")
		month := c.Param("month")
		day := c.Param("day")
		c.JSON(http.StatusOK, gin.H{
			"year":  year,
			"month": month,
			"day":   day,
		})
	})
	r.Run(":9999")
}
