package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)
// Decode解码消息
func Decode(reader *bufio.Reader)(string, error)  {
	// 读取消息长度
	lengthByte, _ := reader.Peek(4) // 读取前4个字节的数据
	lengthBuff := bytes.NewBuffer(lengthByte)
	var length int32
	err := binary.Read(lengthBuff, binary.LittleEndian, &length)
	if err !=nil{
		return "", err
	}
	// Buffered返回缓冲中现有的可读取的字节数。
	if int32(reader.Buffered()) < length+4{
		return "", err
	}
	// 读取真正的消息
	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil{
		return "", err
	}

	return string(pack[4:]), nil

}


func Process3(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {

		msg, err := Decode(reader)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		fmt.Println("收到client发来的数据：", msg)
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
		go Process3(conn)
	}
}
