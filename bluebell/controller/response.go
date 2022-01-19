package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
@author RandySun
@create 2022-01-14-21:40
*/
type Response struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data,omitempty"` //omitempty:字段值为空忽略,不返回
	//Status int         `json:"status"`
}

// ResponseError 返回失败状态码
func ResponseError(c *gin.Context, code ResCode) {
	rd := &Response{
		Code: code,
		Msg:  code.GetMsg(),
		Data: nil,
	}
	c.JSON(http.StatusOK, rd)
}

// ResponseErrorWithMsg 自定义错误
func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}, data interface{}) {
	rd := &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	c.JSON(http.StatusOK, rd)
}

// ResponseSuccess 返回成功状态码
func ResponseSuccess(c *gin.Context, data interface{}) {
	rd := &Response{
		Code: CodeSuccess,
		Msg:  CodeSuccess.GetMsg(),
		Data: data,
	}
	c.JSON(http.StatusOK, rd)
}
