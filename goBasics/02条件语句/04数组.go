
package main

import "fmt"

func main() {
	// 统一类元素的集合， 不允许混合不同元素
	// 定义数据定义了一个大小为4的int类型数组
	//var s [4]string // 默认值为空
	//fmt.Println(s)


	//var a [4]int
	//fmt.Println(a) // int默认值为0
	//if s[0] == ""{
	//	fmt.Println("12")
	//}
	//
	//// 使用数组
	//a[0] = 10
	//fmt.Println(a[0])

	// 3定义并赋值三种
	//var a [4]int = [4]int {1,2,3,4}
	//fmt.Println(a)
	//
	//var b = [4]int {1,2,3,4}
	//fmt.Println(b)
	//
	//c := [4]int {1,2,34}
	//fmt.Println(c)
	////4  定义并赋初值（其他方式）
	//
	//var a [4]int = [4]int{1,2}
	//fmt.Println(a)
	//
	//var b [4]int = [4]int{2:5}
	//fmt.Println(b)
	//
	//var c [4]int = [4]int {3:4}
	//fmt.Println(c)
	//
	//var d [4]int =[4]int{3:4,2:3}
	//fmt.Println(d)

	//下面两个不是同一种类型(不能比较和加减)
	//var a [4]int
	//var b [4]int
	//if a > b{
	//
	//}

	/*a := [4]int{1,2,3}
		fmt.Println(len(a))
	// 遍历数组1
		for i:=0; i < len(a); i++{
			fmt.Println(a[i])
		}
		// 遍历数组2
		for i , v :=range a{ // （range:关键字，不是内置函数）把数组放在range关键字后
			fmt.Println(i, v)
		}

		//当用一个变量来接收，它就是索引
		//当用两个变量来结束，第一个是索引，第二个是真正的值
		//for i,v:=range a{
		for _,v:=range a{
			//fmt.Println(i)
			//fmt.Println("---------")
			fmt.Println(v)
		}*/

	// 多维数组
	//
	//var a [2][3]int
	//fmt.Println(a)
	//// 赋值
	//a[0][1] = 100
	//fmt.Println(a[0][1])

	// 定义并赋值
	//var b [2][3]int = [2][3]int{{1, 34}}
	//fmt.Println(b)
	//var c [2][3]int = [2][3]int{1: {1, 34, 5}}
	//fmt.Println(c)

	//字符串的空值是空字符串  ""
	//int类型空值是  0
	//数组的空值,取决于类型的空值
	//var d [4][3]string
	//fmt.Println(d)
	//var e = [4]int{1, 23}
	//arr(e)


	// 二维数组遍历

	//for _, b := range c {
	//	for _, c := range b {
	//		fmt.Println(c)
	//	}
	//}
	//var b string
	//fmt.Println(b,55555555555)

	//var a1 int       //int类型默认值为 0
	////var b string    //string类型默认值为 空字串
	////var c bool      //bool类型默认值为false
	////var d [2]int    //数组默认值为[0 0]
	//fmt.Println(&a1) //默认已经分配内存地址，所以我们可以直接使用&操作，取到的是一个合法的内存地址，而非nil

	var a =10   //定义变量a
	b := a      //将a的值赋值给b
	b = 101     //修改b的值，此时不会影响a
	fmt.Printf("值%v，内存地址%p\n",a,&a)   //a的值是10，a的内存地址是0xc42000e228
	fmt.Printf("值%v，内存地址%p\n",b,&b)   //b的值是101，b的内存地址是0xc42000e250

	var c =[3]int{1,2,3}    //定义一个长度为3的int类型的数组
	d := c      //将数组c赋值给d
	d[1] = 100  //修改数组d中索引为1的值为100
	fmt.Printf("值%v，内存地址%p\n",c,&c)   //c的值是[1 2 3]，c的内存地址是0xc42000a180
	fmt.Printf("值%v，内存地址%p\n",d,&d)   //d的值是[1 100 3]，d的内存地址是0xc42000a1a0
}

//func arr(a [4]int) {
//	for _, v := range a {
//		fmt.Println(v)
//	}
//
//}
