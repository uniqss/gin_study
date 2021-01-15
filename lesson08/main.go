package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	t := template.New("index.tmpl").Delims("{[", "]}")
	var err error
	t, err = t.ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Println("parse template failed. err:", err)
		return
	}
	msg := "这是index页面"
	// 渲染模板
	err = t.Execute(w, msg)
	if err != nil {
		fmt.Println("Execute err:", err)
	}
}

func xss(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	t, err := template.New("xss.tmpl").Funcs(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	}).ParseFiles("./xss.tmpl")
	//t, err := template.ParseFiles("./xss.tmpl")
	if err != nil {
		fmt.Println("template.ParseFiles err:", err)
		return
	}
	// 渲染模板
	str1 := "<script>alert('asdf');</script>"
	str2 := "<a hre='www.baidu.com'>sss</a>"
	err = t.Execute(w, map[string]string{
		"str1": str1,
		"str2": str2,
	})
	if err != nil {
		fmt.Println("Execute err:", err)
	}
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/xss", xss)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("ListenAndServe err:", err)
		return
	}
}
