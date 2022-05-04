package main

import "fmt"

func main() {

	// 常量定义为不可比变的量
	const ROOT = "root"
	const NAME = "124"
	fmt.Println(ROOT, NAME)
	//const (
	//	n1 = iota //0
	//	n2        //1
	//	n3        //2
	//	n4        //3
	//)
	//
	//const a = iota
	//const b = n2
	//fmt.Println(a)
	//fmt.Println(b)
	const (
		n1 = iota //0
		n2 = 100  //100
		n3 = iota //2
		n4        //3
	)
	const n5 = iota //0 # const关键字出现时将被重置为0
	const n6 = iota //0
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)
	fmt.Println(n5)
	fmt.Println(n6)
	const (
		aa, bb = iota + 1, iota + 2 //1,2
		c, d                      //2,3
		e, f                      //3,4 e 会与上一个值相同
	)

	fmt.Println(aa,bb)
	fmt.Println(c,d)
	fmt.Println(e,f)


	const (
		a = 111 // 第一个声明的常量所在的行，此处虽然没有iota，但iota的值被置为0
		b // b的表达式同上，所以为111，此处虽然没有iota，但因为是新增一个常量声明，所以iota的值加1
		c1 // c的表达式同上，所以为111，此处虽然没有iota，但因为是新增一个常量声明，所以iota的值加1
		d1 = iota // 此处的iota基于前面的基础加1，所以d为3
		e1 = 333
		f1 // 同上为333
		g // 同上为333
	)
	fmt.Println(a,b,c1,d1,e1,f1,g) // 111 111 111 3 333 333 333

	var  name string
	var  age  byte
	var  sal  float32
	var ispass bool
	fmt.Println("请输入姓名：")
	//当程序只是到fmt.Scanln(&name)程序会停止执行等待用户输入
	fmt.Scanln(&name)
	fmt.Println("请输入年龄：")
	//当程序只是到fmt.Scanln(&age)程序会停止执行等待用户输入
	fmt.Scanln(&age)
	fmt.Println("请输入薪水：")
	//当程序只是到fmt.Scanln(&sal)程序会停止执行等待用户输入
	fmt.Scanln(&sal)
	fmt.Println("请输入是否通过：")
	//当程序只是到fmt.Scanln(&name)程序会停止执行等待用户输入
	fmt.Scanln(&ispass)
	fmt.Printf(" The name is:%s,age:%d,sal:%f, ispass :%t",name,age,sal,ispass)
}
