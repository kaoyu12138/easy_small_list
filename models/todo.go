package models

import "bubble/dao"

//存放所有跟Todo数据模型相关的结构体定义，以及相关增删改查操作

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func CreateTodo(todo *Todo) (err error) {
	return dao.DB.Create(&todo).Error
}

func GetTodo(todolist *[]Todo) (err error) {
	return dao.DB.Find(&todolist).Error
}

func GetATodo(id string, todo *Todo) (err error) {
	return dao.DB.Where("id = ?", id).Find(&todo).Error
}

func UpdateTodo(todo *Todo) (err error) {
	return dao.DB.Save(&todo).Error
}

func DeleteTodo(id string) (err error) {
	return dao.DB.Where("id = ?", id).Delete(Todo{}).Error
}
