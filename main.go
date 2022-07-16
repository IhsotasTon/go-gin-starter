package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 测试一个get请求
	r.GET("/rest/n/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint(name),
		})
	})
	// Run http server
	if err := r.Run(":8052"); err != nil {
		fmt.Printf("could notrun server: %v", err)
	}
}
