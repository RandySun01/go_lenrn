package main

import "fmt"

/*
@author RandySun
@create 2021-09-01-8:44
*/

func main() {
	//err := ordinaryInitClient()
	//if err != nil{
	//	fmt.Println("连接redis failed:", err)
	//}
	//fmt.Println("连接redis success", oldRdb)

	//v8Example()
	//sentinelInitClient()
	//err := clusterInitClient()
	//fmt.Println(err)
	err := initClient()
	if err != nil{
		fmt.Println("连接redis failed:", err)
	}
	fmt.Println("连接redis success", rdb)
	//setGetRedisDemo()

	//zsetRedisDemo()
	//deleteKeysRedisDemo()

	//pipelineRedisDemo()
	//
	//
	//pipelinedRedisDemo()

	//transactionDemo()

	type TraceCode string
	a := TraceCode("sdfsd")
	fmt.Println(a)
	b := string("4")
	fmt.Println(b)
	//var b string("sdf")
	//fmt.Println(b)


}