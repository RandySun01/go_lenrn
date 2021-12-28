package main

import (
	"context"


	"github.com/go-redis/redis/v8" // 注意导入的是新版本
	"time"
)
/*
@author RandySun
@create 2021-09-02-8:15
*/
var clusterRdb *redis.ClusterClient

// redis集群
func clusterInitClient()(err error){
	clusterRdb = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{":6379", ":7001", ":7002", ":7003", ":7004", ":7005"},
	})
	ctx, cancel  :=context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel ()
	_, err = clusterRdb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}
