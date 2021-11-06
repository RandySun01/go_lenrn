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
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n')
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" || strings.ToUpper(inputInfo) == "q"{
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