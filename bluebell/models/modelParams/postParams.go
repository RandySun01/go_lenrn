package modelParams

/*
@author RandySun
@create 2022-01-20-8:53
*/

// ParamPostList 获取帖子列表query string参数
type ParamPostList struct {
	CommunityId int64  `json:"community_id" form:"community_id"`   // 可以为空
	Page        int64  `json:"page" from:"page"`                   // 页码
	Size        int64  `json:"size" from:"size"`                   // 每页数量
	Order       string `json:"order" from:"order" example:"score"` // 排序依据
}

// ParamCommunityPostList 按社区获取帖子列表query string参数
type ParamCommunityPostList struct {
	*ParamPostList
}
