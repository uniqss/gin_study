package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	Gender string
	Age    int
}

func f1(w http.ResponseWriter, r *http.Request) {
	fmt.Println("f1")

	kua := func(name string) (string, error) {
		return name + "真帅", nil
	}

	// 定义模板
	t := template.New("hello.tmpl")

	t.Funcs(template.FuncMap{
		"kua": kua,
	})

	//x, err := template.ParseFiles("./hello.tmpl")
	x, err := t.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("parse template failed. err:", err)
		return
	}

	name := "倪达耶"
	// 解析模板
	// 渲染模板
	err = x.Execute(w, name)
	if err != nil {
		fmt.Println("Execute err:", err)
	}
}

func Demo1(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	t, err := template.ParseFiles("./t.tmpl", "./ul.tmpl")
	if err != nil {
		fmt.Println("parse template failed. err:", err)
		return
	}
	// 解析模板
	name := "哈喽"
	// 渲染模板
	err = t.Execute(w, name)
	if err != nil {
		fmt.Println("Execute err:", err)
	}
}

func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/tmplDemo", Demo1)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("ListenAndServe err:", err)
		return
	}
}
