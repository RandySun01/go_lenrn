package main

import (
	"context"
	pb "grpc/02simpleGrpc/proto"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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
)

var grpcClient pb.SimpleClient

func main() {
	// 连接服务器
	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("net.Connect connect: %v", err)
	}
	defer conn.Close()
	// 建立gRpc连接
	ctx := context.Background()
	grpcClient = pb.NewSimpleClient(conn)
	route(ctx, 2)

}
func route(ctx context.Context, deadlines time.Duration) {
	// 设置超时时间
	clientDeadline := time.Now().Add(time.Duration(deadlines * time.Second))
	ctx, cancel := context.WithDeadline(ctx, clientDeadline)
	defer cancel()
	// 创建发送结构体
	req := pb.SimpleRequest{
		Data: "grpc",
	}
	// 调用 Route 方法 同时传入context.Context,  在有需要时可以让我们改变RPC的行为，比如超时/取消一个正在运行的RPC
	res, err := grpcClient.Route(ctx, &req)
	if err != nil {
		// 获取错误状态
		statu, ok := status.FromError(err)
		if ok {
			// 判断是否为调用超时
			if statu.Code() == codes.DeadlineExceeded {
				log.Fatalln("Route timeout!")
			}
		}

		log.Fatalf("Call Route err:%v", err)
	}
	// 打印返回直
	log.Println("服务的返回响应data:", res)

}
