package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func main() {
	r := gin.Default()

	r.LoadHTMLFiles("./index.html")

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.MaxMultipartMemory = 20 << 10
	r.POST("/upload", func(c *gin.Context) {
		// 从请求中读取文件
		f1, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			// 将读取到的文件保存到本地(服务端本地)
			//dst := fmt.Sprintf("./%s", f1.Filename)
			dst := path.Join("./", f1.Filename)
			err = c.SaveUploadedFile(f1, dst)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"status": "OK",
				})
			}
		}
	})

	err := r.Run(":9090")
	if err != nil {
		fmt.Println("Run err:", err)
		return
	}
}
