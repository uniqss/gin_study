package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

var (
	DB *gorm.DB
)

func initMysql() error {
	var err error

	DB, err = gorm.Open("mysql", "root:toorex@(127.0.0.1:3306)/bubble2?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("gorm.Open err:", err)
		return err
	}

	return DB.DB().Ping()
}

func main() {
	// 连接数据库
	//DB, err := gorm.Open("mysql", "root:toorex@(127.0.0.1:3306)/bubble2?charset=utf8mb4&parseTime=True&loc=Local")
	//if err != nil {
	//	fmt.Println("gorm.Open err:", err)
	//	return
	//}
	//defer DB.Close()

	err := initMysql()
	if err != nil {
		panic(err)
	}
	defer DB.Close()

	DB.AutoMigrate(Todo{})

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
			// 前面端面添加待办事项，点击提交，会发请求到这里
			fmt.Println("添加")

			//var todo Todo
			//err := c.ShouldBind(&todo)
			//if err != nil {
			//	c.JSON(http.StatusBadRequest, gin.H{
			//		"error": err.Error(),
			//	})
			//} else {
			//	fmt.Printf("%#v\n", todo)
			//	DB.Create(&todo)
			//	c.JSON(http.StatusOK, gin.H{
			//		"status": "ok",
			//	})
			//}

			// 1.从请求中把数据拿出来
			var todo Todo
			c.BindJSON(&todo)
			// 2.存入数据库
			err := DB.Create(&todo).Error
			// 3.返回响应
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, todo)
			}
			//
		})
		// 查看所有的待办事项
		v1Group.GET("/todo", func(c *gin.Context) {
			fmt.Println("查看所有的待办事项")
			//var todos []Todo
			//result := DB.Find(&todos)
			//if result.Error == nil {
			//	c.JSON(http.StatusOK, todos)
			//} else {
			//	c.JSON(http.StatusInternalServerError, nil)
			//}
			// 查询 todos 表中的所有数据
			var todoList []Todo
			result := DB.Find(&todoList)
			if result.Error != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": result.Error,
				})
			} else {
				c.JSON(http.StatusOK, todoList)
			}
		})
		// 查看某一个待办事项
		v1Group.GET("/todo/:id", func(c *gin.Context) {
			fmt.Println("查看某一个待办事项")
			// 其实没有用到
		})
		// 修改
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			fmt.Println("修改")
			//idStr := c.Param("id")
			//id, _ := strconv.Atoi(idStr)
			//var todo Todo
			//todo.ID = id
			//// ??? update tb set status = !status where id = 1; 搞这么复杂。。。shit
			////DB.Debug().Select(&todo)
			////DB.Debug().Model(&todo).Update(map[string]interface{}{"status": !todo.Status})
			//DB.Debug().First(&todo, id).Update(map[string]interface{}{"status": !todo.Status})
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "get param id error"})
				return
			}
			var todo Todo
			if err = DB.Where("id=?", id).First(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			}
			c.BindJSON(&todo)
			if err = DB.Save(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, todo)
		})
		// 删除
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			fmt.Println("删除")
			//idStr := c.Param("id")
			//id, _ := strconv.Atoi(idStr)
			//var todo Todo
			//todo.ID = id
			//DB.Delete(&todo)

			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "get param id error"})
				return
			}
			if err = DB.Where("id=?", id).Delete(Todo{}).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{id: "deleted"})
		})
	}

	err = r.Run(":9090")
	if err != nil {
		fmt.Println("r.Run err:", err)
	}
}
