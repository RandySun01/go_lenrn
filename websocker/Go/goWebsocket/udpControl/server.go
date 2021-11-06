package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP: net.IPv4(0,0,0,0),
		Port: 9999,
	})

	if err != nil{
		fmt.Println("listen failed, err: ", err)
	}
	defer listen.Close()
	for {
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil{
			fmt.Println("reD UDP FAILED, ERR:", err)
		}
		fmt.Println(data, 544444444444)
		fmt.Println("data: %v, addr: %v, count: %v\n", string(data[:n]), addr, n)
		_, err = listen.WriteToUDP(data[:n], addr) // 发送数据
		if err != nil{
			fmt.Println("write to upp failed, err: ", err)
			continue
		}
	}
}