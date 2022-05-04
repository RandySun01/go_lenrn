package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
@author RandySun
@create 2022-02-22-8:09
*/
var wg1 sync.WaitGroup

// 初始的例子

func worker1() {
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
	}
	// 如何接收外部命令实现退出
	wg1.Done()
}

func goWorker1()  {
	wg1.Add(1)
	go worker1()
	// 如何优雅的实现结束子goroutine
	wg1.Wait()
	fmt.Println("over")
}



var exit bool

// 全局变量方式存在的问题：
// 1. 使用全局变量在跨包调用时不容易统一
// 2. 如果worker中再启动goroutine，就不太好控制了。

func worker2() {
	for {
		fmt.Println("worker2")
		time.Sleep(time.Second)
		if exit {
			break
		}
	}
	wg1.Done()
}

func goWorker2(){
	wg1.Add(1)
	go worker2()
	time.Sleep(time.Second * 3) // sleep3秒以免程序过快退出
	exit = true                 // 修改全局变量实现子goroutine的退出
	wg1.Wait()
	fmt.Println("over2")
}


func worker3(exitChan chan struct{}) {
LOOP:
	for {
		fmt.Println("worker3")
		time.Sleep(time.Second)
		select {
		case <-exitChan: // 等待接收上级通知
			break LOOP
		default:
		}
	}
	wg1.Done()
}

func goWorker3(){
	var exitChan = make(chan struct{})
	wg1.Add(1)
	go worker3(exitChan)
	time.Sleep(time.Second * 3) // sleep3秒以免程序过快退出
	exitChan <- struct{}{}      // 给子goroutine发送退出信号
	close(exitChan)
	wg1.Wait()
	fmt.Println("over3")
}



func worker4(ctx context.Context) {
LOOP:
	for {
		fmt.Println("worker4")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // 等待上级通知
			break LOOP
		default:
		}
	}
	wg1.Done()
}

func goWorker4(){
	ctx, cancel := context.WithCancel(context.Background())
	wg1.Add(1)
	go worker4(ctx)
	time.Sleep(time.Second * 3)
	cancel() // 通知子goroutine结束
	wg1.Wait()
	fmt.Println("over4")
}


func worker5(ctx context.Context) {
	go worker51(ctx)
LOOP:
	for {
		fmt.Println("worker5")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // 等待上级通知
			break LOOP
		default:
		}
	}
	wg1.Done()
}
func worker51(ctx context.Context) {
LOOP:
	for {
		fmt.Println("worker51")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // 等待上级通知
			break LOOP
		default:
		}
	}
}
func goWorker5(){
	ctx, cancel := context.WithCancel(context.Background())
	wg1.Add(1)
	go worker5(ctx)
	time.Sleep(time.Second * 3)
	cancel() // 通知子goroutine结束
	wg1.Wait()
	fmt.Println("over5")
}

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return // return结束该goroutine，防止泄露
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}
func genWithCancel()  {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 当我们取完需要的整数后调用cancel

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}


func goWithDeadline(){
	d := time.Now().Add(10000 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// 尽管ctx会过期，但在任何情况下调用它的cancel函数都是很好的实践。
	// 如果不这样做，可能会使上下文及其父类存活的时间超过必要的时间。
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func goWorker6(ctx context.Context)  {
LOOP:
	for {
		fmt.Println("db connecting ...")
		time.Sleep(time.Millisecond * 10) // 假设正常连接数据库耗时10毫秒
		select {
		case <-ctx.Done(): // 50毫秒后自动调用
			break LOOP
		default:
		}
	}
	fmt.Println("worker done!")
	wg1.Done()
}
func goWithTimeout()  {
	// 设置一个50毫秒的超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	wg1.Add(1)
	go goWorker6(ctx)
	time.Sleep(time.Second * 5)
	cancel() // 通知子goroutine结束
	wg1.Wait()
	fmt.Println("over")
}

type TraceCode string


func goWorker7(ctx context.Context) {
	key := TraceCode("TRACE_CODE")
	traceCode, ok := ctx.Value(key).(string) // 在子goroutine中获取trace code
	if !ok {
		fmt.Println("invalid trace code")
	}
LOOP:
	for {
		fmt.Printf("worker, trace code:%s\n", traceCode)
		time.Sleep(time.Millisecond * 10) // 假设正常连接数据库耗时10毫秒
		select {
		case <-ctx.Done(): // 50毫秒后自动调用
			break LOOP
		default:
		}
	}
	fmt.Println("worker done!")
	wg1.Done()
}
func goWithValue()  {
	// 设置一个50毫秒的超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	// 在系统的入口中设置trace code传递给后续启动的goroutine实现日志数据聚合
	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "12512312234")
	wg1.Add(1)
	go goWorker7(ctx)
	time.Sleep(time.Second * 5)
	cancel() // 通知子goroutine结束
	wg1.Wait()
	fmt.Println("over")
}
func main() {
	//goWorker1()
	//goWorker2()
	//goWorker3()
	//goWorker4()
	//goWorker5()
	//genWithCancel()
	//goWithDeadline()
	//goWithTimeout()
	goWithValue()

}
