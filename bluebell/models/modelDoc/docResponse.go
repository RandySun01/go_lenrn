package modelDoc

import (
	"bluebell/controller"
	"bluebell/models/modelPost"
)

/*
@author RandySun
@create 2022-01-23-14:30
*/

// 专门用来放接口文档用到的model
// 因为我们的接口文档返回的数据格式是一致的,但是具体的data类型不一致
type _ResponsePostList struct {
	Code    controller.ResCode         `json:"code"`    // 业务响应状态码
	Message string                     `json:"message"` // 提示信息
	Data    []*modelPost.ApiPostDetail // 返回数据
}
