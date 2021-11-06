package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	//1. timer基本使用
	//t0 := time.Now()
	//fmt.Printf("t0: %v \n", t0)
	//
	//timer1 := time.NewTimer(100* time.Second)
	//t1 := time.Now()
	//fmt.Printf("t1: %v \n", t1)
	//t2 := <-timer1.C
	//fmt.Printf("t2:%v\n", t2)

	// 2. 验证timer只能响应1次
	//timer2 := time.NewTimer(time.Second)
	//for {
	//	<-timer2.C
	//	fmt.Printf("时间到",)
	//}

	// 3. timer实现延迟的功能
	//time.Sleep(time.Second)
	//timer3 := time.NewTimer(2 * time.Second)
	//<- timer3.C
	//fmt.Printf("2秒到")
	//<- time.After(4*time.Second)
	//fmt.Printf("2秒到")

	// 4. 停止定时器
	//timer4 := time.NewTimer(200 * time.Second)
	//go func() {
	//	fmt.Println("定时器执行了")
	//	<-timer4.C
	//}()
	//
	//b := timer4.Stop()
	//if b{
	//	fmt.Println("timer4已经关闭")
	//	time.Sleep(2 * time.Second)
	//}
	// 5.重置定时器
	//timer5 := time.NewTimer(3 * time.Second)
	//timer5.Reset(1 * time.Second)
	//fmt.Println(time.Now())
	//fmt.Println(<-timer5.C)
	//for {
	//
	//}
	//Ticker：时间到了，多次执行
	// 1.获取ticker对象
	//ticker := time.NewTicker(2 * time.Second)
	//i :=0
	//// 子协程
	//go func() {
	//	for {
	//		i++
	//		fmt.Println(<-ticker.C)
	//		if i == 5{
	//			// 停止
	//			ticker.Stop()
	//		}
	//	}
	//}()
	//for  {
	//
	//}
	//
	// 创建2个管道
	//outPut1 := make(chan string)
	//outPut2 := make(chan string)
	//// 跑2个子协程,写数据
	//go test1(outPut1)
	//go test2(outPut2)
	//select {
	//case s1 := <-outPut1:
	//	fmt.Println("s1=",s1)
	//
	//case s2 := <-outPut2:
	//	fmt.Println("s2=", s2)
	//
	//}

	// 创建2个管道
	//int_chan := make(chan int, 1)
	//string_chan := make(chan string, 1)
	//go func() {
	//	//time.Sleep(2 * time.Second)
	//	int_chan <- 1
	//}()
	//go func() {
	//	string_chan <- "hello"
	//}()
	//select {
	//case value := <-int_chan:
	//	fmt.Println("int:", value)
	//case value := <-string_chan:
	//	fmt.Println("string:", value)
	//}
	//fmt.Println("main结束")
	// 创建管道
	//output1 := make(chan string, 10)
	//// 子协程写数据
	//go write(output1)
	//// 取数据
	//for s := range output1 {
	//	fmt.Println("res:", s)
	//	time.Sleep(time.Second)
	//}

	//wg.Add(1)
	//go hell()
	//fmt.Println("main goroutine done!")
	//wg.Wait()
	//
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			fmt.Println(key, 5555555555555)
			set(key, n)
			fmt.Printf("k=:%v,v:=%v\n", key, get(key))
			wg.Done()
		}(i)
	}
	wg.Wait()

	//var m = sync.Map{}
	//wg := sync.WaitGroup{}
	//for i := 0; i < 20; i++ {
	//	wg.Add(1)
	//	go func(n int) {
	//		key := strconv.Itoa(n)
	//		m.Store(key, n)
	//		value, _ := m.Load(key)
	//		fmt.Printf("k=:%v,v:=%v\n", key, value)
	//		wg.Done()
	//	}(i)
	//}
	//wg.Wait()
}

var m = make(map[string]int)

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

var wg sync.WaitGroup

func hell() {
	defer wg.Done()
	fmt.Println("Hello Goroutine!")
}
func write(ch chan string) {
	for {
		select {
		// 写数据
		case ch <- "hello":
			fmt.Println("write hello")
		default:
			fmt.Println("channel full")
		}
		time.Sleep(time.Millisecond * 500)
	}
}
func test1(ch chan string) {
	time.Sleep(time.Second * 5)
	ch <- "test1"

}
func test2(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "test2"
}
