package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func indexHandler(c *gin.Context) {
	fmt.Println("indexHandler in ...")
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

func indexHandler2(c *gin.Context) {
	fmt.Println("indexHandler2 in ...")
	//c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

// 定义一个中间件m1
func m1(c *gin.Context) {
	fmt.Println("m1 in ...")

	// 计时
	start := time.Now()

	//go funcXX(c.Copy()) // 一旦使用goroutine一定要copy

	c.Next()  // 调用后续
	c.Abort() // 阻止后续的处理函数

	cost := time.Since(start)
	fmt.Println("cost:%v", cost)
	fmt.Println("m1 out ...")
}
func m2(c *gin.Context) {
	fmt.Println("m2 in ...")
	//c.Next()
	c.Abort()
	fmt.Println("m2 out ...")
}

func main() {
	//r := gin.Default() // logger recovery

	r := gin.New()

	//r.GET("/index", m1, indexHandler, indexHandler2)
	//r.GET("/shop", m1, func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{"msg": "/shop"})
	//})

	r.Use(m1, m2)
	r.GET("/index", indexHandler, indexHandler2)
	r.GET("/shop", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "/shop"})
	})

	r.Run(":9090")
}
