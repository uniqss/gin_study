package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name   string
	Age    int64
	Active bool
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

	//u1 := User{Name: "qimi", Age: 18, Active: true}
	//db.Create(&u1)
	//u2 := User{Name: "jinzhu", Age: 20, Active: false}
	//db.Create(&u2)

	var user User
	db.First(&user)

	user.Name = "七米"
	user.Age = 99
	//db.Debug().Save(&user)
	//db.Debug().Model(&user).Update("name", "小王子")

	//m1 := map[string]interface{}{
	//	"name":   "张小飞",
	//	"age":    28,
	//	"active": true,
	//}
	//db.Debug().Model(&user).Updates(m1)                // 更新所有
	//db.Debug().Model(&user).Select("age").Updates(m1)  // 只更新age
	//db.Debug().Model(&user).Omit("active").Updates(m1) // 排除m1中的active更新其他的字段

	//db.Debug().Model(&user).UpdateColumn("age", 30)
	//rowsNum := db.Model(User{}).Updates(User{Name: "张小飞"}).RowsAffected
	//fmt.Println(rowsNum)

	// 让users表中所有的用户年龄在原来基础上+2
	// 这个不建议搞，过于复杂，建议直接操作数据库写sql语句 update user set age = age + 2;

	// lesson24 delete
	//var u = User{}
	//u.ID = 1
	//db.Debug().Delete(&u)

	//var u = User{}
	//u.Name = "qimi"
	//db.Debug().Delete(&u) // 这样会全删

	//db.Where("name = ?", "qimi").Delete(User{})

	var u1 []User
	db.Debug().Unscoped().Where("name=?", "jinzhu2").Find(&u1)
	fmt.Println(u1)

	db.Debug().Where("name=?", "jinzhu2").Find(&u1)
	fmt.Println(u1)
}
