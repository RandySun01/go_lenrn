package main

import (
	"bufio"
	"fmt"
	"net"
)

// 处理函数
func Process(conn net.Conn){
	defer conn.Close() // 关闭连接
	for {
		reader := bufio.NewReader(conn)
		var buf [128] byte
		n, err := reader.Read(buf[:])  // 读取数据
		if err != nil{
			fmt.Println("Read From Client Failed, err: ", err)
		}

		recvStr := string(buf[:n])
		fmt.Println("收到Client端发来的数据: ", recvStr)
		conn.Write([]byte(recvStr)) // 发送数据
	}

}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:9999")
	fmt.Println("监听端口")
	if err !=nil{
		fmt.Println("listen failed err: ", err)
	}

	for {
		conn, err := listen.Accept()  // 建立连接
		if err != nil{
			fmt.Println("accept failed err: ", err)
			continue
		}
		go Process(conn) // 启动一个goroutine处理连接
	}
}

