package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	t, err := template.ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Println("parse template failed. err:", err)
		return
	}
	// 解析模板
	msg := "这是index页面"
	// 渲染模板
	err = t.Execute(w, msg)
	if err != nil {
		fmt.Println("Execute err:", err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	t, err := template.ParseFiles("./home.tmpl")
	if err != nil {
		fmt.Println("parse template failed. err:", err)
		return
	}
	// 解析模板
	msg := "这是home页面"
	// 渲染模板
	err = t.Execute(w, msg)
	if err != nil {
		fmt.Println("Execute err:", err)
	}
}

func index2(w http.ResponseWriter, r *http.Request) {
	// 定义模板 模板继承的方式
	t, err := template.ParseFiles( "./templates/base.tmpl", "./templates/index2.tmpl")
	if err != nil {
		fmt.Println("parse template failed. err:", err)
		return
	}
	// 解析模板
	msg := "这是index2页面"
	// 渲染模板
	err = t.ExecuteTemplate(w, "index2.tmpl", msg)
	if err != nil {
		fmt.Println("Execute err:", err)
	}
}

func home2(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	t, err := template.ParseFiles( "./templates/base.tmpl", "./templates/home2.tmpl")
	if err != nil {
		fmt.Println("parse template failed. err:", err)
		return
	}
	// 解析模板
	msg := "这是home2页面"
	// 渲染模板
	err = t.ExecuteTemplate(w, "home2.tmpl", msg)
	if err != nil {
		fmt.Println("Execute err:", err)
	}
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/index2", index2)
	http.HandleFunc("/home2", home2)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("ListenAndServe err:", err)
		return
	}
}
