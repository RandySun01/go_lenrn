package main

import (
	"fmt"
	pb "grpc/12streamInterceptor/proto"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

/*
@author RandySun
@create 2022-03-28-10:58
*/

// StreamService 定义我们的服务
type StreamService struct{}

// Conversations 实现Conversations方法
func (s *StreamService) Conversations(srv pb.Stream_ConversationsServer) error {
	n := 1
	for {
		fmt.Println(1111)
		req, err := srv.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}
		err = srv.Send(&pb.StreamResponse{
			Answer: "from stream server answer: the " + strconv.Itoa(n) + "question is " + req.Question,
		})

		if err != nil {
			return err
		}

		n++
		log.Printf("from stream client question: %s", req.Question)
	}
}

const (
	// Address 监听地址
	Address string = ":8000"
	// Network 网络通信协议
	Network string = "tcp"
)

type wrappedStream struct {
	grpc.ServerStream
}

func newWrappedStream(s grpc.ServerStream) grpc.ServerStream {
	return &wrappedStream{s}
}

func (w *wrappedStream) RecvMsg(m interface{}) error {
	log.Fatalf("Receive a message (Type: %T) at %s", m, time.Now().Format(time.RFC3339))
	return w.ServerStream.RecvMsg(m)
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	log.Fatalf("Send a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ServerStream.SendMsg(m)
}

func streamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	// 包装 grpc.ServerStream 以替换 RecvMsg SendMsg这两个方法。
	err := handler(srv, newWrappedStream(ss))
	if err != nil {
		log.Fatalf("RPC failed with error %v", err)
	}
	log.Fatalf("一个简单的 server stream interceptor 示例")

	return err
}

func main() {
	// 监听本地端口
	listener, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	log.Println(Address + " net.Listing...")
	// 新建gRPC服务器实例
	grpcServer := grpc.NewServer(grpc.StreamInterceptor(streamInterceptor))
	// 在gRPC服务器注册我们的服务
	pb.RegisterStreamServer(grpcServer, &StreamService{})

	//用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}
