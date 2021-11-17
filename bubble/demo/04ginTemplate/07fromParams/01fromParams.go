package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()
	r.LoadHTMLFiles("../templates/login.html")
	// 获取form表单提交数据
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)

	})
	// 获取数据
	r.POST("/login", func(c *gin.Context) {
		// 方式一
		//username := c.PostForm("username")
		//password := c.PostForm("password")

		// 方式二
		//username := c.DefaultPostForm("username", "body") // 取不到key,则值为默认值body
		//password := c.DefaultQuery("password", "123")

		// 方式三
		username, ok := c.GetPostForm("username")
		if !ok {
			username = "body"
		}

		password, ok := c.GetPostForm("password")
		if !ok {
			password = "sdfsdf"
		}
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})

	})
	r.Run(":9999")
}
