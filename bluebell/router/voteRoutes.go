package routes

import (
	"bluebell/controller"
	"bluebell/middlewares"

	"github.com/gin-gonic/gin"
)

/*
@author RandySun
@create 2022-01-19-8:58
*/
// VoteRouters 帖子相关
func VoteRouters(r *gin.Engine) {
	// 用户注册
	v1 := r.Group("/api/v1")
	v1.Use(middlewares.JWTAuthMiddleware())

	// 投票
	v1.POST("/vote", controller.PostVoteHandler)

}
