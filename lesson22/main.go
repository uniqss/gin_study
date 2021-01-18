package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	Id   int64
	Name string
	Age  int64
}

func main() {

	db, err := gorm.Open("mysql", "root:toorex@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("gorm.Open err:", err)
		return
	}
	defer db.Close()

	// 创建表 自动迁移（把结构体和数据表进行对应）
	db.AutoMigrate(&User{})
	//
	////var user User
	////user := new(User)
	//user := &User{}
	//
	//db.First(&user)
	//fmt.Printf("user:%#v\n", user)
	//
	//var users []User
	//db.Find(&users)
	//fmt.Printf("users:%#v\n", users)

	var user User
	db.Attrs(User{Age: 99}).FirstOrInit(&user, User{Name: "找不到"})
	fmt.Printf("user:%#v\n", user)

}
