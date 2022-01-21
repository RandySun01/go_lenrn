package redis

import (
	"context"
	"errors"
	"math"
	"strconv"
	"time"

	"go.uber.org/zap"

	"github.com/go-redis/redis/v8"
)

/*
@author RandySun
@create 2022-01-19-21:39
*/

/*
投票几种情况
赞成票(1),反对票(-1) 取消(0)
direction=1时,有两种情况:
	1. 之前没有投过票,现在投赞成票 --->更新分数和投票记录  差值绝对值: 1  +432
	2. 之前投反对票,现在改投赞成票 --->更新分数和投票记录  差值绝对值: 2  +432 * 2

direction=0时, 有两种情况:
	1. 之前投过反对票,现在取消投票 --->更新分数和投票记录 差值绝对值: 1  +432
	2. 之前投过赞成票,现在取消投票 --->更新分数和投票记录 差值绝对值: 1  -432
direction=-1时,有两种情况:
	1. 之前没有投过票,现在投反对票 --->更新分数和投票记录  差值绝对值: 1  -432
	2. 之前投赞成票, 现在改投反对票 --->更新分数和投票记录 差值绝对值: 2  -432 * 2

投票的限制:
每个帖子自发表之日起,一个星期之内允许用户投票,超过一个星期就不允许在投票
	1. 到期之后将redis保存的赞成票数和反对票数存储到mysql表中
	2. 到期之后删除保存的key：KeyPostVotedZSetPrefix

*/

const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePerVore     = 432 // 每一票值多少分
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
	ErrVoteRepeated   = errors.New("不允许重复投票")
)

// CreatePost 将创建的帖子放入到redis中
func CreatePost(postId, communityId int64) error {
	zap.L().Debug("CreatePost add redis time postId", zap.Int64("postId", postId))
	//事务
	pipeline := rdb.TxPipeline()
	// 帖子时间
	pipeline.ZAdd(context.Background(), getRedisKey(KeyPostTimeZSet), &redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postId,
	})

	// 帖子分数默认设置为当前创建帖子的时间戳,后续在当前时间戳加减分数
	pipeline.ZAdd(context.Background(), getRedisKey(KeyPostScoreZSet), &redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postId,
	})

	// 把帖子id加到社区的set中
	cKey := getRedisKey(KeyCommunitySetPrefix + strconv.Itoa(int(communityId)))
	pipeline.SAdd(context.Background(), cKey, postId)

	_, err := pipeline.Exec(context.Background())

	return err

}
func VoteForPost(userId, postId string, value float64) error {
	// 判断投票限制

	postTime := rdb.ZScore(context.Background(), getRedisKey(KeyPostTimeZSet), postId).Val()
	// 去redis取帖子发布时间,超过一星期不在投票
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}
	// 更新帖子的分数
	// 先查当前用户给当前帖子的投票记录 赞成票(1),反对票(-1) 取消(0)
	ov := rdb.ZScore(context.Background(), getRedisKey(KeyPostVotedZSetPrefix+postId), userId).Val()
	// 如果这一次投票的值和之前保存一直,就提示不允许重复投票
	if value == ov {
		return ErrVoteRepeated
	}
	var op float64
	if value > ov {
		op = 1
	} else {
		op = -1
	}
	// 计算两次投票的差值
	diff := math.Abs(ov - value)
	// 增加分数
	pipeline := rdb.Pipeline()
	pipeline.ZIncrBy(context.Background(), getRedisKey(KeyPostScoreZSet), op*diff*scorePerVore, postId)

	// 记录用户为该帖子投票的数据
	if value == 0 {
		pipeline.ZRem(context.Background(), getRedisKey(KeyPostVotedZSetPrefix+postId), userId)

	} else {
		pipeline.ZAdd(context.Background(), getRedisKey(KeyPostVotedZSetPrefix+postId), &redis.Z{
			Score:  value, // 赞成票和返回票
			Member: userId,
		})

	}
	_, err := pipeline.Exec(context.Background())
	return err
}
