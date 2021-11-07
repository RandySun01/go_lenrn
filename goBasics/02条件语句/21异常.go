package main

import "fmt"

//  异常处理 不是try except
//特殊  defer panic  recover   结合起来用
// defer 延迟调用，即便程序出了严重错误，也会执行
// panic 主动抛出异常 等同于py中的raise
//recover 恢复程序，继续执行（程序出了异常，它给我恢复）

// 1defer 延迟调用，即便程序出了严重错误，也会执行
/*func main() {
	defer fmt.Println("xxx")  //延迟调用,先注册，后调用

	defer fmt.Println("ddd")  //延迟调用

	fmt.Println("yyy")
}*/

// panic 主动抛出异常 等同于py中的raise
/*func main() {
	defer fmt.Println("lqz lqz lqz ") //加了defer 表示程序执行完成后再调用defer注册的内容执行

	fmt.Println("xxx")
	fmt.Println("yyy")
	panic("我错了") //抛出异常
	fmt.Println("zzz")
}*/


func main() {
	defer func() {
		if error := recover(); error != nil {
			fmt.Println("error",error)
		}
		fmt.Println("这是finally的内容，不管程序是否出错，都会执行")
	}()

	f1()
	f2()
	f3()

}

func f1() {
	fmt.Println("f1  f1")
}
func f2() {
	//f2有异常，捕获异常
	//py 这么写
	//print(f1 f1)
	//try：
	//	fmt.Println("f2  f2")
	//	raise("xxx")
	//	print(pppp)
	//excepet Exception as e：
	//	print(e)
	//
	//finally:
	//	执行代码
	//
	//print(f3 f3)

	//defer func() {
	//	if error := recover(); error != nil {
	//		fmt.Println(error)
	//	}
	//	fmt.Println("这是finally的内容，不管程序是否出错，都会执行")
	//}()
	fmt.Println("f2  f2")
	panic("xxx")
	fmt.Println("f2asdfasdfadsfdsa  f2")  //会不会打印

}
func f3() {
	fmt.Println("f3  f3")
}

//以后的异常处理就是 拿到你要捕获异常的前面
//defer func() {
//	if error := recover(); error != nil {
//		fmt.Println(error)  //如果有异常就执行
//	}
//	fmt.Println("这是finally的内容，不管程序是否出错，都会执行") //finally执行的东西
//}()

