package main

import (
	"errors"
	"log"
	"net/http"
	"net/rpc"
)

/*
@author RandySun
@create 2022-03-20-16:55
*/

// 结构体, 用于注册
type Arith struct {
}

// 声明参数结构体
type ArithRequest struct {
	A, B int
}

// 返回结果给客户端的结果
type ArithResponse struct {
	// 乘积
	Pro int
	// 商
	Quo int
	// 余数
	Rem int
}

// 乘法
func (this *Arith) Multiply(req ArithRequest, res *ArithResponse) error {
	res.Pro = req.A * req.B
	return nil
}

// 商和余数
func (this *Arith) Divide(req ArithRequest, res *ArithResponse) error {
	if req.B == 0 {
		return errors.New("除数不能为0")
	}
	// 除
	res.Quo = req.A / req.B
	// 取模
	res.Rem = req.A % req.B
	return nil
}

func main() {
	// 注册服务
	rect := new(Arith)
	// 注册一个rect的服务
	rpc.Register(rect)
	// 服务处理绑定到http协议上面
	rpc.HandleHTTP()
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}

}
