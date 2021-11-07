package main

import "fmt"

//接口使用其他
//1 接口的值接收器和指针接收器，等同于方法的值接收器和指针接收器（不论指针接收器还是值接收器，都可以用指针或者值来调用）
//2 实现多个接口
//写了俩接口

type SpeakInterface interface {
	speak()
}
type RunInterface interface {
	run()
}

//定义了一个狗结构体
type Dogs struct {
	name string
	age  int
}

//让狗实现run接口，实现speak接口
func (d Dogs) speak() {
	fmt.Println(d.name, "在说话")
}
func (d Dogs) run() {
	fmt.Println(d.name, "在走路")
}

//3 接口嵌套

//定义一个hobby接口，有打印id和打印name方法
type Hobby3 interface {
	printId()
	printName()
}

//再定义一个接口，嵌套Hobby接口
type TestInterface interface {
	Hobby3      //只写接口名字（Hobby接口内的所有方法，就相当于在这写了一遍）
	printWife() //这个接口相当于有三个方法
}

//定义一个结构体来实现上面的方法
type Person33 struct {
	name string
	age  int
}

//实现接口
func (p Person33) printId() {
	fmt.Println("pringid")
}
func (p Person33) printName() {
	fmt.Println("pringname")

}
func (p Person33) printWife() {
	fmt.Println("printwife")

}

func main() {
	//2 实现多个接口
	//var d Dogs=Dogs{"laowang",12}
	//var s SpeakInterface=d
	//var r RunInterface=d
	//s.speak()
	//r.run()

	//3 接口嵌套
	//var p Person33
	////var h Hobby3=p
	//var t TestInterface=p
	////h.printName()
	//t.printName()
	//t.printWife()
	//p.printId()

	//4 接口零值,  是nil，接口是引用类型
	var p Person33
	//var t TestInterface  =&p   //本质  为了函数传递时，节省空间
	var t TestInterface = p //以后用的多的
	fmt.Println(t)
	//为什么要将产生的对象赋值给接口来调用方法
}
