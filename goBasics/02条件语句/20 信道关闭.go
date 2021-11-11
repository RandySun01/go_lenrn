package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

//
//5 信道关闭 close(信道)

//func producer(chnl chan int) {
//	for i := 0; i < 10; i++ {
//		chnl <- i //一写就卡主
//	}
//	close(chnl) // 数据添加完了，需要关闭通道，不然取数据会死锁
//}
//func main() {
//	ch := make(chan int)
//	go producer(ch)
//	for { //死循环
//		v, ok := <-ch
//		if ok == false {
//			break
//		}
//		fmt.Println("Received ", v, ok)
//	}
//}

//6 通过for range循环信道,如果信道不关闭，可以一直取，一旦关闭，for循环结束

//func producer(chnl chan int) {
//	for i := 0; i < 10; i++ {
//		chnl <- i
//	}
//	close(chnl)
//}
//func main() {
//	ch := make(chan int)
//	go producer(ch)
//	//通过for range 就不用判断false了，当信道关闭，循环自动结束（代码量少）
//	for v := range ch {
//		fmt.Println("Received ", v)
//	}
//}

//7 重写那个例子

//func digits(number int, dchnl chan int) {
//	for number != 0 {
//		digit := number % 10
//		dchnl <- digit
//		number /= 10
//	}
//	close(dchnl)
//}
//func calcSquares(number int, squareop chan int) {
//	sum := 0
//	dch := make(chan int)
//	go digits(number, dch)
//	for digit := range dch {
//		sum += digit * digit
//	}
//	squareop <- sum
//}
//
//func calcCubes(number int, cubeop chan int) {
//	sum := 0
//	dch := make(chan int)
//	go digits(number, dch)
//	for digit := range dch {
//		sum += digit * digit * digit
//	}
//	cubeop <- sum
//}
//
//func main() {
//	number := 589
//	sqrch := make(chan int)
//	cubech := make(chan int)
//	go calcSquares(number, sqrch)
//	go calcCubes(number, cubech)
//	squares, cubes := <-sqrch, <-cubech
//	fmt.Println("Final output", squares+cubes)
//}

//func worker(id int, jobs<-chan int, results chan<- int) {
//	for j := range jobs {
//		fmt.Printf("worker:%d start job:%d\n", id, j)
//		time.Sleep(time.Second)
//		fmt.Printf("worker:%d end job:%d\n", id, j)
//		results<- j * 2
//	}
//}
//
//
//func main() {
//	jobs := make(chan int, 100)
//	results := make(chan int, 100)
//	// 开启3个goroutine
//	for w := 1; w<= 3; w++ {
//		go worker(w, jobs, results)
//	}
//	// 5个任务
//	for j := 1; j<= 5; j++ {
//		jobs<- j
//	}
//	close(jobs)
//	// 输出结果
//	for a := 1; a<= 5; a++ {
//		<-results
//	}
//}
////
//func main() {
//	ch := make(chan int, 1)
//	for i := 0; i < 10; i++ {
//		select {
//		case x := <-ch:
//			fmt.Println("x", x)
//		case ch <- i:
//			fmt.Println("i=", i)
//		default:
//			fmt.Println("default=", i)
//
//		}
//	}

//ch := make(chan int, 1)
//for i := 0; i< 10; i++ {
//	select {
//	case x :=<-ch:
//		fmt.Println(x)
//	case ch<- i:
//	}
//}
//}

//var x int64
//var wg sync.WaitGroup
//
//func add() {
//	for i := 0; i< 5000; i++ {
//		x = x + 1
//	}
//	wg.Done()
//}
//func main() {
//	wg.Add(2)
//	go add()
//	go add()
//	wg.Wait()
//	fmt.Println(x)
//}

//var x int64
//var wg1 sync.WaitGroup
//var lock sync.Mutex
//
//func add() {
//	for i := 0; i< 5000; i++ {
//		lock.Lock() // 加锁
//		x = x + 1
//		lock.Unlock() // 解锁
//	}
//	wg1.Done()
//}
//func main() {
//	wg1.Add(2)
//	go add()
//	go add()
//	wg1.Wait()
//	fmt.Println(x)
//}

//func test11(ch chan string) {
//	time.Sleep(time.Second * 5)
//	ch <- "test1"
//}
//func test2(ch chan string) {
//	time.Sleep(time.Second * 2)
//	ch <- "test2"
//}

//func main() {
//	// 2个管道
//	output1 := make(chan string)
//	output2 := make(chan string)
//	// 跑2个子协程，写数据
//	go test11(output1)
//	go test2(output2)
//	// 用select监控
//	select {
//	case s1 := <-output1:
//		fmt.Println("s1=", s1)
//	case s2 := <-output2:
//		fmt.Println("s2=", s2)
//	}
//}

//func main() {
//	// 创建2个管道
//	int_chan := make(chan int, 1)
//	string_chan := make(chan string, 1)
//	go func() {
//		//time.Sleep(2 * time.Second)
//		int_chan <- 1
//	}()
//	go func() {
//		string_chan <- "hello"
//	}()
//	select {
//	case value := <-int_chan:
//		fmt.Println("int:", value)
//	case value := <-string_chan:
//		fmt.Println("string:", value)
//	}
//	fmt.Println("main结束")
//}

//// 判断管道有没有存满
//func main() {
	// 创建管道
	//output1 := make(chan string)
	//// 子协程写数据
	//go write(output1)
	//// 取数据
	//for s := range output1 {
	//	fmt.Println("res:", s)
	//	time.Sleep(time.Second)
	//}

	//select {
	////// 写数据
	//case output1 <- "hello":
	////	fmt.Println("write hello")
	//default:
	//	fmt.Println("channel full")
	//}


//	quit := make(chan bool)
//	for i := 0; i != runtime.NumCPU(); i++ {
//		go func() {
//			for {
//				select {
//				case <-quit:
//					break
//				default:
//				}
//			}
//		}()
//	}
//
//	time.Sleep(time.Second * 15)
//	for i := 0; i != runtime.NumCPU(); i++ {
//		quit <- true
//	}
//}

//func write(ch chan string) {
//	for {
//		select {
//		// 写数据
//		case ch <- "hello":
//			fmt.Println("write hello")
//		default:
//			fmt.Println("channel full")
//		}
//		time.Sleep(time.Millisecond * 500)
//	}
//}

//var x int64
//var wg sync.WaitGroup
//
//func add() {
//	for i := 0; i< 5000; i++ {
//		x = x + 1
//	}
//	wg.Done()
//}
//func main() {
//	wg.Add(2)
//	go add()
//	go add()
//	wg.Wait()
//	fmt.Println(x)
//}
//
//var x int64
//var wg1 sync.WaitGroup
//var lock1 sync.Mutex
//
//func add() {
//	for i := 0; i< 5000; i++ {
//		lock1.Lock() // 加锁
//		x = x + 1
//		lock1.Unlock() // 解锁
//	}
//	wg1.Done()
//}
//func main() {
//	wg1.Add(2)
//	go add()
//	go add()
//	wg1.Wait()
//	fmt.Println(x)
//}

//var (
//	x      int64
//	wg     sync.WaitGroup
//	lock   sync.Mutex
//	rwlock sync.RWMutex
//)
//
//func write() {
//	//lock.Lock()   // 加互斥锁
//	rwlock.Lock() // 加写锁
//	x = x + 1
//	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒 173.2548ms
//	rwlock.Unlock()                   // 解写锁
//	//lock.Unlock()                     // 解互斥锁
//	wg.Done()
//}
//
//func read() {
//	//lock.Lock()                  // 加互斥锁
//	rwlock.RLock()               // 加读锁
//	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
//	rwlock.RUnlock()             // 解读锁
//	//lock.Unlock()                // 解互斥锁
//	wg.Done()
//}
//
//func main() {
//	start := time.Now()
//	for i := 0; i< 10; i++ {
//		wg.Add(1)
//		go write()
//	}
//
//	for i := 0; i< 1000; i++ {
//		wg.Add(1)
//		go read()
//	}
//
//	wg.Wait()
//	end := time.Now()
//	fmt.Println(end.Sub(start))
//}


var x int64
var l sync.Mutex
var wg sync.WaitGroup

// 普通版加函数
func add() {
	// x = x + 1
	x++ // 等价于上面的操作
	wg.Done()
}

// 互斥锁版加函数
func mutexAdd() {
	l.Lock()
	x++
	l.Unlock()
	wg.Done()
}

// 原子操作版加函数
func atomicAdd() {
	atomic.AddInt64(&x, 1)
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		//go add()       // 普通版add函数 不是并发安全的
		//go mutexAdd()  // 加锁版add函数 是并发安全的，但是加锁性能开销大
		go atomicAdd() // 原子操作版add函数 是并发安全，性能优于加锁版
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(x)
	fmt.Println(end.Sub(start))
}