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
func customFunc(w http.ResponseWriter, r *http.Request) {

	// 1. 编写模板

	// 2. 读取模板 解析指定文件生成模板对象

	// 自定义一个夸人的模板函数,要么返回两个值,第二个返回值必须是error类型
	kua := func(arg string) (string, error) {
		return arg + "-自定义函数", nil
	}
	 // 创建一个名字是customFunctionTemplate.tmpl模板对象,名字一定要与模板的名字对应上
	t, err := template.New("customFunctionTemplate.tmpl").Funcs(template.FuncMap{"kua": kua}).ParseFiles("./customFunctionTemplate.tmpl")
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

// 传入单对象
func nestingTmpl(w http.ResponseWriter, r *http.Request) {

	// 解析模板
	// 要将被包含的模板写在后面
	tmpl, err := template.ParseFiles("./nestingTemplate.tmpl", "./ul.tmpl")
	if err != nil {
		fmt.Println("create templates failed, err:", err)
		return
	}
	user := UserInfo{
		Name:   "RandySun",
		Gender: "男",
		Age:    18,
	}
	// 渲染模板
	tmpl.Execute(w, user)
}
func main() {

	http.HandleFunc("/customFunc", customFunc)
	http.HandleFunc("/nestingTmpl", nestingTmpl)

	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		fmt.Printf("http server run failed err:%#v", err)
	}
}
