package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func helloTemplate(w http.ResponseWriter, r * http.Request)  {
	// 2.解析指定文件生产模板对象
	temp, err := template.ParseFiles("./helloTemplate.tmpl")
	if err != nil{
		fmt.Printf("create template failed err:%#v", err)
		return
	}

	// 利用给定数据渲染模板,并将结果写入到w
	temp.Execute(w, "hello template")


}
func main() {
	http.HandleFunc("/helloTemplate", helloTemplate)
	err := http.ListenAndServe(":9999",nil)
	if err != nil{
		fmt.Printf("http server run failed err:%#v", err)
	}

}
