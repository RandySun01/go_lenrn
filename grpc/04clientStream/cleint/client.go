package main

import (
	"context"
	pb "grpc/04clientStream/proto"
	"io"
	"log"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/*
@author RandySun
@create 2022-03-27-22:22
*/

// Address 连接地址
const Address string = ":8000"

var streamClient pb.StreamClientClient

// routeList 调用服务端RouteList方法
func routeList() {
	// 调用服务端RouteList方法,获取流
	stream, err := streamClient.RouteList(context.Background())
	if err != nil {
		log.Fatalf("Upload list err: %v", err)
	}

	for n := 0; n < 5; n++ {
		//向流中发送消息
		err := stream.Send(&pb.StreamRequest{
			StreamData: "stream client rpc " + strconv.Itoa(n),
		})
		// 添加判断服务端是否停止接口数据
		//发送也要检测EOF，当服务端在消息没接收完前主动调用SendAndClose()关闭stream，此时客户端还执行Send()，则会返回EOF错误，所以这里需要加上io.EOF判断
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("stream request err: %v", err)
		}
	}
	// 关闭流并获取返回消息
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("RouteList get response err: %v", err)
	}
	log.Println(res)
}

func main() {
	// 连接服务器
	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()
	// 建立gRPC连接
	streamClient = pb.NewStreamClientClient(conn)
	routeList()
}
