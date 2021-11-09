package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
@author RandySun
@create 2021-11-08-22:42
*/

func sayHello(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "<H1>hello web</H1>")
	_, _ = fmt.Fprintln(w, "<H1 style='color: red'>hello web</H1>")

}

func sayHelloFile(w http.ResponseWriter, r *http.Request) {

	b, _ := ioutil.ReadFile("G:\\goproject\\go\\bubble\\demo\\01httpTest\\helloWeb.txt")
	a, _ := fmt.Fprintln(w, string(b))
	fmt.Println(a)

}
func main() {
	http.HandleFunc("/hello", sayHello)
	http.HandleFunc("/helloFile", sayHelloFile)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		fmt.Printf("http server failed err:%#v", err)
	}

}
