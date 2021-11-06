package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
)

func Encode1(message string)([]byte, error){
	// 读取消息的长度转换成int32类型(占用4个字节)
	length := int32(len(message))
	pkg := new(bytes.Buffer)
	// 写入消息头
	err := binary.Write(pkg, binary.LittleEndian, length)
	fmt.Println("写入消息头", err )
	if err != nil{
		return nil, err
	}
	// 写入消息体
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	fmt.Println("写入消息体", err )
	if err != nil{
		return nil, err
	}
	return pkg.Bytes(), nil
}


// Decode解码消息
func Decode1(reader *bufio.Reader)(string, error)  {
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