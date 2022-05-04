package main

import (
	"fmt"
)

// 接口  规范结构体的行为（只要结构体实现了该接口，）
type DuckInterface interface {
	speak()
	run()
}

//定义一个TDuck（唐老鸭）结构体来实现该接口
type TDuck struct {
	name string
	age  int
	wife string
}

//实现接口就是实现接口中所有方法

func (t TDuck) speak() {
	fmt.Println(t.name, "在说话")

}
func (t TDuck) run() {
	fmt.Println(t.name, "在怕不")
}

//定义一个PDuck（普通鸭子）结构体来实现该接口
type PDuck struct {
	name string
	age  int
}

//实现接口就是实现接口中所有方法
func (t PDuck) speak() {
	fmt.Println(t.name, "在说话")
}
func (t PDuck) run() {
	fmt.Println(t.name, "在走路")
}

// 3. 空接口
type Empty interface {
}

func main() {

	//t := TDuck{"laowang", 18, "laowang"}
	//t.speak()
	//

	//1 实例化得到两个鸭子对象
	//	var t TDuck = TDuck{name: "唐老鸭", age: 2, wife: "刘亦菲"}
	//	var p PDuck = PDuck{name: "普通鸭子", age: 1}
	//
	//	//2 这两个鸭子的对象(因为它实现了该接口)，都可以赋值给接口类型
	//	var i1 DuckInterface
	//	var i DuckInterface
	//	i1 = t
	//	fmt.Println(i1)
	//	fmt.Println(i1.speak)
	//	i1.speak()
	//	//i = p
	//	fmt.Println(i)

	//var _ DuckInterface = &PDuck{}


	//3 空接口（没有方法，所有类型其实都实现了空接口，于是：任意的类型都可以赋值给空接口类型）
	//var a Empty=10
	//a="ssddd"
	//a = TDuck{}
	//fmt.Println(a)

	//4 匿名空接口（空接口没有名字）
	//var a interface{} = "randy"
	//fmt.Println(a)

	//5 类型断言（接口类型传到函数中，它到底是PDuck还是TDuck，我需要断言回来）
	var t TDuck = TDuck{name: "唐老鸭", age: 2, wife: "刘亦菲"}
	var p PDuck = PDuck{name: "普通鸭子", age: 1}
	test(t)
	test(p)


	//6 类型选择
	//var t TDuck=TDuck{name:"唐老鸭",age:2,wife:"刘亦菲"}
	//var p PDuck=PDuck{name:"普通鸭子",age:1}
	//var a=10
	//var b="xxx"
	//test2(t)
	//test2(p)
	//test2(a)
	//test2(b)
	//test33(t)
	//test33(p)
	//test33(a)
	//test33(b)
	//var peo People1 = Student1{} // 不能通过
	//var peo People1 = &Student1{} // 能通过
	//think := "bitch"
	//fmt.Println(peo.Speak(think))

	//var w WashingMachine = &haier{}
	//w.wash()
	//w.dry()
	//var w io.Writer // io包内的Writer接口，初始零值/pair是(nil,nil)
	//
	//// 注意：接口值（一个接口类型的值，简称接口值）的类型是type存放的内容，即它所存放的具体变量的类型
	//fmt.Printf("%T %v\n", w, w)  // <nil> <nil>
	//d := time.Now()
	//fmt.Println(d)
}

// WashingMachine 洗衣机
type WashingMachine interface {
	wash()
	dry()
}

// 甩干器
type dryer struct{}

// 实现WashingMachine接口的dry()方法
func (d dryer) dry() {
	fmt.Println("甩一甩")
}

// 海尔洗衣机
type haier struct {
	dryer //嵌入甩干器
}

// 实现WashingMachine接口的wash()方法
func (h haier) wash() {
	fmt.Println("洗刷刷")
}



type People1 interface {
	Speak(string) string
}

type Student1 struct{}

func (stu *Student1) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

//5 类型断言
func test(d DuckInterface) {
	//暂时不知道他们的真正类型,接口类型只能调用接口的方法，不能调用原来结构体的属性和方法
	//fmt.Println(d.speak)
	//我可以断言你是什么类型，我断言你是TDuck类型   d.(TDuck) 返回两个值，一个是断言成功的TDuck类型的对象，一个是true和false表示是否断言成功
	//fmt.Println(d.(type))
	switch v := d.(type) {
	case TDuck:
		fmt.Printf("x is a string，value is %v\n", v)

	}
	value, ok := d.(TDuck)
	fmt.Printf("%T, %#v",value, ok )
	//if ok {
	//}
	//if v, ok := d.(TDuck); ok {
	//	fmt.Println(v.wife)
	//	fmt.Println(v.age)
	//	v.speak()
	//}
}


//6 类型选择其他
func test33(a interface{})  {
	fmt.Println(a)

	switch v:=a.(type) {
	//注意区分类型断言,这种情况只能用一个值来接收
	//因为a是空接口类型，它不能调用任何属性和方法
	case int:
		fmt.Println(a)
		fmt.Println(v)
	case string:
		fmt.Println(a)
		fmt.Println(v)
	case TDuck:
		fmt.Println(a)
		fmt.Println(v.wife)
	case PDuck:
		fmt.Println(a)
		fmt.Println(v.name)
	}
}

