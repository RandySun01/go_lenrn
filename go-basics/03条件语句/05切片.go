package main

import (
	"fmt"
)

func main() {
	//var a [4]int = [4]int{1,2,3}
	//fmt.Println(a[:3])
	//b  = a[:]
	//print(b)
	//b = append(b, 1)
	//print(b)
	// 1.声明切片 切片是最最常用的数据结构，数组稍微少一些

	//var sc []string
	//var b = []int{}
	//fmt.Println(sc == nil)
	//
	//fmt.Println(b == nil) // 因为初始化了 0


	//var a [9]int = [9]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	//
	//fmt.Println(a, len(a)) //  2 3 4 5 6 7 8] 9
	//var b []int = a[:]
	////b是切片[]int --- int类型的切片
	//fmt.Println(b[1]) // 1
	//var c []int = a[:3]
	//fmt.Println(c) // [0 1 2]
	//
	//var d []int = a[3:6]
	//fmt.Println(d) // [3 4 5]
	//
	//var e []int = a[3:]
	//fmt.Println(e) // [3 4 5 6 7 8]
	// 2改变数据，会不会影响切片，数组或者切片发生改变，都会受到影响
	//b[0] = 100

	//fmt.Println(b) // [100 1 2 3 4 5 6 7 8]
	//fmt.Println(a) // [100 1 2 3 4 5 6 7 8]

	// 3. 切片的长度和容量 len,cap， 容量表示最多存放多少个值
	/*	var a[9]int = [9]int{0,1,2,3,4,5,6,7,8}
		fmt.Println(a) // [0 1 2 3 4 5 6 7 8]

		var b[]int = a[3:7]
		fmt.Println(b) // [3 4 5 6]
		// 长度
		fmt.Println(len(b)) // 4
		// cap 容量: 数组的长度减去切片开始的位置，就是存放数据的容量
		fmt.Println(cap(b)) // 6
	*/

	// 4、改变切片的其他操作，追加值，内置函数append
	/*
		a := [9]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
		fmt.Println(a) // [0 1 2 3 4 5 6 7 8]
		b := a[1:3]
		fmt.Println(b) // [1 2]
		// 第一个参数要添加值得切片，第二参数要添加的值
		b = append(b, 99)
		fmt.Println(b) // [1 2 99]
		b = append(b, 88)
		fmt.Println(b)              // [1 2 99,88]
		fmt.Println(len(b), cap(b)) // 4 8
	*/

	// 5 切片追到到了最后，超过了数组长度
	/*a := [9]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(a) // [0 1 2 3 4 5 6 7 8]
	b := a[6:8]
	fmt.Println(len(b), cap(b)) // 2 3
	fmt.Println(b)              // [6 7]

	// 追加值
	b = append(b, 9)
	fmt.Println(b)              // [6 7 9]
	fmt.Println(len(b), cap(b)) // 3 3

	// 容量中的值已满，再次追加值
	//如果再追加一个会怎么样？它会重写定义一个数组，把原来的值copy到新的数组上，切片是对新数组的引用
	////新数组大小是多少？如何看新数组大小（切片的容量）：原切片容量的两倍
	b = append(b, 10)
	fmt.Println(b) // 6 7 9 10]
	fmt.Println(len(b), cap(b))
	// 在此追加
	b = append(b, 11)
	fmt.Println(b) // [6 7 9]
	fmt.Println(len(b), cap(b))

	//原数组改变，会不会影响现在的切片
	//切片会变吗？不会变了，因为不引用原来的数组了
	a[1] = 100
	fmt.Println(a) // 	[0 100 2 3 4 5 6 7 9]
	fmt.Println(b) // [6 7 9 10 11]*/

	//现在新的数组(切片引用的数组)，可不可以打印出来？
	//不能赋值给一个变量，看不到了，在底层藏着

	// 数组定义就必须指定长度，数组不能变长（一开始就定了，定了就不能改了）
	//这种定义方式，了解

	/*var a [4]int = [...]int{1, 2, 3, 4}
	fmt.Println(a)      // [1 2 3 4]
	fmt.Println(len(a)) // 4
	var b []int
	fmt.Println(b)      // []
	fmt.Println(len(b)) //0*/

	// 空切片，只定义。没有赋值
	//只定义，没有赋值，空值是什么？空切片：nil
	//var b []int
	//fmt.Println(b)
	//if b == nil {
	//	fmt.Println("我是空的")
	//}
	//b[0] = 100 //空指针异常
	////nil[0]=100   //空指针异常
	//fmt.Println(b)

	// 6 切片的另一种定义方式(中括号内加东西，都是数组,不加东西，是切片)
	// make(类型，长度，容量),底层也引用了数组，只是看不到
	/*var a = make([]int, 3, 6)
	fmt.Println(a)              // [0 0 0]
	fmt.Println(len(a), cap(a)) // 	3 6
	// 传一个值表示，长度是3，容量也是3
	var b []int = make([]int, 3)
	fmt.Println(b)
	fmt.Println(len(b), cap(b))

	//第0个位置放1
	b[0] = 1
	//追加上1(会不会创建新数组？) 不会，只会切片进行扩容
	b = append(b, 1)
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(cap(b))*/
	// 7 补充（了解，切片的数据结构）不论切片中存长度有多大，切片这个变量占得内存空间都是一样大的
	/*
		{
			point:地址
			length：切片长度
			cap：切片容量
		}
	*/

	// 8 切片的函数传递(因为切片是引用类型，copy传递，改动原切片)
	/*var b []int=make([]int,4,6)
	fmt.Println(b)
	test1(b)
	fmt.Println(b)
	[0 0 0 0]
	[0 0 0 0]
	[0 0 0 0]

	*/

	//9 切片定义的第三种方式(定义并初始化，类似于数组)
	/*var b []int = []int{1, 2, 3}
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(cap(b))*/

	//10 多维切片（用的很少,了解）
	/*var b=[][]string{{"a","b"},{"c","d"}}
	var c [][]string=[][]string{{"a","b"},{"c","d"}}
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(cap(b))
	fmt.Println(c)
	fmt.Println(len(c))
	fmt.Println(cap(c))
	//也是2
	fmt.Println(cap(b[0]))*/
	// 11 切片的copy

	//var a []int = make([]int, 3, 10000)
	////var b []int=make([]int,3,4)
	////var b []int=make([]int,2,4)  //如果b的长度是2 会怎么样,
	//var b []int = make([]int, 4, 4) //如果b的长度是4 会怎么样
	//a[0] = 777
	//a[1] = 888
	//a[2] = 999
	//fmt.Println(a)
	//fmt.Println(b)
	////fmt.Println(a[3])  //错误的
	////把a的值copy到b中
	//copy(b, a)     // 如果a长度大于b， b只会保留b的容量的值，如果小于b则是b中的默认值
	//fmt.Println(b) //切片一定依附于数组，可以无限扩容（）
	//c := [...]string{"北京", "上海", "广州", "深圳", "成都", "重庆"}
	//fmt.Printf("a:%v type:%T len:%d  cap:%d\n", c, c, len(c), cap(c))
	//s := []int{2,3,5,7,11,13}
	//print(s[0])
	//s = s[:0]
	//print(s)
	//s = s[:4]
	//print(s)

	//a := []int{55, 56, 57, 58, 59}
	//a = a[1:4]                     //基于数组a创建切片，包括元素a[1],a[2],a[3]
	//fmt.Println(a)                  //[56 57 58]
	//fmt.Printf("type of a:%T\n", a) //type of b:[]int
	//var a1 = make([]string, 5, 10)
	//for i := 0; i < 10; i++ {
	//	a1 = append(a1, fmt.Sprintf("%v", i))
	//}
	//fmt.Println(a1)
	//fmt.Println(len(a1))
	//sort.IntsAreSorted(a[:])
	//fmt.Println(a)
	var a1 []string
	a1 = append(a1, "6")
	fmt.Println(a1)
}

func test1(a []int) {
	fmt.Println(a)

}
