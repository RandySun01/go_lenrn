package main

import (
	"context"
	"flag"
	"fmt"
	pb "grpc/13manyInterceptor/proto"
	"log"
	"net"
	"time"

	"google.golang.org/grpc/health"

	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

/*
@author RandySun
@create 2022-03-27-20:03
*/
var (
	port  = flag.Int("port", 50051, "the port to serve on")
	sleep = flag.Duration("sleep", time.Second*5, "duration between changes in health")

	system = "" // empty string represents the health of the system
)

// SimpleService 定义我们的服务
type SimpleService struct {
}

// Route 实现Route方法
func (s *SimpleService) Route(ctx context.Context, req *pb.SimpleRequest) (*pb.SimpleResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	log.Printf("md: %+v", md)
	ports := fmt.Sprintf(":%d", *port)
	res := pb.SimpleResponse{
		Code:  200,
		Value: "hello " + req.Data + ports,
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
	listener, err := net.Listen(NetWork, fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("net.Listen err: %V", err)
	}
	log.Println(Address, "net.Listing...")
	// 创建grpc服务实例

	grpcServer := grpc.NewServer()
	// 健康检查
	healthCheck := health.NewServer()
	healthpb.RegisterHealthServer(grpcServer, healthCheck)

	// 在grpc服务器注册我们的服务
	pb.RegisterSimpleServer(grpcServer, &SimpleService{})

	go func() {
		// asynchronously inspect dependencies and toggle serving status as needed
		next := healthpb.HealthCheckResponse_SERVING

		for {
			healthCheck.SetServingStatus(system, next)

			if next == healthpb.HealthCheckResponse_SERVING {
				next = healthpb.HealthCheckResponse_NOT_SERVING
			} else {
				next = healthpb.HealthCheckResponse_SERVING
			}

			time.Sleep(*sleep)
		}
	}()
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcService.Serve err:%v", err)
	}
	log.Println("grpcService.Serve run succ")
}
