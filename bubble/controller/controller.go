package controller

import (
	"bubble/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// controller只调用具体的逻辑
/*
	url      ----->  controller ---->  logic ----->   model
  请求进进来			  控制器           逻辑处理       模型层的增删改查
*/
// 首页
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)

}

// 创建单条
func CreateATodo(c *gin.Context) {
	// 前端页面填写代办事项,点击提交接收数据
	// 获取数据保存并响应前端
	var todo models.Todo
	c.BindJSON(&todo)
	err := models.CreateATodo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		// 返回成功响应
		c.JSON(http.StatusOK, todo)
	}
}

// 查询展示
func GetTodoList(c *gin.Context) {

	todoList, err := models.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func UpdateATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "传入参数错误",
		})
		return
	}

	todo, err := models.GetATodo(id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.BindJSON(&todo)

	if err = models.UpdateATodo(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func DeleteATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "参数错误",
		})
		return
	}

	if err := models.DeleteATodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	// 都是零值
	//fmt.Println(todo.Title, todo.ID, todo.Status)
	c.JSON(http.StatusOK, gin.H{
		id: "delete succ",
	})
}
