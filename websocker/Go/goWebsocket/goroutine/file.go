package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// 新建文件
	//file, err := os.Create("./xxx.txt")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer file.Close()
	//for i := 0; i < 5; i++ {
	//	file.WriteString("ab\n")
	//	file.Write([]byte("cd\n"))
	//}

	// 打开文件
	//file, err := os.Open("./xxx.txt")
	//if err != nil {
	//	fmt.Println("open file err :", err)
	//	return
	//}
	//defer file.Close()
	//// 定义接收文件读取的字节数组
	//var buf [128]byte
	//var content []byte
	//for {
	//	n, err := file.Read(buf[:])
	//	fmt.Println(n, 4444444444444444)
	//	if err == io.EOF {
	//		// 读取结束
	//		break
	//	}
	//	if err != nil {
	//		fmt.Println("read file err ", err)
	//		return
	//	}
	//	content = append(content, buf[:n]...)
	//}
	//fmt.Println(string(content))

	re()
}
func wr() {
	// 参数2：打开模式，所有模式d都在上面
	// 参数3是权限控制
	// w写 r读 x执行   w  2   r  4   x  1
	file, err := os.OpenFile("./xxx.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return
	}
	defer file.Close()
	// 获取writer对象
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("hello\n")
	}
	// 刷新缓冲区，强制写出
	writer.Flush()
}

func re() {
	file, err := os.Open("./xxx.txt")
	if err != nil {
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {

		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}
		fmt.Println(string(line))
	}

}
