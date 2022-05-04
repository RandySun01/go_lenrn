package main

import (
	"fmt"
)

//不同goroutine之间通信，怎么做？
func main() {
	//1 信道的定义，就是一个变量（特殊）,两个协程之间通信，传递什么类型（int，bool）
	//定义了一个可以传递int类型的信道
	//var a chan int

	//2 信道的0值  <nil>  引用类型，当做参数传递，直接就是操作该信道
	//var a chan int
	//fmt.Println(a)

	//3 信道的初始化
	//var a chan int=make(chan int)
	//fmt.Println(a)  // 打印出来是个地址：0xc000016120

	//4 使用，不通协程之间传递数据
	//var a chan int = make(chan int)
	//go test55(a) //test5执行多长时间，不清楚，睡多久不知道，这种方式有问题

	//把值从信道中取出
	//<-a  //取出来，没有赋值给一个变量
	//var i int = <-a //取出来，赋值给一个变量i
	//fmt.Print(i)
	//time.Sleep(3 * time.Second)
	//i := <-a //取出来，赋值给一个变量i  i是int类型   重点：如果信道中没有值，会阻塞，等另一个goroutine放进值取，再继续往下走
	//fmt.Println(i)

	//5 信道，默认都是阻塞的（重点）
	//本质，默认的信道，是放不进东西去，不管是放还是取，都阻塞
	ch := make(chan int, 4)
	var c chan  int
	c1 := make(chan int, 4)
	fmt.Println(c)
	fmt.Println(ch)
	close(c1)
	x, ok := <-c1
	fmt.Println(x,ok)
	//go recv(ch) // 启用goroutine从通道接收值

	ch<- 10
	fmt.Println("发送成功")
}

func test55(a chan int) {
	fmt.Println("go go go ")
	//当它执行完成，往信道放一个值
	a <- 1 //往信道中放值，因为它是管子，不能直接赋值方式放值,在这个位置，往里放值，也会阻塞,没有将数据取走也会报错
	fmt.Println("xxxx")

}

//6 例子，234   2的平方+3的平方+4的平方  +   2的立方+3的立方+4的立方
//计算平方和与立方和的和

func recv(c chan int) {
	ret, ok :=<-c
	fmt.Println("接收成功", ret, ok)
}
