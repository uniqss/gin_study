package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.LoadHTMLFiles("./login.html", "./index.html")

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	r.POST("/login", func(c *gin.Context) {
		// 获取form表单提交的数据
		//username := c.PostForm("username")
		//password := c.PostForm("password")

		// 这个比较好
		username := c.DefaultPostForm("username", "somebody")
		password := c.DefaultPostForm("password", "ppp")

		//// 这个方法貌似不行
		//username, ok := c.GetPostForm("username")
		//if !ok {
		//	username = "哈哈"
		//}
		//password, ok := c.GetPostForm("password")
		//if !ok {
		//	password = "呵呵"
		//}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"Name":     username,
			"Password": password,
		})
	})

	err := r.Run(":9090")
	if err != nil {
		fmt.Println("Run err:", err)
		return
	}
}
