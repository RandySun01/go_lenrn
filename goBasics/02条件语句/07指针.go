package main

import "fmt"

// 存储变量内存地址的   变量
//什么类型的指针，就是什么类型前面加 *
//取一个变量的内存地址，用 &  取地址符号
func main() {
	//指针  ，c，c++有指针，go有指针，但是不完全一样（go的指针不支持运算） ，py和java，php都没有指针
	//存储变量内存地址的   变量

	// *  定义指针变量时，和解引用时使用
	// &  只用在取地址的时候
	//1 指针定义和初始化
	//a := 10
	////定义一个指针，存储a的内存地址
	////什么类型的指针，就是什么类型前面加 *
	////取一个变量的内存地址，用 &
	////指针类型的空值是  nil  -----> 引用类型的空值都是nil类型

	//var b *int = &a
	//fmt.Println(b)
	//fmt.Println(*b)
	//fmt.Println(a)

	// 2 解引用（反解）
	//a:=10
	//b:=&a
	//fmt.Println(b)
	//////b这个指针对应的值打印出来
	//fmt.Println(*b)

	//3 骚操作(只要是变量就会有地址，只要有地址，就可以取地址 &)
	//a := 10
	//b := &a
	//c := &b
	//d := &c
	//	//var b *int = &a
	//	//var c **int = &b
	//	//var d ***int = &c
	//fmt.Println(d)
	//fmt.Println(*d)
	//fmt.Println(**d)
	//fmt.Println(***d)

	//4 向函数传递指针参数
	// 通过函数传递，把a的值改为100
	//a := 10
	//b := &a
	////test4(a)
	////fmt.Println(a)
	//print(*b,1)
	//*b = 30
	//print(*b,2)

	//test5(b)
	//fmt.Println(b)
	//fmt.Println(a)

	//5 不要向函数传递数组的指针（取数组的地址），而要传递切片
	//通过函数传递，把原数组的值改变
	//var a =[5]int{9,8,7,6}
	//test6(&a) // 还有一个原因数组的大小不一致
	//fmt.Println(a)
	//test7(&a)  //这个不行 还有一个原因数组的大小不一致
	//test7(a[:]) // //这个不行 还有一个原因数组的大小不一致
	//fmt.Println(a)

	////6 go 不允许指针运算(c中，经常内存溢出)
	//var a=[5]int{9,8,7,6}
	//b:=&a
	//fmt.Println(b)
	//(*b)[1] = (*b)[1] + 1
	//fmt.Println((*b)[1] + 2)
	// c语言  b++指的是数组的第1个值（从0开始）
	// fmt.Println(b++)*/

	//d := map[int]string{1: "randy"}
	//var e *map[int]string = &d
	//fmt.Println((*e)[1])

	//7 数组指针和指针数组
	// 数组指针---》指向数组的指针
	//指针数组----》数组中放指针

	//var a = [5]int{9, 8, 7, 6}
	//var b *[5]int = &a //b是指向数组的指针
	//fmt.Println(b)
	//
	//x, y := 10, 11
	//var c [4]*int = [4]*int{&x, &y} //数组中放指针
	//fmt.Println(c)

	//var a *int
	//*a = 100
	//fmt.Println(*a)

	var b map[string]int
	fmt.Println(b)
	b = make(map[string]int)
	fmt.Println(b)
	b["沙河娜扎"] = 100
	fmt.Println(b)

	//var a *int
	//fmt.Println(a)
	////
	//a = new(int)
	//fmt.Println(a)
	//*a = 10
	//fmt.Println(*a)
	m := make(map[string]*student1)
	stus := []student1{
		{name: "randy", age: 18},
		{name: "bary", age: 23},
		{name: "jack", age: 9000},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
	}
	for k, v := range m {
		//fmt.Println(k, "=>", v.name)
		fmt.Printf("地址:%p, k:%v => v: %v \n", v, k, v.name)
	}
}

type student1 struct {
	name string
	age  int
}

func test4(a int) {
	a = 100
	fmt.Println(a)
}

func test5(a *int) {
	// 传递的是地址
	fmt.Println(a)
	*a = 100
	fmt.Println(*a)
}

func test6(a *[5]int) {
	(*a)[0] = 999
	fmt.Println(a) //理应该是个地址但是人家给你打印成了 &[999 8 7 6] 表示指向这个数组的指针
	fmt.Println(*a)
}
func test7(a []int) {
	a[0] = 999
	fmt.Println(a)
}
