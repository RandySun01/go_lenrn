package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)
var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println("Hello Goroutine!", i)
}
//这是一个普通函数
func testgo() {
	time.Sleep(1 * time.Second)
	fmt.Println("hello world")

}
func a() {
	for i := 1; i< 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i< 10; i++ {
		fmt.Println("B:", i)
	}
}

func main() {
	//1 让普通函数testgo并发起来
	//
	//本质上：其他语言一般很//在函数前加一个go关键字，不需要关注开的是线程还是协程，因为它内部自己处理，你只管用就行了
	//	//go testgo()
	//	//go testgo()
	//	//go testgo()
	//	//go testgo()
	//	//go testgo()
	//	//go testgo()
	//	//go testgo()
	//	//
	//	////跟py稍微有点区别
	//	//time.Sleep(3*time.Second)
	//	//少很少很少开进程，开线程就够了，其他语言的多线程就是多线程，可以利用多核优势，因为py有gil的限制，作者没办法了，给你搞了一个开进程
	//for i := 0; i< 10; i++ {
	//	wg.Add(1) // 启动一个goroutine就登记+1
	//	go hello(i)
	//}
	//wg.Wait() // 等待所有登记的goroutine都结束

	runtime.GOMAXPROCS(1)
	go a()
	go b()
	time.Sleep(time.Second)

}


//1s支持10个请求，  10个人进来了查10个表，每个人占用10s钟，在接下来10s钟，只有这10个人在访问网站
//10个人进来了查10个表，每个人占用0.001s钟,接下来10s中，可以允许多少人访问，读写分离

//celery异步  请求来了----10s--处理完
//			请求来了---》直接返回了     起了一个异步任务去处理10s该做做
