package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type UserInfo struct {
	Name   string
	Age    int
	Gender string
}

// 传入单对象
func grammarFunc(w http.ResponseWriter, r *http.Request) {

	// 1. 编写模板

	// 2. 读取模板 解析指定文件生成模板对象
	t, err := template.ParseFiles("./customFunctionTemplate.tmpl")
	if err != nil {
		fmt.Printf("read templates faild err: %#v", err)
	}

	// 3. 渲染模板
	// 利用给定数据渲染模板，并将结果写入w
	userInfo := UserInfo{
		"RandySun",
		18,
		"男",
	}
	t.Execute(w, userInfo)
}

// 传入多对象
func grammarFuncMany(w http.ResponseWriter, r *http.Request) {
	// 1. 编写模板

	// 2. 读取模板 解析指定文件生成模板对象
	t, err := template.ParseFiles("./customFunctionTemplate.tmpl")
	if err != nil {
		fmt.Printf("read templates faild err: %#v", err)
	}

	// 3. 渲染模板
	// 传入map键值对
	m := map[string]interface{}{}
	m["name"] = "jack"
	m["age"] = 14
	m["gender"] = "男"
	//t.Execute(w, userInfo)
	t.Execute(w, m)
}

func grammarFuncManyToMany(w http.ResponseWriter, r *http.Request) {
	// 1. 编写模板

	// 2. 读取模板 解析指定文件生成模板对象
	t, err := template.ParseFiles("./customFunctionTemplate.tmpl")
	if err != nil {
		fmt.Printf("read templates faild err: %#v", err)
	}

	// 3. 渲染模板
	// 传入map键值对
	m := map[string]interface{}{}
	m["name"] = "jack"
	m["age"] = 14
	m["gender"] = "男"

	userInfo := UserInfo{
		"RandySun",
		18,
		"男",
	}
	um := map[string]interface{}{
		"m": m,
		"u": userInfo,
		"s": []int{1, 3, 4},
	}
	t.Execute(w, um)
}
func main() {

	http.HandleFunc("/grammar", grammarFunc)

	http.HandleFunc("/grammarMany", grammarFuncMany)
	http.HandleFunc("/grammarManyToMany", grammarFuncManyToMany)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		fmt.Printf("http server run failed err:%#v", err)
	}
}
