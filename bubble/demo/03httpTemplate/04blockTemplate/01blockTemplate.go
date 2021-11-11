package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, r *http.Request)  {
	// 定义模板
	// 解析模板
	t, err :=template.ParseFiles("./index.tmpl")
	if err != nil{
		fmt.Printf("parse tmpl failed err:%#v", err)
	}
	// 渲染模板
	msg :="这是index页面"
	t.Execute(w, msg)
}

func home(w http.ResponseWriter, r *http.Request)  {
	// 定义模板
	// 解析模板
	t, err :=template.ParseFiles("./home.tmpl")
	if err != nil{
		fmt.Printf("parse tmpl failed err:%#v", err)
	}
	// 渲染模板
	msg :="这是home页面"
	t.Execute(w, msg)
}

func blockIndex(w http.ResponseWriter, r *http.Request)  {
	// 定义模板
	// 解析模板，继承的模板在前 渲染模板在后
	t, err :=template.ParseFiles("./templates/base.tmpl", "./templates/blockIndex.tmpl")
	if err != nil{
		fmt.Printf("parse tmpl failed err:%#v", err)
	}
	// 渲染模板
	msg :="这是block index页面"
	// 指定渲染模板
	t.ExecuteTemplate(w, "blockIndex.tmpl", msg)
}

func blockHome(w http.ResponseWriter, r *http.Request)  {
	// 定义模板
	// 解析模板 继承的模板在前 渲染模板在后
	t, err :=template.ParseFiles("./templates/base.tmpl", "./templates/blockHome.tmpl")
	if err != nil{
		fmt.Printf("parse tmpl failed err:%#v", err)
	}
	// 渲染模板
	msg :="这是 block home页面"
	// 指定渲染模板
	t.ExecuteTemplate(w, "blockHome.tmpl", msg)
}
func main() {

	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/blockIndex", blockIndex)
	http.HandleFunc("/blockHome", blockHome)
	err := http.ListenAndServe(":9999", nil)
	if err != nil{
		fmt.Printf("create http service faild err:%#v", err)
		return
	}

}
