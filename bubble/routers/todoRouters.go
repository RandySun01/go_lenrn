package routers

import (
	"bubble/controller"
	"github.com/gin-gonic/gin"
)

// todo路由
func TodoRouters(r *gin.Engine) {
	// 代办事项
	v1Group := r.Group("v1")

	// 添加
	v1Group.POST("/todo", controller.CreateATodo)

	// 查看所有代办事项
	v1Group.GET("/todo", controller.GetTodoList)
	// 查看摸一个代办事项
	v1Group.GET("/todo:id", func(c *gin.Context) {

	})
	// 修改
	v1Group.PUT("/todo/:id", controller.UpdateATodo)

	// 删除
	v1Group.DELETE("/todo/:id", controller.DeleteATodo)
}
