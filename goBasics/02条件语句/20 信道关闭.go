package main

import (
	"fmt"
	"sync"
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
//
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
//}


//var x int64
////var wg sync.WaitGroup
////
////func add() {
////	for i := 0; i< 5000; i++ {
////		x = x + 1
////	}
////	wg.Done()
////}
////func main() {
////	wg.Add(2)
////	go add()
////	go add()
////	wg.Wait()
////	fmt.Println(x)
////}


var x int64
var wg1 sync.WaitGroup
var lock sync.Mutex

func add() {
	for i := 0; i< 5000; i++ {
		lock.Lock() // 加锁
		x = x + 1
		lock.Unlock() // 解锁
	}
	wg1.Done()
}
func main() {
	wg1.Add(2)
	go add()
	go add()
	wg1.Wait()
	fmt.Println(x)
}