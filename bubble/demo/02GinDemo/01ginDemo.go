package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func satHelloWeb(c *gin.Context)  {
	// c.Json:返回JSON格式的数据
	c.JSONP(http.StatusOK, gin.H{
		"message": "helloWeb",
	})

}
func main() {

	// 创建一个默认的路由引擎
	r := gin.Default()
	// GET:请求方式: /helloWeb:请求路径
	r.GET("/helloWeb", satHelloWeb)
	// 启动HTTP服务,默认在0.0.0.0:8080服务
	r.Run()
}
