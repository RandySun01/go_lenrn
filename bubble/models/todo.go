package models

import (
	"bubble/dao"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// TODO curd 都在这里面

// CreateATodo 创建todo
func CreateATodo(todo *Todo) (err error) {
	if err := dao.DB.Debug().Create(&todo).Error; err != nil {
		return err
	}
	return
}

// GetAllTodo 查询
func GetAllTodo() (todoList []*Todo, err error) {
	if err := dao.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}

// GetATodo // 查询单条记录
func GetATodo(id string) (todo *Todo, err error) {

	// 需要new
	todo = new(Todo)
	if err := dao.DB.Debug().Where("id=?", id).Find(&todo).Error; err != nil {
		return nil, err
	}
	return

}

// UpdateATodo更新数据
func UpdateATodo(todo *Todo) (err error) {
	if err = dao.DB.Debug().Save(&todo).Error; err != nil {
		return err
	}
	return
}

// 删除
func DeleteATodo(id string) (err error) {
	if err = dao.DB.Debug().Where("id = ?", id).Delete(Todo{}).Error; err != nil {
		return err
	}
	return
}
