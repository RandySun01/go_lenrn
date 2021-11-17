package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	// http重定向
	r.GET("/index", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")

	})

	// 路由重定向
	r.GET("/home", func(c *gin.Context) {
		// 指定重定向路由
		c.Request.URL.Path = "/homeRedirect"

		// 继续后续处理
		r.HandleContext(c)
	})

	r.GET("/homeRedirect", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "路由重定向",
		})
	})

	r.Run(":9999")
}
