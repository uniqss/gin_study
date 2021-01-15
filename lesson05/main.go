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

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 2.解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("template.ParseFiles err:", err)
		return
	}
	// 3.渲染模板
	u1 := User{
		Name:   "小王子",
		Gender: "男",
		Age:    18,
	}
	_ = u1

	m1 := map[string]interface{}{
		"Name":   "张三",
		"Gender": "男",
		"Age":    18,
	}
	_ = m1

	//err = t.Execute(w, u1)

	hobbyList := []string{
		"篮球",
		"足球",
		"乒乓球",
	}

	err = t.Execute(w, map[string]interface{}{
		"u1": u1,
		"m1": m1,
		"hobby": hobbyList,
	})

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
