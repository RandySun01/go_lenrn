package main

import "fmt"

//1 定义一个方法（需要跟结构体搭配）
type Person3 struct {
	Name string
	Age  int
}

////给结构体绑定方法Run方法
//
////def run(self)
////	print(self.name ,"在跑步")
////func (p Person1) Run()  {
////	fmt.Println(p.Name,"在跑步")
////}
////
////func (p Person1) speak()  {
////	fmt.Println(p.Name,"在说话")
////}

//5 匿名字段方法

type Hobby1 struct {
	id        int
	hobbyname string
}

type Person4 struct {
	name   string
	age    int
	Hobby1 //匿名字段
}

func (h Hobby1) printName() {
	fmt.Println(h.hobbyname)
}

func (p Person3) Speak() {
	fmt.Println(p.Name, "在说话")
}
//func (p Person3) Run1()  {
//	fmt.Println(p.Name,"在跑步")
//}
type Myint int8

func (i *Myint) add() {
	(*i)++ //自增1
}
func main() {
	//方法（面向对象中的概念） ：函数和方法
	//方法有特殊之处，绑定给对象的，在方法内部，可以修改对象

	//2 使用方法
	//	p:=Person3{"randy",19}
	//	p.Run()
		//p.speak()
		//p.sex=1  //上节课同学问的

	//3  为什么有了函数，还要方法？用函数完全能实现方法的功能（包括py中，函数完全能实现方法的功能）
	//	p:=Person3{"randy",19}
	//	//	p.Run()
	//	//	Run(p)

	//4 值接收器和指针接收器
		//值类型接收器，不会修改原来的
		//p := Person3{"randy", 19}
		//p.changeNmae("laowang")
		//fmt.Println(p) //这个打印出 {randy,19} 名字并没有改
		//指针接收器，修改原来的
		//p1 := Person3{"randy", 19}
		//p1.changeNmae1("laowang")
		//fmt.Println(p1) //这个打印出 {randy 19} 名字并没有改

	//5 匿名字段方法（方法提升）
	//p := Person4{"randy", 18, Hobby1{1, "篮球"}}

	// 调用打印hobby名字的方法printName
	//p.Hobby1.printName() //大家都想到
	//p.printName()        //直接点出来

	//面向对的继承Person1继承了Hobby1，对象.方法 可以直接用父类的方法

	//6 非结构体上的方法(需要重命名)
	//想给int 绑定一个add方法，以后每add一次，自加1
	//var a Myint=10
	//a.add()
	//a.add()
	//a.add()
	//fmt.Println(a)

}

func Run(p Person3) {
	fmt.Println(p.Name, "在跑步")
}
func (p Person3) Run() {
	fmt.Println(p.Name, "在跑步")
}

// 值接收器
func (p Person3) changeNmae(name string) {
	p.Name = name
	fmt.Println(p)
}

//指针类型接收器
func (p *Person3) changeNmae1(name string) {
	//(*p).Name=name
	p.Name = name
	fmt.Println(p)
}
