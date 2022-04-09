package main

import (
	"context"
	"fmt"
	pb "grpc/12streamInterceptor/proto"
	"io"
	"log"
	"strconv"
	"time"

	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
)

/*
@author RandySun
@create 2022-03-28-11:17
*/

// Address 连接地址
const Address string = ":8000"

var streamClient pb.StreamClient

// conversations 调用服务端的Conversations方法
func conversations() {
	//调用服务端的Conversations方法，获取流
	stream, err := streamClient.Conversations(context.Background())
	if err != nil {
		log.Fatalf("get conversations stream err: %v", err)
	}
	for n := 0; n < 5; n++ {
		fmt.Println(12223)
		err := stream.Send(&pb.StreamRequest{Question: "stream client rpc " + strconv.Itoa(n)})
		if err != nil {
			log.Fatalf("stream request err: %v", err)
		}
		// 接收服务端消息
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Conversations get stream err: %v", err)
		}
		// 打印返回值
		log.Println(res.Answer)
	}
	//最后关闭流
	err = stream.CloseSend()
	if err != nil {
		log.Fatalf("Conversations close stream err: %v", err)
	}
}

// wrappedStream  用于包装 grpc.ClientStream 结构体并拦截其对应的方法。
type wrappedStream struct {
	grpc.ClientStream
}

func newWrappedStream(s grpc.ClientStream) grpc.ClientStream {
	return &wrappedStream{s}
}

func (w *wrappedStream) RecvMsg(m interface{}) error {
	log.Fatalf("Receive a message2 (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ClientStream.RecvMsg(m)
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	log.Fatalf("Send a message1 (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ClientStream.SendMsg(m)
}

// streamInterceptor 一个简单的 stream interceptor 示例。
func streamInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	s, err := streamer(ctx, desc, cc, method, opts...)
	if err != nil {
		return nil, err
	}
	log.Fatalf("一个简单的 client stream interceptor 示例")

	return newWrappedStream(s), nil
}
func main() {
	// 连接服务器
	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithStreamInterceptor(streamInterceptor))
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	// 建立gRPC连接
	streamClient = pb.NewStreamClient(conn)
	conversations()
}
