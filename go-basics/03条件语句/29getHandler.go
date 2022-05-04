package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
@author RandySun
@create 2022-02-15-8:49
*/
func getHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data := r.URL.Query()
	fmt.Println(data.Get("name"))
	fmt.Println(data.Get("age"))
	answer := `{"status": "ok"}`
	w.Write([]byte(answer))
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// 1. 请求类型是application/x-www-form-urlencoded时解析form数据
	r.ParseForm()
	fmt.Println(r.PostForm) // 打印form数据
	fmt.Println(r.PostForm.Get("name"), r.PostForm.Get("age"))
	// 2. 请求类型是application/json时从r.Body读取数据
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read request.Body failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
	answer := `{"status": "ok"}`
	w.Write([]byte(answer))
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "say Hello go！")
}

func main() {
	http.HandleFunc("/", sayHello)

	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/post", postHandler)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}