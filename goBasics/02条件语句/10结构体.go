package main

import "fmt"

type Hobby struct {
	id   int
	name string
}
//type Person struct {
//	name string
//	// Name  string  // Name大写表示外部包可以被访问到
//	age int
//	sex int
//	//hobby Hobby
//}
type Person struct {
	name string
	// Name  string  // Name大写表示外部包可以被访问到
	age int
	sex int
	//hobby Hobby
	 Hobby // 字段没有名字
}

//8 匿名字段（字段没有名字，一种类型只能写一次）
type Person1 struct {
	string
	int
	//int  报错
}
//10 提升字段（匿名字段）
//写一个爱好结构体
//type Hobby struct {
//	id int
//	hobbyName string
//}
func main() {

	//1. 定义结构体对比面向对象中的实例化
	//p:=Person{}
	//var p Person   //结构体对象是值类型
	//实例化并赋初值
	/*	var p Persion = Persion{"randy",13,32}
		fmt.Println(p)
		fmt.Println(p.age)*/

	//2 实例化得到人对象  变量类型就是结构体的名字
	//var p Person=Person{}   //相当于调用__init__方法完成初始化
	//var p Person=Person{Name:"randy",age:18,sex:1}   //按关键字初始化
	//var p Person=Person{"lqz",18,1}   //按位置初始化
	//有什么区别？按关键字，可以忽略位置，可以传一部分
	//var p Person=Person{sex:1,age:18,Name:"randy"}
	//var p Person=Person{sex:1,age:18}
	//fmt.Println(p)

	//3 结构体类型的零值是：属性的空值   ，不是nil，它不是引用类型，当做参数传递，不会修改原来的值

		//var p Person  //只定义，没有初始化
		//var p Person = Person{name: "randy", age: 18, sex: 1}
		//fmt.Println(p)
		//stest5(p)
		//fmt.Println(p)

	//4 当参数传递，要修改原来的结构体
	//var p Person=Person{name:"randy",age:18,sex:1}
	//fmt.Println(p)
	//stest6(&p)
	//fmt.Println(p)

	//5 结构体属性获取和修改
	//var p Person=Person{name:"randy",age:18,sex:1}
	//fmt.Println(p.name)
	//p.age=20
	//fmt.Println(p)

	//6 匿名结构体(定义再函数内部，没有名字和type关键字)
	//	//一堆属性的集合  (其实就是面向对象的封装)
	//	a := struct {
	//		Name string
	//		age  int
	//	}{Name:"randy",age:19}
	//	//var a int
	//	//var b  string
	//	//var c [3]int
	//	fmt.Println(a.Name)

	//7 结构体指针
	//var p Person=Person{name:"randy",age:18,sex:1}
	//fmt.Println(p)
	//stest6(&p)
	////var p1 *Person=&Person{name:"randy",age:18,sex:1}
	//p1:=&Person{name:"randy",age:18,sex:1}
	//(*p1).name="xxx"  //正统用法
	//p1.name="yyy"   //推荐用法（别人的代码，这种多）
	//p2:=&Person{name:"randy",age:18,sex:1}
	//test6(p)
	//var p Person=Person{name:"randy",age:18,sex:1}
	//fmt.Printf("%T",p)

		//var p2 = new(Person)
		//fmt.Println(p2)
		//fmt.Printf("%T", p2)
		//p2.name = "randy"
		//fmt.Println(p2)
		//fmt.Printf("%v", p2)

	/*
		type student struct {
			name string
			age  int
		}

		m := make(map[string]*student)
		stus := []student{
			{name: "小王子", age: 18},
			{name: "娜扎", age: 23},
			{name: "大王八", age: 9000},
		}

		for _, stu := range stus {
			m[stu.name] = &stu
		}
		for k, v := range m {
			fmt.Println(k, "=>", v.name)
		}*/

	//8  匿名字段（字段没有名字）（暂时先不讲它有什么用，做变量提升，面向对象中的继承）
	//实例化，匿名字段名字就是类型名
	//var p1 Person1 = Person1{string: "randy", int: 18} //按关键字,有点奇特（这就是为什么一种类型只能写一次）
	//var p2 Person1 = Person1{"randys", 19}             //按位置
	//fmt.Println(p1)
	//fmt.Println(p2)

	/*h := Hobby{1, "randy"}
	p := Person{"randy", 12, 44, h}
	fmt.Println(p.hobby.id)
	p.hobby.id = 10
	fmt.Println(h.id)
	fmt.Println(p.hobby.id)
*/

	//9 结构体嵌套（结构体内套结构体）
/*	p1:=Person{name:"randy",age:19,sex:1,hobby:Hobby{id:1,name:"篮球"}}   //按关键字
	p2:=Person{"randy",19,1,Hobby{id:1,name:"篮球"}}   //按位置
	p3:=Person{}
	var p4 Person  //注意区分引用类型和值类型（数组，字符串，数字，只需要定义，不需要初始化就可以使用）
	//修改hobby的id
	p4.hobby.id=9
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
	fmt.Println(p4)*/


	//10 提升字段（匿名字段一起用）
	//p:=Person{name:"lqz",age:19,sex:1,Hobby:Hobby{id:1,name:"篮球"}}  //Hobby是类型名
	////神奇的一模发生了
	//fmt.Println(p.name)  //在外层可以调用内层的属性，提升了，匿名的时候提升
	//p.name="xxx"
	//p.Hobby.name="yyy"  //以后这两种方式都可以了
	//
	//你去想面向对象的继承  Person继承了Hobby  self.id可以直接拿到Hobby的id
	//Hobby中字段跟Person中字段重名了怎么办
	//p1:=&Person{name:"lqz",age:19,sex:1,Hobby:Hobby{id:1,name:"篮球"}}
	//fmt.Printf(p1.name)
	//fmt.Println(p1.Hobby.name)
	//p.Hobby=&Hobby{id:1,Name:"足球"}
	//var h1 Hobby
	//h1=Hobby{id:1,name:"足球"}
	//go的结构体嵌套就是面向对象的继承
	//
	//type Name struct {
	//	firstName string
	//	lastName string
	//}
	//type AA struct {
	//	XX map[int]string
	//}

	//11 结构体相等性
	//结构体内部所有字段都可以比较，结构体才可以比较，如果内部有不可比较的字段，结构体就不能比较
	//结构体内部所有字段都可以比较，结构体才可以比较
	//a:=Name{"randy","sun"}
	//	//b:=Name{"ra","s"}
	//	//if a==b{
	//	//	fmt.Printf("xxx")
	//	//}


	//如果内部有不可比较的字段，结构体就不能比较
	//c:=AA{}
	//d:=AA{}
	//if c==d {
	//
	//}


	//int 和int比较  string和string   数组和数组比      引用类型不能比较：切片不能和切片比  map不能和map比
	//	//int和string能比么？不是同一个类型就不能比较
	//	//Myint 和int


}

func stest5(p Person) {
	p.name = "randysun"
	fmt.Println(p)

}

func stest6(p *Person) () {
	//	//p是个指针
	//	//(*p).Name="randysun"  //正统的用法，需要解引用
	p.name = "randysun" //这种用法也可以，推荐这种用法，没有解引用，go自动处理
	fmt.Println(p)

}
