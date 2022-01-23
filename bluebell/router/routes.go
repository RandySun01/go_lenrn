package routes

import (
	"bluebell/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
@author RandySun
@create 2022-01-11-8:31
*/

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	SwaggerRouters(r)
	// 路由
	VersionRouters(r)
	// 用户
	UserRouters(r)

	// 分类
	CommunityRouters(r)

	// 创建贴在
	PostRouters(r)

	// 投票
	VoteRouters(r)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": http.StatusNotFound,
		})

	})
	return r
}
