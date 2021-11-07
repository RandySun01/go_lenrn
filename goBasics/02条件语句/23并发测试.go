package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//// 合起来写
	//go func() {
	//	i := 0
	//	for {
	//		i++
	//		fmt.Printf("new goroutine: i = %d\n", i)
	//		time.Sleep(time.Second)
	//	}
	//}()
	//i := 0
	//for {
	//	i++
	//	fmt.Printf("main goroutine: i = %d\n", i)
	//	time.Sleep(time.Second)
	//	if i == 2 {
	//		break
	//	}
	//}

	//go func(s string) {
	//	for i := 0; i < 2; i++ {
	//		fmt.Println(s)
	//	}
	//}("world")
	//// 主协程
	//for i := 0; i < 2; i++ {
	//	// 切一下，再次分配任务
	//	runtime.Gosched()
	//	fmt.Println("hello")
	//}

	//go func() {
	//	defer fmt.Println("A.defer")
	//	func() {
	//		defer fmt.Println("B.defer")
	//		// 结束协程
	//		runtime.Goexit()
	//		defer fmt.Println("C.defer")
	//		fmt.Println("B")
	//	}()
	//	fmt.Println("A")
	//}()
	//for {
	//}

	//runtime.GOMAXPROCS(1)
	runtime.GOMAXPROCS(2)
	go a1()
	go b1()
	time.Sleep(time.Second)
}

func a1() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b1() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}
