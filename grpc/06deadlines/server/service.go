package main

import (
	"context"
	pb "grpc/06deadlines/proto"
	"log"
	"net"
	"runtime"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

/*
@author RandySun
@create 2022-03-27-20:03
*/

// SimpleService 定义我们的服务
type SimpleService struct {
}

// Route 实现Route方法
func (s *SimpleService) Route(ctx context.Context, req *pb.SimpleRequest) (*pb.SimpleResponse, error) {
	data := make(chan *pb.SimpleResponse)

	go handle(ctx, req, data)
	select {
	case res := <-data:
		return res, nil
	case <-ctx.Done():
		return nil, status.Errorf(codes.Canceled, "Client cancelled abandoning")

	}
}

// 超时处理
func handle(ctx context.Context, req *pb.SimpleRequest, data chan<- *pb.SimpleResponse) {

	select {
	case <-ctx.Done():
		log.Println(ctx.Err())
		runtime.Goexit() // 超时后退出该Go协程
	case <-time.After(4 * time.Second): // 模拟耗时操作
		res := pb.SimpleResponse{
			Code:  200,
			Value: "hello " + req.Data,
		}
		data <- &res

	}
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
