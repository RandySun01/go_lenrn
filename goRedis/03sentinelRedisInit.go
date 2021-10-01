package main
import (
	"context"

	"github.com/go-redis/redis/v8" // 注意导入的是新版本
	"time"
)
/*
@author RandySun
@create 2021-09-02-8:10
*/
var sentinelRdb *redis.Client
func sentinelInitClient()(err error){
	sentinelRdb = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName: "master",
		SentinelAddrs: []string{"127.0.0.1:6379", "127.0.0.1:6379", "127.0.0.1:6379"},
	})
	ctx, cancel  :=context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel ()
	_, err = sentinelRdb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}
