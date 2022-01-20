package controller

import (
	"bluebell/models/modelParams"
	"bluebell/service"

	"go.uber.org/zap"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

/*
@author RandySun
@create 2022-01-19-9:00
*/

// PostVoteHandler 投票
func PostVoteHandler(c *gin.Context) {
	// 参数校验
	p := new(modelParams.ParamVote)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			zap.L().Error("service.VoteForPost(userId, p)", zap.Error(err))

			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans)) // 翻译并去除掉错误提示中的结构体

		ResponseErrorWithMsg(c, CodeInvalidParam, errData, nil)
		return
	}

	// 投票业务逻辑
	userId, err := GetCurrentUserId(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	// 具体投票逻辑
	if err := service.VoteForPost(userId, p); err != nil {
		zap.L().Error("service.VoteForPost(userId, p)", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
