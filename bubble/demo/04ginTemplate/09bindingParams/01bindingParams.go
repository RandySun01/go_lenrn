package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

func main() {

	r := gin.Default()
	// querystring方式
	r.GET("/user", func(c *gin.Context) {
		//username := c.Query("username")
		//password := c.Query("password")
		//user := userInfo{
		//	Username: username,
		//	Password: password,
		//}
		// 数据绑定
		var userBind UserInfo // 声明变量

		err := c.ShouldBind(&userBind)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", userBind)
		}

		c.JSON(http.StatusOK, userBind)

	})

	// form方式
	r.POST("/form", func(c *gin.Context) {

		// 数据绑定
		var userBind UserInfo // 声明变量

		err := c.ShouldBind(&userBind)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", userBind)
		}

		c.JSON(http.StatusOK, userBind)

	})

	// json方式
	r.POST("/json", func(c *gin.Context) {

		// 数据绑定
		var userBind UserInfo // 声明变量

		err := c.ShouldBind(&userBind)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", userBind)
		}

		c.JSON(http.StatusOK, userBind)

	})
	r.Run(":9999")
}
