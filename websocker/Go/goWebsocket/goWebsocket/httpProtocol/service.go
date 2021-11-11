package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 单独写回调函数
	http.HandleFunc("/go", myHandler)
	//http.HandleFunc("/ungo",myHandler2 )
	// addr：监听的地址  handler：回调函数
	http.ListenAndServe("127.0.0.1:8000",nil)
}

// handler函数
func myHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println(r.RemoteAddr, "连接成功")
	// 请求方式：GET POST DELETE PUT UPDATE
	fmt.Println("method: ", r.Method)
	fmt.Println("URL: ",r.URL.Path)
	fmt.Println("header: ", r.Header)
	fmt.Println("body: ", r.Body)
	//响应
	w.Write([]byte("randySun.com"))
}