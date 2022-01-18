package controller

import (
	"errors"
	"strconv"

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

// getPageInfo 获取分页
func getPageInfo(c *gin.Context) (int64, int64, error) {
	pageStr := c.Query("page")
	sizeStr := c.Query("size")
	var (
		page int64
		size int64
		err  error
	)
	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}
	size, err = strconv.ParseInt(sizeStr, 10, 64)

	if err != nil {
		ResponseError(c, CodeInvalidParam)
		size = 10
	}
	return page, size, nil

}
