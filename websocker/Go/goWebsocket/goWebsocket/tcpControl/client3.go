package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

// Encode 将消息编码
func Encode(message string)([]byte, error){
	// 读取消息的长度转换成int32类型(占用4个字节)
	var length = int32(len(message))
	var pkg = new(bytes.Buffer)
	// 写入消息头
	err := binary.Write(pkg, binary.LittleEndian, length)
	fmt.Println("写入消息头", err, message )
	if err != nil{
		return nil, err
	}
	// 写入消息体
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	fmt.Printf("写入消息体T%", err )
	if err != nil{
		return nil, err
	}
	return pkg.Bytes(), nil
}


func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil{
		fmt.Println("err: ", err)
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello, Hello. How are you?`
		data, err := Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed, err:", err)
			return
		}
		conn.Write(data)
	}
}