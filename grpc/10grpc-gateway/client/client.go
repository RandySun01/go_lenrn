package main

import (
	"context"
	pb "grpc/10grpc-gateway/proto"
	"log"

	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
)

/*
   @author RandySun
   @create 2022-04-28-11:51
*/
// Address 连接地址
const Address string = ":8000"

var grpcClient pb.SimpleClient

func main() {
	//从输入的证书文件中为客户端构造TLS凭证
	//creds, err := credentials.NewClientTLSFromFile("G:\\goproject\\go\\grpcGateway\\pkg\\tls\\ca_cert.pem", "x.test.example.com")
	//if err != nil {
	//	log.Fatalf("Failed to create TLS credentials %v", err)
	//}

	// 连接服务器 grpc.WithTransportCredentials(insecure.NewCredentials())
	//conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(&token))
	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	// 建立gRPC连接
	grpcClient = pb.NewSimpleClient(conn)
	route()
}

// route 调用服务端Route方法
func route() {
	// 创建发送结构体
	req := pb.InnerMessage{
		SomeInteger: 99,
		SomeFloat:   1,
	}
	// 调用我们的服务(Route方法)
	// 同时传入了一个 context.Context ，在有需要时可以让我们改变RPC的行为，比如超时/取消一个正在运行的RPC
	res, err := grpcClient.Route(context.Background(), &req)
	if err != nil {
		log.Fatalf("Call Route err: %v", err)
	}
	// 打印返回值
	log.Println(res)
}
