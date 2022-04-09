package main

import (
	"context"
	pb "grpc/13manyInterceptor/proto"
	"log"
	"net"
	"runtime/debug"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"

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
	res := pb.SimpleResponse{
		Code:  200,
		Value: "hello " + req.Data,
	}
	return &res, nil
}

// LoggingInterceptor RPC 方法的入参出参的日志输出
func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("gRPC method: %s, %v", info.FullMethod, req)
	resp, err := handler(ctx, req)
	log.Printf("gRPC method: %s, %v", info.FullMethod, resp)
	return resp, err
}

// RecoveryInterceptor RPC 方法的异常保护和日志输出
func RecoveryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	//RPC 方法的异常保护和日志输出
	defer func() {
		if e := recover(); e != nil {
			debug.PrintStack()
			err = status.Errorf(codes.Internal, "Panic err: %v", e)
		}
	}()

	return handler(ctx, req)
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
	opts := []grpc.ServerOption{
		grpc_middleware.WithUnaryServerChain(
			RecoveryInterceptor,
			LoggingInterceptor,
		),
	}
	grpcServer := grpc.NewServer(opts...)
	// 在grpc服务器注册我们的服务
	pb.RegisterSimpleServer(grpcServer, &SimpleService{})

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcService.Serve err:%v", err)
	}
	log.Println("grpcService.Serve run succ")
}
