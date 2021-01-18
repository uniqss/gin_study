package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type User struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64 `gorm:"column:user_age"`
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`
	MemberNumber *string `gorm:"unique;not null"`
	Num          int     `gorm:"AUTO_INCREMENT"`
	Address      string  `gorm:"index:addr"`
	IgnoreMe     int     `gorm:"-"` // 忽略
}

type Animal struct {
	AnimalId int64 `gorm:"primary_key"`
	Name     string
	Age      int64
}

// 指定表名
func (Animal) TableName() string {
	return "qimi"
}

func main() {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "tbl_" + defaultTableName
	}

	db, err := gorm.Open("mysql", "root:toorex@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("gorm.Open err:", err)
		return
	}
	defer db.Close()
	db.SingularTable(true) // 禁用复数

	// 创建表 自动迁移（把结构体和数据表进行对应）
	db.AutoMigrate(&User{})

	db.AutoMigrate(&Animal{})

	// 这样写不推荐
	db.Table("xiaowangzi").CreateTable(&User{})

}
