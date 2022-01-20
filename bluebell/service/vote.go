package service

import (
	"bluebell/dao/redis"
	"bluebell/models/modelParams"
	"strconv"

	"go.uber.org/zap"
)

/*
@author RandySun
@create 2022-01-19-8:57
*/

// 基于用户投票相关算法: www.ruanyifeng.com/blog/algorithm

// redis实战
// 本项目使用简化版的投票分数
// 投一票就加432分 86400/200 -> 200张赞成票可以给帖子续一天 -->  《redis实战》

/*
投票几种情况
赞成票(1),反对票(-1) 取消(0)
direction=1时,有两种情况:
	1. 之前没有投过票,现在投赞成票 --->更新分数和投票记录
	2. 之前投反对票,现在改投赞成票 --->更新分数和投票记录

direction=0时, 有两种情况:
	1. 之前投过赞成票,现在取消投票 --->更新分数和投票记录
	2. 之前投过反对票,现在取消投票 --->更新分数和投票记录
direction=-1时,有两种情况:
	1. 之前没有投过票,现在投反对票 --->更新分数和投票记录
	2. 之前投赞成票, 现在改投反对票 --->更新分数和投票记录

投票的限制:
每个帖子自发表之日起,一个星期之内允许用户投票,超过一个星期就不允许在投票
	1. 到期之后将redis保存的赞成票数和反对票数存储到mysql表中
	2. 到期之后删除保存的key：KeyPostVotedZSetPrefix

*/

// VoteForPost 为帖子投票
func VoteForPost(userId int64, p *modelParams.ParamVote) error {
	// 记录帖子分数
	zap.L().Debug(
		"VoteForPost params",
		zap.Int64("userId", userId),
		zap.String("postId", p.PostId),
		zap.Int8("direction", p.Direction),
	)
	return redis.VoteForPost(strconv.Itoa(int(userId)), p.PostId, float64(p.Direction))

}
