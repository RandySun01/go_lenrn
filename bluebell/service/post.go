package service

import (
	"bluebell/dao/mysql"
	"bluebell/models/modelPost"
	"bluebell/pkg/snowflake"

	"go.uber.org/zap"
)

/*
@author RandySun
@create 2022-01-18-9:07
*/
func CreatePost(p *modelPost.Post) (err error) {
	// 生成post id
	p.Id = snowflake.GenId()
	//保存到数据库
	return mysql.CreatePost(p)
}

// GetPostDetail 获取单条数据
func GetPostDetail(postId int64) (data *modelPost.ApiPostDetail, err error) {
	// 查询并组合接口想用的数据
	post, err := mysql.GetPostDetailById(postId)
	if err != nil {
		zap.L().Error("mysql.GetPostDetailById failed", zap.Int64("postId", postId), zap.Error(err))
		return
	}
	// 根据作者id查询作者信息
	user, err := mysql.GetUserById(post.AuthorId)
	if err != nil {
		zap.L().Error("mysql.GetUserById(post.AuthorId) failed", zap.Int64("author_id", post.AuthorId), zap.Error(err))
		return
	}

	// 根据社区id查询社区详细信息
	communityDetail, err := mysql.GetCommunityDetailById(post.CommunityId)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailById(post.CommunityId)", zap.Int64("community_id", post.CommunityId), zap.Error(err))
		return
	}
	data = &modelPost.ApiPostDetail{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: communityDetail,
	}
	return
}
