package main

import (
	pb "grpc/04clientStream/proto"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

/*
@author RandySun
@create 2022-03-27-22:22
*/

// SimpleService 定义服务
type SimpleService struct {
}

//RouteList 实现RouteList方法
func (s *SimpleService) RouteList(srv pb.StreamClient_RouteListServer) error {
	for {
		// 从流中获取信息
		res, err := srv.Recv()
		if err == io.EOF {
			// 发送结果并关闭
			return srv.SendAndClose(&pb.SimpleResponse{
				Code:  200,
				Value: "ok",
			})
		}
		if err != nil {
			return err
		}
		// 打印结果
		log.Println(res.StreamData)
		// 接收一次就停止接收数据
		return srv.SendAndClose(&pb.SimpleResponse{Value: "ok"})

	}
}

const (
	// Address 监听地址
	Address string = ":8000"
	// Network 网络通信协议
	Network string = "tcp"
)

func main() {
	// 监听本地端口
	listener, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	log.Println(Address + "net.Listing...")
	// 创建gRPC服务器实例
	grpcServer := grpc.NewServer()
	// gRPC服务器注册服务
	pb.RegisterStreamClientServer(grpcServer, &SimpleService{})

	//用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}
