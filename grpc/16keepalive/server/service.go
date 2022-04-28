package main

import (
	"context"
	"flag"
	pb "grpc/13manyInterceptor/proto"
	"log"
	"net"
	"time"

	"google.golang.org/grpc/keepalive"

	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"
)

/*
@author RandySun
@create 2022-04-10-20:03
*/

var kaep = keepalive.EnforcementPolicy{
	MinTime:             5 * time.Second, // If a client pings more than once every 5 seconds, terminate the connection
	PermitWithoutStream: true,            // Allow pings even when there are no active streams
}

var kasp = keepalive.ServerParameters{
	MaxConnectionIdle:     15 * time.Second, // If a client is idle for 15 seconds, send a GOAWAY
	MaxConnectionAge:      30 * time.Second, // If any connection is alive for more than 30 seconds, send a GOAWAY
	MaxConnectionAgeGrace: 5 * time.Second,  // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
	Time:                  5 * time.Second,  // Ping the client if it is idle for 5 seconds to ensure the connection is still active
	Timeout:               1 * time.Second,  // Wait 1 second for the ping ack before assuming the connection is dead
}

// SimpleService 定义我们的服务
type SimpleService struct {
}

// Route 实现Route方法
func (s *SimpleService) Route(ctx context.Context, req *pb.SimpleRequest) (*pb.SimpleResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	log.Printf("md: %+v", md)
	res := pb.SimpleResponse{
		Code:  200,
		Value: "hello " + req.Data,
	}
	return &res, nil
}

const (
	// Address 监听地址
	Address string = ":8001"
	// NetWork 网络通信协议
	NetWork string = "tcp"
)

func main() {
	flag.Parse()

	// 监听本地端口
	listener, err := net.Listen(NetWork, Address)
	if err != nil {
		log.Fatalf("net.Listen err: %V", err)
	}
	log.Println(Address, "net.Listing...")
	// 创建grpc服务实例

	grpcServer := grpc.NewServer(grpc.KeepaliveEnforcementPolicy(kaep), grpc.KeepaliveParams(kasp))
	pb.RegisterSimpleServer(grpcServer, &SimpleService{})
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcService.Serve err:%v", err)
	}
	log.Println("grpcService.Serve run succ")
}
