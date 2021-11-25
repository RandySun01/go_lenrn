package routers

// 代办事项

// 添加
V1Group.POST("/todo", controller.CreateATodo)

// 查看所有代办事项
V1Group.GET("/todo", controller.GetTodoList)
// 查看摸一个代办事项
//v1Group.GET("/todo:id", func (c *gin.Context) {
//
//})
// 修改
V1Group.PUT("/todo/:id", controller.UpdateATodo)

// 删除
V1Group.DELETE("/todo/:id", controller.DeleteATodo)
