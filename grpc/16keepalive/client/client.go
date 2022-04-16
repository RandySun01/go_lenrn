package main

import (
	"context"
	"fmt"
	pb "grpc/14metadata/proto"
	"log"

	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/resolver/manual"

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

var serviceConfig = `{
	"loadBalancingPolicy": "round_robin",
	"healthCheckConfig": {
		"serviceName": ""
	}
}`

func main() {
	// 连接服务器
	r := manual.NewBuilderWithScheme("whatever")
	r.InitialState(resolver.State{
		Addresses: []resolver.Address{
			{Addr: "localhost:50051"},
			{Addr: "localhost:50052"},
		},
	})
	address := fmt.Sprintf("%s:///unused", r.Scheme())
	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithResolvers(r),
		grpc.WithDefaultServiceConfig(serviceConfig),
	}

	conn, err := grpc.Dial(address, options...)
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
	for {
		res, err := grpcClient.Route(context.Background(), &req)
		if err != nil {
			log.Fatalf("Call Route err:%v", err)
		}

		// 打印返回直
		log.Println("服务的返回响应data:", res)

	}

}
