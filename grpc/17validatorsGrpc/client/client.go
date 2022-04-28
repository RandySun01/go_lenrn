package main

import (
	"context"
	pb "grpc/02simpleGrpc/proto"
	"log"

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
	// 调用 Route 方法 同时传入context.Context,  在有需要时可以让我们改变RPC的行为，比如超时/取消一个正在运行的RPC
	res, err := grpcClient.Route(context.Background(), &req)
	if err != nil {
		log.Fatalf("Call Route err:%v", err)
	}
	// 打印返回直
	log.Println("服务的返回响应data:", res)

}
