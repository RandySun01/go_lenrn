package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil{
		fmt.Println("err: ", err)
	}
	defer conn.Close() // 关闭连接
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n') // 读取用户输入
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" || strings.ToUpper(inputInfo) == "q"{ // 如果输入q就退出
			return
		}

		_, err = conn.Write([]byte(inputInfo)) // 发送数据
		if err != nil{
			return
		}

		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil{
			fmt.Println("recv failed , err: ", err)
		}
		fmt.Println(string(buf[:n]))
	}
}