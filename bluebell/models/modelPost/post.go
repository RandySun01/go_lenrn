package modelPost

import (
	"bluebell/models/modelCommunity"
	"time"
)

/*
@author RandySun
@create 2022-01-18-8:38
*/

// Post 帖子 内存对齐(相同字段放在一起)
type Post struct {
	Id          int64     `json:"id" db:"post_id"`
	AuthorId    int64     `json:"author_id" db:"author_id"`
	CommunityId int64     `json:"community_id" db:"community_id" binding:"required"`
	Status      int32     `json:"status" db:"status"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}

// ApiPostDetail 获取帖子的详情
type ApiPostDetail struct {
	AuthorName                      string                    `json:"author_name"`
	VoteNum                         int64                     `json:"vote_num"`
	*Post                                                     // 嵌入帖子的结构体
	*modelCommunity.CommunityDetail `json:"community_detail"` // 嵌入社区信息
}
