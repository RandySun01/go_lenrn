package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func index(c *gin.Context) {
	fmt.Println("index")
	name, _ := c.Get("name")

	fmt.Println("获取在中间中设置的值name:", name)
	c.JSON(http.StatusOK, gin.H{
		"msg": "index",
	})
}

// 定义一个中间件m1统计请求的耗时时间
func m1(c *gin.Context) {
	fmt.Println("m1 in ....")

	// 计算请求时间
	start := time.Now()

	c.Next() // 调用后续处理的函数
	//c.Abort() // 阻止调用后续处理的函数
	end := time.Since(start)
	fmt.Printf("消耗时间time:%v\n", end)
	fmt.Println("m1 out ....")

}

func m2(c *gin.Context) {
	fmt.Println("m2 in ....")
	// 在中间中设置值
	c.Set("name", "randySun")

	c.Next() // 调用后续处理的函数
	fmt.Println("m2 out ....")

}

// 认证中间件
//func authMiddleware(c *gin.Context) {
//	// 是否登录判断
//	username := c.Query("username")
//	// 判断是否登录用户
//	if username == "RandySun" {
//
//		c.Next()
//	} else {
//		// 认证失败
//		c.JSON(http.StatusUnauthorized, gin.H{
//			"msg": "没有权限",
//		})
//		c.Abort()
//	}
//}

// 通过闭包认证中间件
func authMiddleware(doCheck bool) gin.HandlerFunc {
	// 连接数据库
	// 或者其他准备工作
	fmt.Println(33333333)
	return func(c *gin.Context) {
		// 是否登录判断
		username := c.Query("username")
		// 判断是否登录用户
		if doCheck {
			if username == "RandySun" && doCheck {
				c.Next()
			} else {
				// 认证失败
				c.JSON(http.StatusUnauthorized, gin.H{
					"msg": "没有权限",
				})
				c.Abort()
			}

		} else {
			// 放行认证
			c.Next()
		}

	}
}

func main() {
	r := gin.Default()

	// 全局注册中间
	r.Use(m1, m2, authMiddleware(true))
	r.GET("/index", index)
	// 局部中间件
	r.GET("/home", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "home",
		})
	})
	r.GET("/add", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "add",
		})
	})

	shopGroup := r.Group("/shop")
	// 为路由组添加中间件
	shopGroup.Use(authMiddleware(true))
	{
		shopGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"mgs": "shop index",
			})
		})

	}

	r.Run(":9999")

}
