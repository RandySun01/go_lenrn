package main

import (
	"fmt"
	"net"
)

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP: net.IPv4(0,0,0, 0),
		Port: 9999,
	})
	if err != nil{
		fmt.Println("连接服务端失败, err: ", err)
		return
	}
	defer socket.Close()
	sendDate := []byte("hello server")
	_, err = socket.Write(sendDate)
	if err != nil{
		fmt.Println("发送数据失败,err:", err)
	}
	data := make([]byte, 4096)
	n, remoteAddr, err := socket.ReadFromUDP(data)
	if err != nil{
		fmt.Println("接收数据失败err: ", err)
		return
	}
	fmt.Println("recv: %v addr: %v conut: %v \n", string(data[:]), remoteAddr, n)



}