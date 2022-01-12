package controller

import (
	"bluebell/models/params"
	"bluebell/service"
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

/*
@author RandySun
@create 2022-01-12-8:47
*/

func SignUpHandler(c *gin.Context) {
	// 1. 参数校验
	var p params.UserParamSignUp
	if err := c.ShouldBindJSON(&p); err != nil {
		// 请求参数有误,直接返回响应
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "请求参数有误",
		})
		return
	}
	fmt.Println(p)
	// 2.业务处理
	service.SignUp()
	// 3.返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
