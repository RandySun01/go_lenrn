package modelParams

/*
@author RandySun
@create 2022-01-20-8:53
*/

// ParamPostList 获取帖子列表query string参数
type ParamPostList struct {
	Page  int64  `json:"page" from:"page"`
	Size  int64  `json:"size" from:"size"`
	Order string `json:"order" from:"order"`
}
