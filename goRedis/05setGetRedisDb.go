package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8" // 注意导入的是新版本

)
/*
@author RandySun
@create 2021-09-02-8:21
*/

func setGetRedisDemo() {
	ctx := context.Background()

	err := rdb.Set(ctx, "score", 100, 0).Err()
	if err != nil{
		fmt.Printf("set score failed, err: %v\n", err)
	}

	val, err :=rdb.Get(ctx,"sscore").Result()
	if err != nil{
		fmt.Printf("get score failed, err: %v\n", err)

	}
	fmt.Println("score", val)

	val2, err := rdb.Get(ctx, "name").Result()
	if err == redis.Nil {
		fmt.Println("name does not exist")
	} else if err != nil {
		fmt.Printf("get name failed, err:%v\n", err)
		return
	} else {
		fmt.Println("name", val2)
	}
}