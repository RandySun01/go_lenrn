package main

import "fmt"

func calcSquares1(number int, squareop chan int) {
	sum := 0
	for number != 0 { //589   9的平方+8的平方+5的平方
		digit := number % 10 //取余数 --9    8  5
		sum += digit * digit
		number /= 10 //除以10等于  58        5
	}
	squareop <- sum
}

func calcCubes1(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}

func main() {
	number := 123
	//定义了两个信道
	sqrch := make(chan int)
	cubech := make(chan int)
	//起了两个协程，分别计算平方和和立方和
	go calcSquares1(number, sqrch)
	go calcCubes1(number, cubech)
	//squares, cubes := <-sqrch, <-cubech
	//从信道中取出值，赋给两个变量  cubes是什么类型？  int
	cubes := <-sqrch
	squares := <-cubech

	fmt.Println("Final output", squares+cubes)
}
