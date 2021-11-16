package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// 解析模板
	r.LoadHTMLGlob("./templates/**/*")
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			// H is a shortcut for map[string]interface{}
			//type H map[string]interface{}
			// 渲染模板
			"title": "posts/index.tmpl",
		})

	})
	r.GET("/users/index", func(c *gin.Context) {

		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			// H is a shortcut for map[string]interface{}
			//type H map[string]interface{}
			// 渲染模板
			"title": "users/index.tmpl",
		})

	})

	r.Run(":9999") // 启动项目

}
