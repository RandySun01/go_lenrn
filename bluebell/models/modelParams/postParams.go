package modelParams

/*
@author RandySun
@create 2022-01-20-8:53
*/

// ParamPostList 获取帖子列表query string参数
type ParamPostList struct {
	CommunityId int64  `json:"community_id" form:"community_id"` // 可以为空
	Page        int64  `json:"page" from:"page"`
	Size        int64  `json:"size" from:"size"`
	Order       string `json:"order" from:"order"`
}

// ParamCommunityPostList 按社区获取帖子列表query string参数
type ParamCommunityPostList struct {
	*ParamPostList
}
