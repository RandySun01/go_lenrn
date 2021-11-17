package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.HEAD("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "head",
		})

	})
	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "get",
		})

	})

	r.POST("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "post",
		})

	})

	r.PUT("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "put",
		})
	})
	r.PATCH("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "patch",
		})
	})
	r.DELETE("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "delete",
		})
	})

	//// Any: 接收所有请求方式
	r.Any("/any", func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodGet:
			c.JSON(http.StatusOK, gin.H{
				"method": http.MethodGet,
			})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{
				"method": http.MethodPost,
			})
		case http.MethodHead:
			c.JSON(http.StatusOK, gin.H{
				"method": http.MethodHead,
			})
		case http.MethodDelete:
			c.JSON(http.StatusOK, gin.H{
				"method": http.MethodDelete,
			})

		}
	})

	// 路由组
	bookRouteGroup := r.Group("/book")
	{
		bookRouteGroup.GET("/details", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "book",
			})
		})

		bookRouteGroup.POST("/add", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "add book",
			})
		})
		bookRouteGroup.DELETE("/move", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "delete book",
			})
		})
	}

	// 路由组嵌套

	shopRouteGroup := r.Group("/shop")
	{
		shopRouteGroup.GET("/details", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "shop",
			})
		})

		shopRouteGroup.POST("/add", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "add shop",
			})
		})

		// 路由嵌套
		shopOrderRouteGroup := shopRouteGroup.Group("shopOrder")
		shopOrderRouteGroup.GET("/details", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "shop order",
			})
		})

		shopOrderRouteGroup.POST("/add", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "add shop order",
			})
		})
		shopOrderRouteGroup.DELETE("/move", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "delete shop order",
			})
		})
	}

	// NoRoute:处理不存在的路由
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "访问资源不存在",
		})

	})
	r.Run(":9999")
}
