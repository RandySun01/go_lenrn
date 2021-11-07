package main

import "fmt"
//a := 200

func main() {
	// 定义变量必须使用
	// 方式一 var关键字 变量名 变量类型 = 变量值
	var a int = 10
	fmt.Println(a)
	a = 11
	fmt.Println(a)
	// 方式二 第二种（类型推导）类型可以省略

	var str = "randy"
	fmt.Println(str)
	str = "randysun"
	fmt.Println(str)

	// 方式三 第三（简略声明）冒号和等号是一起的
	c := 10
	fmt.Println(c)
	c = 12
	fmt.Println(c)
	// 演变
	// 先定义在执行
	var b int
	b = 123
	fmt.Println(b)
	// 声明多个变量

	var d, e int = 110, 123
	dd, ee := "randy", 123
	fmt.Println(d, e)
	fmt.Println(dd, ee)

	// 匿名变量
	_, aa := 1, 3
	fmt.Println(aa)
	//fmt.Println(_) 不为 _分配内存空间
	str = "sun"
	_, age1 := "randysun", 18
	fmt.Println( age1)
}
