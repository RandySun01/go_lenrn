package main

import (
	"fmt"
	"math/rand"
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
	//runtime.GOMAXPROCS(2)
	//go a1()
	//go b1()
	//time.Sleep(time.Second)

	// 需要2个管道
	// 1.job管道
	jobChan := make(chan *Job, 128)
	// 2.结果管道
	resultChan := make(chan *Result, 128)
	// 3.创建工作池
	createPool(64, jobChan, resultChan)
	// 4.开个打印的协程
	go func(resultChan chan *Result) {
		// 遍历结果管道打印
		for result := range resultChan {
			fmt.Printf("job id:%v randnum:%v result:%d\n", result.job.Id,
				result.job.RandNum, result.sum)
		}
	}(resultChan)
	var id int
	// 循环创建job，输入到管道
	for {
		id++
		// 生成随机数
		r_num := rand.Int()
		job := &Job{
			Id:      id,
			RandNum: r_num,
		}
		jobChan <- job
	}
}

type Job struct {
	// id
	Id int
	// 需要计算的随机数
	RandNum int
}

type Result struct {
	// 这里必须传对象实例
	job *Job
	// 求和
	sum int
}

// 创建工作池
// 参数1：开几个协程
func createPool(num int, jobChan chan *Job, resultChan chan *Result) {
	// 根据开协程个数，去跑运行
	for i := 0; i < num; i++ {
		go func(jobChan chan *Job, resultChan chan *Result) {
			// 执行运算
			// 遍历job管道所有数据，进行相加
			for job := range jobChan {
				// 随机数接过来
				r_num := job.RandNum
				// 随机数每一位相加
				// 定义返回值
				var sum int
				for r_num != 0 {
					tmp := r_num % 10
					sum += tmp
					r_num /= 10
				}
				// 想要的结果是Result
				r := &Result{
					job: job,
					sum: sum,
				}
				//运算结果扔到管道
				resultChan <- r
			}
		}(jobChan, resultChan)
	}
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
