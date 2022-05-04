package main

import "fmt"

func main() {
	/*
		for 初始化； 条件； 自增/自减
	*/

	// 第一种方式
	//for a := 0; a < 10 ; a++{
	//	fmt.Println(a)
	//}

	// 第二种方式
	/*	a := 0 // 作用于全局

		for ; a < 10; a++ {
			fmt.Println(a)
		}
		fmt.Println(" 第二种方式 全局作用局：", a)*/

	// 第三种方式
	/*b := 1
	for ; b < 10; {
		fmt.Println(" 第二种方式 全局作用局：", b)
		b++
	}

	fmt.Println("第三种方式： ", b)*/

	// 第四种方式 省略初始化和自增自减--进阶版，两个分号都省略（while）
	// for 条件(){	}

	/*a := 0
	for a < 10{
		fmt.Println("升级版while", a)
		a++
	}*/
	// 第五种方式 死循环
	// 死循环1
	//for true {
	//	fmt.Println("sss")
	//}

	// 死循环2
/*	for ; ; {
		fmt.Println("sss")
	}*/
	// 死循环3
	/*for  {
		fmt.Println("sss")
	}*/
	// break
	for  {
		fmt.Println("sss")
		break
	}
	// continue
	for  {
		fmt.Println("sss")
		continue
	}
}
