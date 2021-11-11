package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func customDefined(w http.ResponseWriter, r *http.Request) {
	// 定义模板

	// 解析模板
	// 自定默认的标识符
	t, err := template.New("customDefined.tmpl").
		Delims("{[", "]}").
		ParseFiles("./customDefined.tmpl")
	if err != nil {
		fmt.Printf("parse tmpl falied err:%#V", err)
	}
	// 渲染模板
	msg := "我是自定义标识符语法,快使用我吧!"
	err = t.Execute(w, msg)

	if err != nil {
		fmt.Printf("execute template failed  err: %#v", err)
		return
	}
}

func xss(w http.ResponseWriter, r *http.Request) {
	// 定义模板

	// 解析模板
	// 自定默认的标识符
	t, err := template.ParseFiles("./xss.tmpl")
	if err != nil {
		fmt.Printf("parse tmpl falied err:%#v\n", err)
	}
	// 渲染模板
	str := "<script>alert(123);</script>"
	//str := "<a href='www.baidu.com'>百度</a>"
	err = t.Execute(w, str)

	if err != nil {
		fmt.Printf("execute template failed  err: %#v", err)
		return
	}
}

func xssSafe(w http.ResponseWriter, r *http.Request) {
	// 定义模板

	// 解析模板
	// 自定默认的标识符
	t, err := template.New("xss.tmpl").Funcs(
		template.FuncMap{
			"safe": func(s string) template.HTML {
				return template.HTML(s)
			},
		},
	).ParseFiles("./xss.tmpl")
	if err != nil {
		fmt.Printf("parse tmpl falied err:%#v\n", err)
	}
	// 渲染模板
	str1 := "<script>alert(123);</script>"
	str2 := "<a href='https://www.cnblogs.com/randysun/'>博客</a>"
	err = t.Execute(w, map[string]string{
		"str1": str1,
		"str2": str2,
	})

	if err != nil {
		fmt.Printf("execute template failed  err: %#v", err)
		return
	}
}
func main() {
	http.HandleFunc("/customDefined", customDefined)
	http.HandleFunc("/xss", xss)
	http.HandleFunc("/xssSafe", xssSafe)

	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		fmt.Printf("http server run faild err:%#V", err)
	}

}
