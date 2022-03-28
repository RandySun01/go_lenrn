package main

import (
	pb "grpc/03serverStream/proto"
	"log"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

/*
@author RandySun
@create 2022-03-27-20:41
*/

// StreamService 定义流服务
type StreamService struct {
}

// ListValue 实现ListValue方法
func (s *StreamService) ListValue(req *pb.SimpleRequest, srv pb.StreamServer_ListValueServer) error {
	for n := 0; n < 15; n++ {
		// 向流中发送消息,默认每次send送消息最大的长度为 math.MaxInt32 bytes
		err := srv.Send(&pb.StreamResponse{
			StreamValue: req.Data + strconv.Itoa(n),
		})
		if err != nil {
			return err
		}

		log.Println("发送数据n:", strconv.Itoa(n))
		time.Sleep(1 * time.Second)
	}
	return nil
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
	log.Println(Address + " net.Listing...")

	// 新建gRPC服务器实例
	// 默认单次接收最大消息长度为`1024*1024*4`bytes(4M)，单次发送消息最大长度为`math.MaxInt32`bytes
	// grpcServer := grpc.NewServer(grpc.MaxRecvMsgSize(1024*1024*4), grpc.MaxSendMsgSize(math.MaxInt32))
	grpcServer := grpc.NewServer()
	pb.RegisterStreamServerServer(grpcServer, &StreamService{})

	// 用服务器 serve() 方法已以及端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}

}
