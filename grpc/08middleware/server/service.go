package main

import (
	"context"
	pb "grpc/06deadlines/proto"
	"log"
	"net"

	"google.golang.org/grpc/credentials"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
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
	// 添加拦截器后，方法里省略Token认证
	// //检测Token是否有效
	// if err := Check(ctx); err != nil {
	// 	return nil, err
	// }
	res := pb.SimpleResponse{
		Code:  200,
		Value: "hello " + req.Data,
	}
	return &res, nil
}

// Check 验证token
func Check(ctx context.Context) error {
	//从上下文中获取元数据
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "获取Token失败")
	}
	var (
		appID     string
		appSecret string
	)
	if value, ok := md["app_id"]; ok {
		appID = value[0]
	}
	if value, ok := md["app_secret"]; ok {
		appSecret = value[0]
	}
	if appID != "grpc_token" || appSecret != "123456" {
		return status.Errorf(codes.Unauthenticated, "Token无效: app_id=%s, app_secret=%s", appID, appSecret)
	}
	return nil
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

	// 从输入证书文件和密钥文件为服务端构造TLS凭证
	creds, err := credentials.NewServerTLSFromFile("../pkg/tls/server_cert.pem", "../pkg/tls/server_key.pem")
	if err != nil {
		log.Fatalf("Failed to generate credentials %v", err)
	}
	//普通方法：一元拦截器（grpc.UnaryInterceptor）
	var interceptor grpc.UnaryServerInterceptor
	interceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		//拦截普通方法请求，验证Token
		err = Check(ctx)
		if err != nil {
			return
		}
		// 继续处理请求
		return handler(ctx, req)
	}

	// 创建grpc服务实例
	grpcServer := grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(interceptor))

	// 在grpc服务器注册我们的服务
	pb.RegisterSimpleServer(grpcServer, &SimpleService{})
	log.Println(Address, "net.Listing whth TLS and token...")

	//用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatalf("grpcService.Serve err:%v", err)
	}
	log.Println("grpcService.Serve run succ")
}
