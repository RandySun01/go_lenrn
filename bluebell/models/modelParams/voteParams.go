package modelParams

/*
@author RandySun
@create 2022-01-19-9:03
*/

// ParamVote 投票数据
type ParamVote struct {
	//UserId
	PostId    string `json:"post_id" binding:"required"`              // 帖子id
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1"` // 赞成票(1),反对票(-1) 取消(0)
}
