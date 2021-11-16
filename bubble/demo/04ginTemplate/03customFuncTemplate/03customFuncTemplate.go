package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	r := gin.Default()
	// 自定义函数在解析模板之前定义

	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})


	// 解析模板
	r.LoadHTMLGlob("./templates/**/*")
	r.GET("/custom/func/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "customFunc.tmpl", gin.H{
			// H is a shortcut for map[string]interface{}
			//type H map[string]interface{}
			// 渲染模板
			"title": "<a href='https://baidu.com'>自定义函数</a>",
		})

	})

	r.Run(":9999") // 启动项目

}
