package controller

import (
	"bluebell/common"
	"bluebell/models/modelParams"
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
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, nil)
}

// GetPostDetailHandler 获取帖子电调数据
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

// GetPostListHandler 获取帖子分页
func GetPostListHandler(c *gin.Context) {
	// 获取分页参数
	page, size, err := getPageInfo(c)
	zap.L().Error(
		"service.GetPostList page size failed",
		zap.Int64("page", page),
		zap.Int64("size", size),
		zap.Error(err),
	)
	// 获取数据
	data, err := service.GetPostList(page, size)
	if err != nil {
		zap.L().Error("service.GetPostList failed", zap.Error(err))
		return
	}
	// 返回响应
	ResponseSuccess(c, data)

}

// GetPostListHandler2 升级版帖子列表接口
// 根据前端传来的参数动态的获取帖子的列表
// 按照时间排序 或者按照分数排序
// 1.获取分数
// 2. 去redis查询id列表
// 3.根据id去数据库查询帖子详细信息
func GetPostListHandler2(c *gin.Context) {
	// 获取分页参数 /api/v1/posts/?page=1&size=10&order=time

	// 初始化函数结构体时,指定初始化参数
	p := &modelParams.ParamPostList{
		Page:  1,
		Size:  10,
		Order: common.OrderTime, // magic string
	}

	// c.ShouldBind() 根据请求的数据类型选择相应的方法去获取数据
	// c.ShouldBindJSON() 如果请求中携带是json数据格式,才能用这个方法获取到数据
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetPostListHandler2 with invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	page, size, err := getPageInfo(c)
	zap.L().Error(
		"service.GetPostList page size failed",
		zap.Int64("page", page),
		zap.Int64("size", size),
		zap.Error(err),
	)
	// 获取数据
	data, err := service.GetPostList(page, size)
	if err != nil {
		zap.L().Error("service.GetPostList failed", zap.Error(err))
		return
	}
	// 返回响应
	ResponseSuccess(c, data)

}
