package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func Process2(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	var buf [1024]byte
	for {
		n, err := reader.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client发来的数据：", recvStr)
	}

}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:9999")
	fmt.Println("监听端口")
	if err != nil {
		fmt.Println("listen failed err: ", err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed err: ", err)
			continue
		}
		go Process2(conn)
	}
}
