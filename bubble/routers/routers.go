package routers

import (
	"bubble/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// 加载静态文件
	r.Static("/static", "static")
	// 加载模板
	r.LoadHTMLGlob("./templates/*")
	r.GET("/", controller.IndexHandler)

	// 路由
	TodoRouters(r)

	return r
}
