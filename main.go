package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 2.解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("template.ParseFiles err:", err)
		return
	}
	// 3.渲染模板
	name := "哈哈哈哈"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Println("t.Execute err:", err)
		return
	}

}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("ListenAndServe err:", err)
		return
	}
}
