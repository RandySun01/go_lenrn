package routes

import (
	"bluebell/controller"
	"bluebell/middlewares"

	"github.com/gin-gonic/gin"
)

/*
@author RandySun
@create 2022-01-12-21:50
*/

// UserRouters 用户相关
func UserRouters(r *gin.Engine) {
	// 用户注册
	v1 := r.Group("/api/v1")
	v1.POST("/signup", controller.SignUpHandler)
	v1.POST("/login", controller.LoginHandler)
	v1.Use(middlewares.JWTAuthMiddleware())
	v1.POST("/ping", middlewares.JWTAuthMiddleware(), controller.LoginPingHandler)
}
