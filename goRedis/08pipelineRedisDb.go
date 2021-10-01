package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

/*
@author RandySun
@create 2021-09-03-21:37
*/
func pipelineRedisDemo() {
	ctx := context.Background()
	pipe := rdb.Pipeline()
	incr := pipe.Incr(ctx, "pipeline_counter")
	pipe.Expire(ctx, "pipeline_counter", time.Hour)
	_, err := pipe.Exec(ctx)
	fmt.Println(incr.Val(), err)

}

func pipelinedRedisDemo() {
	var incr *redis.IntCmd
	ctx := context.Background()
	_, err := rdb.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		incr = pipe.Incr(ctx,"pipeline_counter")
		pipe.Expire(ctx, "pipeline_counter", time.Hour)
		return nil
	})

	fmt.Println(incr.Val(), err)
}