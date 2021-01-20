package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"strconv"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func main() {
	// 连接数据库
	db, err := gorm.Open("mysql", "root:toorex@(127.0.0.1:3306)/bubble2?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("gorm.Open err:", err)
		return
	}
	defer db.Close()

	db.AutoMigrate(Todo{})

	r := gin.Default()

	r.Static("/static", "static")

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	v1Group := r.Group("/v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", func(c *gin.Context) {
			fmt.Println("添加")

			var todo Todo
			err := c.ShouldBind(&todo)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				fmt.Printf("%#v\n", todo)
				db.Create(&todo)
				c.JSON(http.StatusOK, gin.H{
					"status": "ok",
				})
			}
		})
		// 查看所有的待办事项
		v1Group.GET("/todo", func(c *gin.Context) {
			fmt.Println("查看所有的待办事项")
			var todos []Todo
			result := db.Find(&todos)
			if result.Error == nil {
				c.JSON(http.StatusOK, todos)
			} else {
				c.JSON(http.StatusInternalServerError, nil)
			}
		})
		// 查看某一个待办事项
		v1Group.GET("/todo/:id", func(c *gin.Context) {
			fmt.Println("查看某一个待办事项")

		})
		// 修改
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			idStr := c.Param("id")
			id, _ := strconv.Atoi(idStr)
			fmt.Println("修改")
			var todo Todo
			todo.ID = id
			// ??? update tb set status = !status where id = 1; 搞这么复杂。。。shit
			//db.Debug().Select(&todo)
			//db.Debug().Model(&todo).Update(map[string]interface{}{"status": !todo.Status})
			db.Debug().First(&todo, id).Update(map[string]interface{}{"status": !todo.Status})
		})
		// 删除
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			fmt.Println("删除")
			idStr := c.Param("id")
			id, _ := strconv.Atoi(idStr)
			var todo Todo
			todo.ID = id
			db.Delete(&todo)
		})
	}

	err = r.Run(":9090")
	if err != nil {
		fmt.Println("r.Run err:", err)
	}
}
