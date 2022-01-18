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
func CommunityRouters(r *gin.Engine) {
	// 用户注册
	v1 := r.Group("/api/v1")
	v1.Use(middlewares.JWTAuthMiddleware())

	v1.GET("/community", controller.CommunityHandler)
	v1.GET("/community/:id", controller.CommunityDetailHandler)

}
