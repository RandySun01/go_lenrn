package main

import (
	"context"
	pb "grpc/13manyInterceptor/proto"
	"log"
	"net"
	"time"

	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"
)

/*
@author RandySun
@create 2022-03-27-20:03
*/

// SimpleService 定义我们的服务
type SimpleService struct {
}

const (
	timestampFormat = time.StampNano
	streamingCount  = 10
)

// Route 实现Route方法
func (s *SimpleService) Route(ctx context.Context, req *pb.SimpleRequest) (*pb.SimpleResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	log.Printf("md: %+v", md)
	res := pb.SimpleResponse{
		Code:  200,
		Value: "hello " + req.Data,
	}
	trailer := metadata.Pairs("timestamp", time.Now().Format(timestampFormat))
	grpc.SetTrailer(ctx, trailer)
	header := metadata.New(map[string]string{"location": "MTV", "timestamp": time.Now().Format(timestampFormat)})
	grpc.SendHeader(ctx, header)
	return &res, nil
}

const (
	// Address 监听地址
	Address string = ":8001"
	// NetWork 网络通信协议
	NetWork string = "tcp"
)

func main() {
	// 监听本地端口
	listener, err := net.Listen(NetWork, Address)
	if err != nil {
		log.Fatalf("net.Listen err: %V", err)
	}
	log.Println(Address, "net.Listing...")
	// 创建grpc服务实例

	grpcServer := grpc.NewServer()
	// 在grpc服务器注册我们的服务
	pb.RegisterSimpleServer(grpcServer, &SimpleService{})

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcService.Serve err:%v", err)
	}
	log.Println("grpcService.Serve run succ")
}
