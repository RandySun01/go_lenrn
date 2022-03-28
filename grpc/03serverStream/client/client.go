package main

import (
	"context"
	pb "grpc/03serverStream/proto"
	"io"
	"log"

	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
)

/*
@author RandySun
@create 2022-03-27-21:07
*/

// Address 连接地址
const Address string = ":8000"

var grpcClient pb.StreamServerClient

func listValue() {
	// 创建发送结构体
	req := pb.SimpleRequest{
		Data: "stream server grpc",
	}

	// 调用ListValue 方法 路由
	// 同时传入了一个 context.Context ，在有需要时可以让我们改变RPC的行为，比如超时/取消一个正在运行的RPC
	stream, err := grpcClient.ListValue(context.Background(), &req)
	if err != nil {
		log.Fatalf("Call ListStr err: %v", err)
	}
	for {
		// Recv() 方法接收服务端消息,默认每类Recv() 最大消息长度 1024*1024*4 bytes(4M)
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("ListStr get stream err: %v", err)
		}
		// 打印返回值
		log.Println(res.StreamValue)

		// 模拟终止服务端发送消息
		break // 退出接收数据
	}

	//可以使用CloseSend()关闭stream，这样服务端就不会继续产生流消息
	//调用CloseSend()后，若继续调用Recv()，会重新激活stream，接着之前结果获取消息
	stream.CloseSend()
	// 继续接收
	res, err := stream.Recv()

	if err == io.EOF {
		log.Fatalf("ListStr get stream err: %v", err)
	}
	if err != nil {
		log.Fatalf("ListStr get stream err: %v", err)
	}
	// 打印返回值
	log.Println(res.StreamValue)
}

func main() {
	// 连接服务器
	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	// 建立gRPC连接
	grpcClient = pb.NewStreamServerClient(conn)
	listValue()
}
