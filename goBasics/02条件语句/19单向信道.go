package main

//4 单向信道（只能读或只能写）
// 使用场景，传到另外的协程函数中，转成唯送信道（只能往里写），不允许在函数中读出值

func sendData(sendch chan<- int) { //接收一个唯送信道，只允许往里写
	sendch <- 10
	//<-sendch   // 取值就会报错
}

func main() {
	// 双向通道
	//sendch := make(chan int) //双向信道，可以放值，也可以取值
	//go sendData(sendch)
	//fmt.Println(<-sendch) //取值没问题

	// 单向通道
	sendch := make(chan<- int)
	go sendData(sendch)
	//fmt.Println(<-sendch) //取就报错

}
