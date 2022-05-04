package main

import "fmt"

func main() {

	//a:=find(12,12,34,45,56,67)
	//fmt.Println(a)
	//test333(1,2,3,4)
	//a := []int{1, 2, 3, 4}
	////test44(a)  //不可以
	//test44(a...) //相当于打散了,对数据类型有要求，也就是a必须是int类型切片才能打散了传过去

	type Empty interface {

	}
	var a []interface{}=[]interface{}{1,2,3}
	//var a []Empty=[]Empty{1,2,3}  //报错，不是同一种类型
	fmt.Println(a...)

	//a只能是空接口类型的切片
	//fmt.Println(a...)
}

func find(num int,nums ...int) bool {
	b:=false
	for _,v:=range nums{
		if num==v{
			fmt.Println("在里面")
			b=true
			break
		}else {
			b=false
		}
	}
	return b
}

func test333(a ...int) {
	//a 是个切片类型
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	//a是个数组？是个切片
	//查看变量a的类型（切片类型）
	fmt.Printf("%T", a)
	//fmt.Print()   //不换行
	//	//fmt.Println()  // 换行
	//	//fmt.Printf("%T",1)
}

////直接传切片
func test44(a ...int) {
	//a是切片，但是不能直接传切片
	//如果想直接传切片，可以吗？
	fmt.Println(a)

}


//package main
//
//import (
//	"fmt"
//)
//
//func change(s ...string) {
//	s[0] = "Go"
//	s = append(s, "playground")
//	fmt.Println(s)
//}
//
//func main() {
//	welcome := []string{"hello", "world"}
//	change(welcome...)
//	fmt.Println(welcome)
//}