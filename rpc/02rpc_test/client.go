package main

import (
	"fmt"
	"log"
	"net/rpc"
)

/*
@author RandySun
@create 2022-03-20-16:55
*/

type ArithRequest1 struct {
	A, B int
}

// 返回客户端的结果
type ArithResponse1 struct {

	// 乘积
	Pro int
	// 商
	Quo int
	// 余数
	Rem int
}

func main() {
	conn, err := rpc.DialHTTP("tcp", ":8000")
	if err != nil {
		log.Fatalln(err)
	}
	req := ArithRequest1{9, 2}
	var res ArithResponse1
	err = conn.Call("Arith.Multiply", req, &res)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%d * %d = %d\n", req.A, req.B, res.Pro)

	err = conn.Call("Arith.Divide", req, &res)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%d / %d 商 %d，余数 = %d\n", req.A, req.B, res.Quo, res.Rem)

}
