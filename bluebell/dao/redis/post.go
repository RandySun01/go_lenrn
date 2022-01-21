package redis

import (
	"bluebell/common"
	"bluebell/models/modelParams"
	"context"
)

/*
@author RandySun
@create 2022-01-21-8:46
*/

// GetPostIdsOrder 根据用户查询的顺序获取post表的id
func GetPostIdsOrder(p *modelParams.ParamPostList) ([]string, error) {
	// 从redis获取id
	// 根据用户请求中携带的redis.cn
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == common.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)

	}
	// 确定查询索引的起始点
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1
	//3 ZRevRange 查询分数从大到小顺序查询指定数量的元素
	return rdb.ZRevRange(context.Background(), key, start, end).Result()
}
