package main

import "fmt"

//有缓冲信道(无缓冲，取值，赋值都是阻塞)
func main() {

	//0 无缓冲信道演示

	//var a chan int = make(chan int)
	//a <- 1 //一直阻塞在这，所以报死锁错误*/

	//1 定义有缓冲信道（信道里可以真正的放值，没放满之前，不阻塞）
	//定义了一个缓冲大小为3的信道
	//var a chan int=make(chan int,3)
	//a<-1   //放一个值
	//a<-2   //在放一个
	//a<-3   //在放一个
	////a<-4   //会怎么样？死锁
	//fmt.Println(<-a) //取出一个
	//fmt.Println(<-a) //取出一个
	//fmt.Println(<-a) //再取出一个，会死锁（会一直等待，等待死锁）
	//

	//2 无缓冲信道的本质是，缓冲大小为0 ，不是1：记住了
	//var a chan int=make(chan int,0)
	//a<-1
	//a<-3

	//3 长度和容量
	var a chan int = make(chan int, 3)
	a <- 1
	fmt.Println(len(a)) //1
	fmt.Println(cap(a)) //3

	//4 单向信道
	//var a chan int=make(chan int)
}
