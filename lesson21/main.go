package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	Id int64
	//Name string `gorm:"default:'小王子'"`
	Name sql.NullString `gorm:"default:'小王子'"`
	//Age  int64
	Age sql.NullInt64
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

	u := User{Name: sql.NullString{String: "", Valid: true}, Age: sql.NullInt64{Int64: 100, Valid: true}}

	fmt.Println(db.NewRecord(&u)) // 判断主键是否为空
	db.Debug().Create(&u)
	fmt.Println(db.NewRecord(&u)) // 判断主键是否为空
}
