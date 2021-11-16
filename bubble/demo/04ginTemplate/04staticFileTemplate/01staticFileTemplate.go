package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	// 加载静态文件
	r.Static("/css", "../statics")

	r.LoadHTMLFiles("../templates/index.tmpl")

	r.GET("/index/css", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			// H is a shortcut for map[string]interface{}
			//type H map[string]interface{}
			// 渲染模板
			"title": "gin渲染模板",
		})
	})

	r.Run(":9999")
}
