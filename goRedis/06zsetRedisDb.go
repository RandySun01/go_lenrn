package main

/*
@author RandySun
@create 2021-09-03-8:39
*/
import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8" // 注意导入的是新版本

)
func zsetRedisDemo() {

	zsetKey := "language_rank"
	languages := []*redis.Z{
	&redis.Z{Score: 90.0, Member: "Golang"},
	&redis.Z{Score: 98.0, Member: "Java"},
	&redis.Z{Score: 95.0, Member: "Python"},
	&redis.Z{Score: 97.0, Member: "JavaScript"},
	&redis.Z{Score: 99.0, Member: "C/C++"},
	}
	ctx := context.Background()

	// ZADD
	num, err := rdb.ZAdd(ctx,zsetKey, languages...).Result()
	if err != nil {
		fmt.Printf("zadd failed, err:%v\n", err)
		return
	}
	fmt.Printf("zadd %d succ.\n", num)


	// 把Golang的分数加10
	newScore, err := rdb.ZIncrBy(ctx, zsetKey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed, err:%v\n", err)
		return
	}

	fmt.Printf("Golang's score is %f now.\n", newScore)

	// 取分数最高的3个
	ret, err := rdb.ZRevRangeWithScores(ctx, zsetKey, 0, 2).Result()
	if err != nil {
		fmt.Printf("zrevrange failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println("取分数最高的3个",z.Member, z.Score)
	}

	// 取95~100分的
	op := &redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = rdb.ZRangeByScoreWithScores(ctx, zsetKey, op).Result()
	if err != nil {
		fmt.Printf("zrangebyscore failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println("取95~100分的:",z.Member, z.Score)
	}
}