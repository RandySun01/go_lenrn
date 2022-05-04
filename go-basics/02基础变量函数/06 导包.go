package main



/*

/*1 你写的所有包必须在gopath的src路径下
	包名一般是文件夹的名字
	在外部包可以直接调用， 包名.方法名（）
  2 包下的函数，如果大写字母开头，表示导出（可以给外部包使用），如果小写表示，只能给包内部使用
*/

//深入了解：你为什么配置goroot和gopath：如果不配goroot，内置包找不到，如果不配gopath，自己写的包找不到


import "daobao"
import "fmt"
func main() {
	a := daobao.Add(1)
	fmt.Println(a)

}
