package service

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/models/modelParams"
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
	err = mysql.CreatePost(p)
	if err != nil {
		return err
	}
	// 记录保存到reds中
	err = redis.CreatePost(p.Id, p.CommunityId)
	return err
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

// GetPostList 获取帖子列表
func GetPostList(page, size int64) (data []*modelPost.ApiPostDetail, err error) {
	postList, err := mysql.GetPostList(page, size)
	if err != nil {
		zap.L().Error("mysql.GetPostList failed", zap.Error(err))
		return
	}
	data = make([]*modelPost.ApiPostDetail, 0, len(postList))
	for _, post := range postList {
		// 根据作者id查询作者信息
		user, err := mysql.GetUserById(post.AuthorId)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorId) failed", zap.Int64("author_id", post.AuthorId), zap.Error(err))
			continue
		}

		// 根据社区id查询社区详细信息
		communityDetail, err := mysql.GetCommunityDetailById(post.CommunityId)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailById(post.CommunityId)", zap.Int64("community_id", post.CommunityId), zap.Error(err))
			continue
		}
		postDetail := &modelPost.ApiPostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: communityDetail,
		}
		// 添加
		data = append(data, postDetail)

	}
	return
}

// GetPostList2 根据时间很分数获取帖子列表
func GetPostList2(p *modelParams.ParamPostList) (data []*modelPost.ApiPostDetail, err error) {
	// 去redis查询id列表
	ids, err := redis.GetPostIdsOrder(p)
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIdsOrder(p) return 0 data")
		return
	}
	zap.L().Debug("GetPostList2", zap.Any("ids", ids))
	// 根据postId到数据库中查询 返回的数据还要我提供的id顺序返回数据
	postList, err := mysql.GetPostListByIds(ids)
	data = make([]*modelPost.ApiPostDetail, 0, len(postList))

	// 提前查询好每篇帖子的投票数量
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return
	}

	//将帖子的作者和分区信息查询的结果查询出来填充到帖子当中
	for idx, post := range postList {
		// 根据作者id查询作者信息
		user, err := mysql.GetUserById(post.AuthorId)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorId) failed", zap.Int64("author_id", post.AuthorId), zap.Error(err))
			continue
		}

		// 根据社区id查询社区详细信息
		communityDetail, err := mysql.GetCommunityDetailById(post.CommunityId)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailById(post.CommunityId)", zap.Int64("community_id", post.CommunityId), zap.Error(err))
			continue
		}
		postDetail := &modelPost.ApiPostDetail{
			AuthorName:      user.Username,
			VoteNum:         voteData[idx],
			Post:            post,
			CommunityDetail: communityDetail,
		}
		// 添加
		data = append(data, postDetail)

	}
	return
}

func GetCommunityPostList(p *modelParams.ParamPostList) (data []*modelPost.ApiPostDetail, err error) {
	// 去redis查询id列表
	ids, err := redis.GetPostCommunityIdsOrder(p)
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIdsOrder(p) return 0 data")
		return
	}
	zap.L().Debug("GetPostList2", zap.Any("ids", ids))
	// 根据postId到数据库中查询 返回的数据还要我提供的id顺序返回数据
	postList, err := mysql.GetPostListByIds(ids)
	data = make([]*modelPost.ApiPostDetail, 0, len(postList))

	// 提前查询好每篇帖子的投票数量
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return
	}

	//将帖子的作者和分区信息查询的结果查询出来填充到帖子当中
	for idx, post := range postList {
		// 根据作者id查询作者信息
		user, err := mysql.GetUserById(post.AuthorId)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorId) failed", zap.Int64("author_id", post.AuthorId), zap.Error(err))
			continue
		}

		// 根据社区id查询社区详细信息
		communityDetail, err := mysql.GetCommunityDetailById(post.CommunityId)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailById(post.CommunityId)", zap.Int64("community_id", post.CommunityId), zap.Error(err))
			continue
		}
		postDetail := &modelPost.ApiPostDetail{
			AuthorName:      user.Username,
			VoteNum:         voteData[idx],
			Post:            post,
			CommunityDetail: communityDetail,
		}
		// 添加
		data = append(data, postDetail)

	}
	return
}

// GetPostListNew 获取所有帖子和根据社区查询帖子
func GetPostListNew(p *modelParams.ParamPostList) (data []*modelPost.ApiPostDetail, err error) {
	// 根据参数不同，执行不同的逻辑
	if p.CommunityId == 0 {
		//查询所有
		data, err = GetPostList2(p)
	} else {
		// 根据社区id查询
		data, err = GetCommunityPostList(p)
	}
	if err != nil {
		zap.L().Error("GetPostListNew failed", zap.Error(err))
		return nil, err
	}
	return
}
