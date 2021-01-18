package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfo struct {
	Id     int
	Name   string
	Gender string
	Hobby  string
}

func main() {
	db, err := gorm.Open("mysql", "root:toorex@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("gorm.Open err:", err)
		return
	}
	defer db.Close()

	// 创建表 自动迁移（把结构体和数据表进行对应）
	db.AutoMigrate(&UserInfo{})

	// 创建记录行
	u1 := UserInfo{1, "七米", "男", "蛙泳"}

	db.Create(&u1)

	// 查询
	var u UserInfo
	db.First(&u)

	fmt.Printf("u:%#v\n", u)

	db.Model(&u).Update("hobby", "双色球")

	db.Delete(u)
}
