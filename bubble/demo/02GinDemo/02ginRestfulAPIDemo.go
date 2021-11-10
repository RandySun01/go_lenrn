package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// 获取书籍信息
	r.GET("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"messages": "GET",
		})
	})

	// 添加书记信息
	r.POST("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"messages": "POST",
		})
	})

	// 整理修改
	r.PUT("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"messages": "PUT",
		})
	})

	// 部分修改
	r.PATCH("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"messages": "PATCH",
		})

	})

	// 删除书籍
	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"messages": "DELETE",
		})

	})
	r.Run()
}
