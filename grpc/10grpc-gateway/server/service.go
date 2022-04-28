package main

import (
	"context"
	pb "grpc/10grpc-gateway/proto"
	"grpc/10grpc-gateway/server/gateway"
	"log"
	"net"

	"google.golang.org/grpc"
)

/*
@author RandySun
@create 2022-03-27-20:03
*/

// SimpleService 定义我们的服务
type SimpleService struct{}

// Route 实现Route方法
func (s *SimpleService) Route(ctx context.Context, req *pb.InnerMessage) (*pb.OuterMessage, error) {
	res := pb.OuterMessage{
		ImportantString: "hello grpc validator",
		Inner:           req,
	}
	return &res, nil
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

	// 新建gRPC服务器实例
	grpcServer := grpc.NewServer()
	// 在gRPC服务器注册我们的服务
	pb.RegisterSimpleServer(grpcServer, &SimpleService{})
	log.Println(Address + " net.Listing whth TLS and token...")
	// grpc-->http
	httpServer := gateway.ProvideHTTP(Address, grpcServer)
	// 证书认证 https
	//if err = httpServer.Serve(tls.NewListener(listener, httpServer.TLSConfig)); err != nil {
	//	log.Fatal("ListenAndServe: ", err)
	//}

	// 无需证书认证
	if err = httpServer.Serve(listener); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	//用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}
