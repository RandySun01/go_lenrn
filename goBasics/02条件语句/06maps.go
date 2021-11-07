package main

import "fmt"

func main() {
	//1 maps的定义           字典/hash
	//map[key的类型]value的类型
	//map的空值是什么？nil
	//var a map[int]string
	//fmt.Println(a)
	//if a == nil{
	//	fmt.Println("我是空的")
	//}

	//a[0] = "randysun"
	//fmt.Println(a[0]) // 报错，因为他是空的 */

	// 2.定义并初始化方式一、三种方式 make
	// make(类型)

	//var a map[int]string = make(map[int]string)
	//var b = make(map[int]string)
	//c := make(map[int]string)
	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Println(c)
	//if a == nil {
	//	fmt.Println("我是空的")
	//} else {
	//	fmt.Println("我不是空的")
	//}
	//
	//a[0] = "randy"
	//fmt.Println(a)
	//
	////3 定义并初始化方式二   直接赋值
	//
	//d := map[int]string{1: "randy", 0: "barry", 3:"wer"}
	//fmt.Println(d) // 打印结果是这种形式：map[0:barry 1:randy] 无序的

	// 4 maps的取值
	//	d := map[int]string{1: "randy", 0: "barry", 3: "wer"}
	//	fmt.Println(d[1])  // randy 打印结果是这种形式：map[0:barry 1:randy] 无序的
	//	fmt.Println(d[35]) // 为空
		//var b map[int]map[int][]int = map[int]map[int][]int{}   以后尽量不要这么玩


	//  5 maps取值如果key不存在，结果是value的空值。问题来了，因为取出来的可能为0  ""  nil。。。，如何判断呢？
	//	e := map[int]string{1: "randy", 0: "barry", 3: "wer"}
	//
	//	var a map[int]int= map[int]int{1:1,2:2}
	//	c:=a[9]
		//if c==""{
		//	fmt.Println("不存在")
		//}else c==0{
		//
		//}
		//神奇的使用,取值时，可以用两个变量来接收，第一个就是value值（可能为value值的空值），第二个是布尔，有就是true，没有就是false
		//c,ok:=e[9]
		//c1,ok1:=e[1]
		//fmt.Println(c)
		//fmt.Println(ok)
		//
		//fmt.Println(c1)
		//fmt.Println(ok1)

	// 6 给map添加元素
	//	var a map[int]string = map[int]string{1: "randy", 2: "barry"}
	//	//不存在的key（就是放进去）
	//	a[9] = "xxx"
	//	//存在的kye(替换)
	//	a[2] = "yyy"
	//	fmt.Println(a)

	// 7  删除map中元素
	/*var a map[int]string= map[int]string{1:"randy",2:"barry"}
	//内置函数delete
	delete(a, 2)
	fmt.Println(a)*/

	//8 map的长度 len()a
	//var a map[int]string= map[int]string{1:"randy",2:"barry"}
	//fmt.Println(len(a))
	// 的疑问
	//	//数字不需要写   写了没错，到底代表什么？底层,表示初始容量（不需要掌握，知道就可以了）
	//var a map[int]string= make(map[int]string, 3)
	//fmt.Println(cap(a)) //没有 报错
	//3 并不是map的长度
	//fmt.Println(len(a))
	//
	//a[1]="xxx"
	//a[2]="yy"
	//a[3]="zz"
	//a[4]="ii"
	//a[7]="ii"
	//fmt.Println(len(a))

	// 9 maps的迭代，循环9

	//a := map[int]string{1:"randy", 2:"barry"}
	//fmt.Println(a)
	//////for i:=0;i<len(a);i++{  //这个不可以
	//////	fmt.Println(a[i])
	//////}
	//////range 迭代
	//for k,v:=range a{
	//	fmt.Println(k)
	//	fmt.Println(v)
	//}
	// map中  go  map是无序的       py中从3.6以后字典有序，3.6之前无序，一旦做成有序，占得内存大了，


	//10 maps当参数传递,引用类型（地址），会改变原值
	/*	//var a map[int]string = map[int]string{1:"randy", 2:"barry"}
		a := map[int]string{1:"randy", 2:"barry"}
		fmt.Println(a)

		test3(a)
		fmt.Println(a)*/


	////11 相等性，不能直接比较
	//var a=map[int]string{1:"lqz"}
	////var b=map[int]string{1:"lqz"}
	////if a==b{
	////	//报错,不能比较
	////}
	//if a==nil{
	//fmt.Println("空的")
	//}else {
	//	fmt.Println("不是空的")
	//
	//}
	//var a map[int]string = map[int]string{1:"randy", 2:"barry"}
	//a := map[int]string{1:"randy", 2:"barry"}
	//fmt.Println(a)
	//fmt.Println("函数执行前", a)
	//test4(a)
	//fmt.Println("函数执行完", a)
	//var mapSlice = make([]map[string]string, 3)
	//for index, value := range mapSlice {
	//	fmt.Printf("index:%d value:%v\n", index, value)
	//}
	//fmt.Println("after init")
	//// 对切片中的map元素进行初始化
	//mapSlice[0] = make(map[string]string, 10)
	//mapSlice[0]["name"] = "小王子"
	//mapSlice[0]["password"] = "123456"
	//mapSlice[0]["address"] = "沙河"
	//for index, value := range mapSlice {
	//	fmt.Printf("index:%d value:%v\n", index, value)
	//}

	 var map1 map[int]string
	map1[1] = "4"
	fmt.Println(map1)


	m := make(map[int]string)
	m[1] = "4"
	fmt.Println(m)


}

func test4(a map[int]string)  {
	a[1] =  "老王"
	fmt.Println("函数内",a)


}
func test3(a map[int]string)  {
	a[1] =  "老王"
	fmt.Println(a)


}
