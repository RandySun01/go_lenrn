package routes

import (
	"bluebell/controller"
	"bluebell/middlewares"

	"github.com/gin-gonic/gin"
)

/*
@author RandySun
@create 2022-01-16-16:13
*/
func PostRouters(r *gin.Engine) {
	// 用户注册
	v1 := r.Group("/api/v1")
	v1.Use(middlewares.JWTAuthMiddleware())

	v1.POST("/post", controller.CreatePostHandler)
	v1.GET("/post/:id", controller.GetPostDetailHandler)
	v1.GET("/postList", controller.GetPostListHandler)
}
