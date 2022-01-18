package controller

import (
	"bluebell/dao/mysql"
	"bluebell/models/modelParams"
	"bluebell/service"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

/*
@author RandySun
@create 2022-01-12-8:47
*/

func SignUpHandler(c *gin.Context) {
	// 1. 参数校验
	var p = new(modelParams.UserParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		// 格式校验
		// 请求参数有误,直接返回响应
		zap.L().Error("SignUp with invalid param", zap.Error(err))

		// 判断err是不是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		// 翻译
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)), nil)
		return
	}
	fmt.Println(p)
	// 手动对请求参数进行详细的业务判断
	//if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.Password != p.RePassword {
	//	zap.L().Error("SignUp with invalid param")
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": "请求参数有误",
	//	})
	//	return
	//}
	// 2.业务处理
	if err := service.SignUp(p); err != nil {
		zap.L().Error("logic.SigUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3.返回响应
	ResponseSuccess(c, nil)
}

func LoginHandler(c *gin.Context) {
	// 获取请求参数校验
	var p = new(modelParams.UserParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		// 格式校验
		// 请求参数有误,直接返回响应
		zap.L().Error("Login with invalid param", zap.String("username", p.Username), zap.Error(err))

		// 判断err是不是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}

		// 翻译错误
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)), nil)
		return
	}
	fmt.Println(p)

	// 业务逻辑处理
	token, err := service.Login(p)
	if err != nil {
		zap.L().Error("logic.Login failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidParam)
		return
	}
	//  返回响应信息

	ResponseSuccess(c, token)
}

func LoginPingHandler(c *gin.Context) {
	// 如果是登录用户,判断请求头中是否是有效的JWT

	ResponseSuccess(c, nil)
}
