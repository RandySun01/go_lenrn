package main

import "fmt"

func fun1() {
	fmt.Println("fun1")
}

func fun2(a int, b string) {
	fmt.Println(a, b)
}

func fun3(a, b string) {
	fmt.Println(a, b)
}

func fun4(a int) int {
	return a
}
func fun5(a int, b string) (int, string) {
	return a, b
}

func fun6(a int, b, c string) (int, string) {
	return a, b
}

func intSum2(x ...int) int {
	fmt.Println("T", x) //x是一个切片
	fmt.Printf("类型%T\n", x)

	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}

func test(a ...int) {
	//a是切片，但是不能直接传切片
	//如果想直接传切片，可以吗？
	fmt.Println(a)
	//查看变量a的类型（切片类型）
	fmt.Printf("%T", a)

}

func calc(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

func add(x, y int) int {
	return x + y
}
func calc1(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func adder() func(int) int {
	var x int
	fmt.Println("x", x)
	return func(y int) int {
		fmt.Println("y", y)
		x += y
		return x
	}
}

func f1() int {
	x := 5
	defer func() {
		x++
		//fmt.Println(x)
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func calc2(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
func test10() {

}
func main() {
	test10()
	/*
		func关键字  函数名(参数1 参数类型，参数2 参数类型)(){
		函数体的内容
		}

	*/
	//// 无惨函数
	//fun1()
	//// 有参函数
	//fun2(1, "23")
	//// 有参函数 升级版
	//fun3("1", "23")
	//
	//// 返回值
	//a := fun4(1)
	//fmt.Println(a)
	//
	//// 多个返回值
	//b, c := fun5(1, "123")
	//fmt.Println(b, c)
	//// y中 _ 是一个真正的变量，go中，就是没有
	//_, d := fun5(1, "123")
	//fmt.Println(d)
	//
	//fmt.Println(	fun6(1,"3","4")	)

	//print(intSum2(1, 2, 3, 4))
	//a := []int{1, 2, 3, 4}
	//fmt.Println()
	//test(a...)
	//sum, sub := calc(1,2)
	//fmt.Println(sum, sub)

	//ret2 := calc1(10, 20, add)
	//fmt.Println(ret2) //30

	//var f = adder()
	//fmt.Println(f(10)) //10
	//fmt.Println(f(20)) //30
	//f1 := adder()
	//fmt.Println(f1(40)) //40
	//fmt.Println(f1(50)) //90

	//fmt.Println(f1())
	//fmt.Println(f2())
	//fmt.Println(f3())
	//fmt.Println(f4())

	//x := 1
	//y := 2
	//defer calc2("AA", x, calc2("A", x, y))
	//x = 10
	//defer calc2("BB", x, calc2("B", x, y))
	//y = 20
	//x=3333
	//say()
	//
	//x:=222
	//{
	//	x:=333
	//	{
	//		x:=444
	//		fmt.Println(x)
	//	}
	//	fmt.Println(x)
	//}
	//fmt.Println(x)
	//say()

	type typeFunc func(int, int) int
	var name []typeFunc = []typeFunc{add1}
	for _, key := range name {
		fmt.Println(key(2, 4))

	}

}

var x int = 111

func say() {
	fmt.Println(x)
}

func add1(x, y int) int {
	return x + y
}
