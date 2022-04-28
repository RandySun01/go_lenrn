package main

import (
	"context"
	"fmt"
	pb "grpc/14metadata/proto"
	"log"

	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/*
@author RandySun
@create 2022-03-27-20:03
*/

const (
	// Address 监听地址
	Address string = ":8001"
	// NetWork 网络通信协议
	NetWork string = "tcp"
)

func main() {
	// 连接服务器

	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("net.Connect connect: %v", err)
	}
	defer conn.Close()
	// 建立gRpc连接
	grpcClient := pb.NewSimpleClient(conn)

	// 创建发送结构体
	req := pb.SimpleRequest{
		Data: "grpc",
	}
	ctx := context.Background()
	//// 追加自定义字段
	//newCtx := metadata.AppendToOutgoingContext(ctx, "token", "RandySun")

	md := metadata.New(map[string]string{"go": "programming", "tour": "book"})
	newCtx := metadata.NewOutgoingContext(ctx, md)
	// 调用 Route 方法 同时传入context.Context,  在有需要时可以让我们改变RPC的行为，比如超时/取消一个正在运行的RPC
	var header, trailer metadata.MD

	res, err := grpcClient.Route(newCtx, &req, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		log.Fatalf("Call Route err:%v", err)
	}
	fmt.Println("timestamp from header:\n", header, trailer)

	if t, ok := header["timestamp"]; ok {
		fmt.Printf("timestamp from header:\n")
		for i, e := range t {
			fmt.Printf(" %d. %s\n", i, e)
		}
	} else {
		log.Fatal("timestamp expected but doesn't exist in header")
	}
	if l, ok := header["location"]; ok {
		fmt.Printf("location from header:\n")
		for i, e := range l {
			fmt.Printf(" %d. %s\n", i, e)
		}
	} else {
		log.Fatal("location expected but doesn't exist in header")
	}
	fmt.Printf("response:\n")

	if t, ok := trailer["timestamp"]; ok {
		fmt.Printf("timestamp from trailer:\n")
		for i, e := range t {
			fmt.Printf(" %d. %s\n", i, e)
		}
	} else {
		log.Fatal("timestamp expected but doesn't exist in trailer")
	}

	// 打印返回直
	log.Println("服务的返回响应data:", res)

}
