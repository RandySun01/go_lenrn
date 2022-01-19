package controller

import (
	"bluebell/models/modelParams"
	"bluebell/service"

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
			ResponseError(c, CodeInvalidParam)
		}
		errData := removeTopStruct(errs.Translate(trans)) // 翻译并去除掉错误提示中的结构体
		ResponseErrorWithMsg(c, CodeInvalidParam, errData, nil)
		return
	}

	service.PostVote()
	ResponseSuccess(c, nil)
}
