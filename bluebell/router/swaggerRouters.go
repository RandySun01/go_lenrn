package routes

import (
	_ "bluebell/docs" // 千万不要忘了导入把你上一步生成的docs

	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

/*
@author RandySun
@create 2022-01-23-14:39
*/

// SwaggerRouters 生成接口文档
func SwaggerRouters(r *gin.Engine) {
	// 用户注册
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
}
