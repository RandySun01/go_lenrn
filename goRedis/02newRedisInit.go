package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8" // 注意导入的是新版本
	"time"
)

/*
@author RandySun
@create 2021-09-01-8:53
*/
var rdb *redis.Client

// 初始化连接
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		Password: "", // 密码
		DB: 5, // 选择连接的库
		PoolSize: 100, //连接池大小
	})

	ctx, cancel  :=context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel ()
	_, err = rdb.Ping(ctx).Result()
	return err
}


func v8Example(){
	ctx := context.Background()
	if err := initClient(); err != nil{
		return
	}
	err := rdb.Set(ctx, "name", "RandySun", 0).Err()
	if err != nil{
		panic(err)
	}
	val, err :=rdb.Get(ctx,"name").Result()
	if err != nil{
		panic(err)
	}
	fmt.Println("name", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil{
		fmt.Println("key2, does not exist")
	}else if err != nil{
		panic(err)
	}else {
		fmt.Println("key2", val2)
	}


}
