package main

import "fmt"

/*
@author RandySun
@create 2021-10-16-22:31
*/
func f5() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f6() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f7() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f8() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func main() {
	//fmt.Println(f5())
	//fmt.Println(f6())
	//fmt.Println(f7())
	//fmt.Println(f8())

	x := 1
	y := 2
	defer calc11("AA", x, calc11("A", x, y))
	x = 10
	defer calc11("BB", x, calc11("B", x, y))
	y = 20
}

func calc11(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}