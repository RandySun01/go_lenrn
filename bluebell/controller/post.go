package controller

import (
	"bluebell/models/modelPost"
	"bluebell/service"
	"strconv"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

/*
@author RandySun
@create 2022-01-18-8:39
*/
func CreatePostHandler(c *gin.Context) {
	// 获取参数及参数校验

	// c.ShouldBindJSON()  //调用 gin框架 validator ----> binding tag
	p := new(modelPost.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("c.ShouldBindJSON(p) error", zap.Any("err", err))
		zap.L().Error("post with invalid param", zap.Error(err))
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

	// 从c获取当前请求的用户id值
	authorId, err := GetCurrentUserId(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
	}
	p.AuthorId = authorId
	// 创建帖子
	if err := service.CreatePost(p); err != nil {
		zap.L().Error("service.CreatePost failed", zap.Error(err))
		return
	}
	// 返回响应
	ResponseSuccess(c, nil)
}

func GetPostDetailHandler(c *gin.Context) {
	// 获取参数(从URL中获取帖子的id)
	postIdStr := c.Param("id")
	postId, err := strconv.ParseInt(postIdStr, 10, 64)
	if err != nil {
		zap.L().Error("GetPostDetailHandler Get params invalid err", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 根据id取出帖子的数据(查数据库)
	data, err := service.GetPostDetail(postId)
	if err != nil {
		zap.L().Error("GetPostDetailHandler Get params err", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, data)

}
