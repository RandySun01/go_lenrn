package main

import "fmt"

func gotoDemo2() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	return
	// 标签
breakTag:
	fmt.Println("结束for循环")
}
func main() {
	//第一种
	//var a=1
	/* a :=12
	switch a {
	case 1:
		fmt.Println(1)
	case 4:
		fmt.Println(9)
	default:
		fmt.Println("条件都不满足执行我")
	}
	// 第二种 default
	//var a=1
	/* a :=12
	switch a {
	case 1:
		fmt.Println(1)
	case 4:
		fmt.Println(9)

	}
	*/
	// 3 多表达式实现， 只要符合其中的一个就会执行响应代码

	//a := 10
	//switch a {
	//case 1, 2, 3:
	//	fmt.Println("1-3")
	//case 4, 5, 6:
	//	fmt.Println("4-6")
	//case 7, 8, 9:
	//	fmt.Println("7-9")
	//
	//default:
	//	fmt.Println("其他种情况")
	//
	//}
	//
	//s := "a"
	//switch {
	//case s == "a":
	//	fmt.Println("a")
	//	fallthrough
	//case s == "b":
	//	fmt.Println("b")
	//case s == "c":
	//	fmt.Println("c")
	//default:
	//	fmt.Println("...")
	//}
	gotoDemo2()

	//switch b := 10; b  {
	//case 10 :
	//	fmt.Println("10-",b)
	//	fallthrough
	//
	//case 9:
	//	fmt.Println("9-",b)
	//case 8:
	//	fmt.Println(b)
	//
	//default:
	//	fmt.Println(b)
	//



	//}


	switch num:=30; { // num is not a constant
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is lesser than 200", num)

	}
}
