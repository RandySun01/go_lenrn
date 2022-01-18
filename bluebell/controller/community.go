package controller

import (
	"bluebell/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

/*
@author RandySun
@create 2022-01-16-16:15
*/

// CommunityHandler 社区相关
func CommunityHandler(c *gin.Context) {
	// 从数据库中读取所有社区(community_id, community_name)以列表形式返回
	data, err := service.GetCommunityList()
	if err != nil {
		zap.L().Error("service.GetCommunityList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不轻易将服务的错误暴露给外面
		return
	}
	// 返回给前端
	ResponseSuccess(c, data)

}

// CommunityDetailHandler 根据id获取详情
func CommunityDetailHandler(c *gin.Context) {

	//  获取社区id
	idStr := c.Param("id") // 获取url参数
	// 转换十六进制
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
	}
	data, err := service.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("community.CommunityDetailHandler failed", zap.Int64("id", id), zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不轻易将服务的错误暴露给外面
		return
	}
	ResponseSuccess(c, data)

}
