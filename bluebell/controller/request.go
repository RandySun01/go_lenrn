package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
)

/*
@author RandySun
@create 2022-01-15-14:25
*/

const ContextUserIdKey = "userId"

var ErrorUserNotLogin = errors.New("用户未登录")

// GetCurrentUserId 获取用户信息
func GetCurrentUserId(c *gin.Context) (userId int64, err error) {
	uid, ok := c.Get(ContextUserIdKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	// 转换类型
	userId, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
