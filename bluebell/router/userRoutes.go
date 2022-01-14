package routes

import (
	"bluebell/controller"

	"github.com/gin-gonic/gin"
)

/*
@author RandySun
@create 2022-01-12-21:50
*/
func UserRouters(r *gin.Engine) {
	// 用户注册
	r.POST("/signup", controller.SignUpHandler)
	r.POST("/login", controller.LoginHandler)
}
