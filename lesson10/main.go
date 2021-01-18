package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/json", func(c *gin.Context) {
		// 方法1.1 使用map
		//data := map[string]interface{}{
		//	"name":    "小王子",
		//	"message": "hello world!",
		//	"age":     18,
		//}

		// 方法1.2
		data := gin.H{
			"name":    "小王子",
			"message": "world hello!",
			"age":     18,
		}

		c.JSON(http.StatusOK, data)
	})

	// 方法2
	type msg struct {
		Name    string `json:"name"`
		Message string
		Age     int
	}

	r.GET("/json2", func(c *gin.Context) {
		data := msg{
			"老王子",
			"hello haha",
			88,
		}
		c.JSON(http.StatusOK, data)
	})

	err := r.Run(":9090")
	if err != nil {
		fmt.Println("Run err:", err)
		return
	}
}
