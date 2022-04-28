package main

import (
	"context"
	pb "grpc/14metadata/proto"
	"log"
	"time"

	"google.golang.org/grpc/keepalive"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/*
@author RandySun
@create 2022-04-10-20:03
*/

const (
	// Address 监听地址
	Address string = ":8001"
	// NetWork 网络通信协议
	NetWork string = "tcp"
)

var kacp = keepalive.ClientParameters{
	Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
	Timeout:             time.Second,      // wait 1 second for ping ack before considering the connection dead
	PermitWithoutStream: true,             // send pings even without active streams
}

func main() {
	// 连接服务器
	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithKeepaliveParams(kacp))
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
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	res, err := grpcClient.Route(ctx, &req)

	if err != nil {
		log.Fatalf("Call Route err:%v", err)
	}

	// 打印返回直
	log.Println("服务的返回响应data:", res)
	select {} // Block forever; run with GODEBUG=http2debug=2 to observe ping frames and GOAWAYs due to idleness.

}
