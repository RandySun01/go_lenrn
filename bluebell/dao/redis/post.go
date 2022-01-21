package redis

import (
	"bluebell/common"
	"bluebell/models/modelParams"
	"context"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

/*
@author RandySun
@create 2022-01-21-8:46
*/

func getIdsFormKey(key string, page, size int64) ([]string, error) {
	// 确定查询索引的起始点
	start := (page - 1) * size
	end := start + size - 1
	//3 ZRevRange 查询分数从大到小顺序查询指定数量的元素
	return rdb.ZRevRange(context.Background(), key, start, end).Result()
}

// GetPostIdsOrder 根据用户查询的顺序获取post表的id
func GetPostIdsOrder(p *modelParams.ParamPostList) ([]string, error) {
	// 从redis获取id
	// 根据用户请求中携带的redis.cn
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == common.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)

	}
	//// 确定查询索引的起始点
	//start := (p.Page - 1) * p.Size
	//end := start + p.Size - 1
	//3 ZRevRange 查询分数从大到小顺序查询指定数量的元素
	return getIdsFormKey(key, p.Page, p.Size)
}

// GetPostVoteData 根据ids查询每篇帖子的投赞成票的数据
func GetPostVoteData(ids []string) (data []int64, err error) {
	//data = make([]int64, len(ids))
	//for _, id := range ids {
	//	key := getRedisKey(KeyPostVotedZSetPrefix + id)
	//	// 查找key中分数是1的元素数量--->统计每篇帖子的赞成票的数量
	//	v := rdb.ZCount(context.Background(), key, "1", "1").Val()
	//	data = append(data, v)
	//}

	// 使用pipeline一次发送多条命令,检索rtt  避免重复访问数据库
	pipeline := rdb.Pipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZSetPrefix + id)
		pipeline.ZCount(context.Background(), key, "1", "1")
	}
	// 执行redis命令
	cmders, err := pipeline.Exec(context.Background())
	if err != nil {
		return nil, err
	}
	// 获取数据
	data = make([]int64, 0, len(ids))
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return
}

// GetPostCommunityIdsOrder 按社区根据ids查询每篇帖子的投赞成票的数据
func GetPostCommunityIdsOrder(p *modelParams.ParamPostList) ([]string, error) {
	// 使用zinterstore 把分区的帖子set与帖子分数的zset 生成一个新的zset
	// 针对新的zset 按之前的逻辑取数据
	orderKey := getRedisKey(KeyPostTimeZSet)
	if p.Order == common.OrderScore {
		orderKey = getRedisKey(KeyPostScoreZSet)

	}
	//社区的key
	cKey := getRedisKey(KeyCommunitySetPrefix + orderKey + strconv.Itoa(int(p.CommunityId)))

	// 利用缓存key减少zinterstore执行的次数
	key := orderKey + strconv.Itoa(int(p.CommunityId))

	if rdb.Exists(context.Background(), key).Val() < 1 {
		// 不存在需要计算
		pipeline := rdb.Pipeline()
		weights, _ := strconv.ParseFloat(orderKey, 64)

		weightsList := []float64{weights}
		keys := []string{cKey}
		// 将公共的ids设置到缓存中
		pipeline.ZInterStore(context.Background(), key, &redis.ZStore{
			Keys:      keys,
			Weights:   weightsList,
			Aggregate: "MAX",
		}) // zinterstore 计算
		pipeline.Expire(context.Background(), key, 60*time.Second) // 设置超时时间
		_, err := pipeline.Exec(context.Background())
		if err != nil {
			return nil, err
		}
	}
	// 存在的话就直接根据Key查询ids
	return getIdsFormKey(key, p.Page, p.Size)

}
