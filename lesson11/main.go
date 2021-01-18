package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/web", func(c *gin.Context) {
		//name := c.Query("query")
		//name := c.DefaultQuery("query", "somebody")
		//name, ok := c.GetQuery("query")
		//if !ok {
		//	name= "xxx"
		//}

		name := c.Query("query")
		age := c.Query("age")

		c.JSON(http.StatusOK, gin.H{
			"name":   name,
			"age":    age,
			"status": "ok",
		})
	})

	err := r.Run(":9090")
	if err != nil {
		fmt.Println("Run err:", err)
		return
	}
}
