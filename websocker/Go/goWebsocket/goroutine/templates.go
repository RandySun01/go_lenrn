package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件生成模板对象

	tmpl, err := template.ParseFiles("E:\\Lenrn_notes\\03前端\\04websocker\\websocker\\Go\\goWebsocket\\goroutine\\hello.html")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	// 利用给定数据渲染模板，并将结果写入w
	tmpl.Execute(w, "5lmh.com")
}
func main() {
	//http.HandleFunc("/", sayHello)
	//err := http.ListenAndServe(":9090", nil)
	//if err != nil {
	//	fmt.Println("HTTP server failed,err:", err)
	//	return
	//}

	//resp, err := http.Get("https://liwenzhou.com.com/")
	//if err != nil {
	//	fmt.Println("get failed, err:", err)
	//	return
	//}
	//defer resp.Body.Close()
	//fmt.Println(resp.Body)
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println("read from resp.Body failed,err:", err)
	//	return
	//}
	//fmt.Print(string(body))
	//apiUrl := "http://127.0.0.1:9090/get"
	//// URL param
	//data := url.Values{}
	//data.Set("name", "枯藤")
	//data.Set("age", "18")
	//u, err := url.ParseRequestURI(apiUrl)
	//fmt.Println(u)
	//if err != nil {
	//	fmt.Printf("parse url requestUrl failed,err:%v\n", err)
	//}
	//u.RawQuery = data.Encode() // URL encode
	//fmt.Println(u.String())
	//resp, err := http.Get(u.String())
	//if err != nil {
	//	fmt.Println("post failed, err:%v\n", err)
	//	return
	//}
	//defer resp.Body.Close()
	//b, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println("get resp failed,err:%v\n", err)
	//	return
	//}
	//fmt.Println(string(b))

	//url := "http://127.0.0.1:9090/post"
	//// 表单数据
	////contentType := "application/x-www-form-urlencoded"
	////data := "name=枯藤&age=18"
	//// json
	//contentType := "application/json"
	//data := `{"name":"枯藤","age":18}`
	//resp, err := http.Post(url, contentType, strings.NewReader(data))
	//if err != nil {
	//	fmt.Println("post failed, err:%v\n", err)
	//	return
	//}
	//defer resp.Body.Close()
	//b, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println("get resp failed,err:%v\n", err)
	//	return
	//}
	//fmt.Println(string(b))

}
