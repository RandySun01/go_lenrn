package main

/*
@author RandySun
@create 2021-09-03-9:02
*/
import (
	"context"
	"fmt"
	//"github.com/go-redis/redis/v8" // 注意导入的是新版本

)
func deleteKeysRedisDemo()  {
	ctx := context.Background()

	vals, err := rdb.Keys(ctx, "na*").Result()
	if err != nil {
		fmt.Printf("get keys failed, err:%v\n", err)
		return
	}
	fmt.Println("根据通配符获取key前:", vals)

	iter := rdb.Scan(ctx, 0, "na*", 0).Iterator()
	for iter.Next(ctx) {
		fmt.Println("删除key: ", iter.Val())

		err := rdb.Del(ctx, iter.Val()).Err()
		if err != nil {
			panic(err)
		}
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
	vals2, err := rdb.Keys(ctx, "na*").Result()
	if err != nil {
		fmt.Printf("get keys failed, err:%v\n", err)
		return
	}
	fmt.Println("根据通配符删除key后:", vals2)
}