package controller

//将路由后的函数拆分出来

import (
	"bubble/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateTodo(c *gin.Context) {
	//取到用户在网页输入的数据
	var todo models.Todo
	c.BindJSON(&todo)

	//将取到的数据存到数据库中
	if models.CreateTodo(&todo) != nil {
		c.JSON(http.StatusOK, gin.H{"error": models.CreateTodo(&todo).Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func GetTodo(c *gin.Context) {
	var todolist []models.Todo

	if models.GetTodo(&todolist) != nil {
		c.JSON(http.StatusOK, gin.H{"error": models.GetTodo(&todolist).Error()})
	} else {
		c.JSON(http.StatusOK, todolist)
	}
}

func UpdateTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "id not useful"})
	}
	var todo models.Todo

	if models.GetATodo(id, &todo) != nil {
		c.JSON(http.StatusOK, gin.H{"error": models.GetATodo(id, &todo).Error()})
		return
	}

	c.BindJSON(&todo)

	if models.UpdateTodo(&todo) != nil {
		c.JSON(http.StatusOK, gin.H{"error": models.UpdateTodo(&todo).Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "id not useful"})
	}

	if models.DeleteTodo(id) != nil {
		c.JSON(http.StatusOK, gin.H{"error": models.DeleteTodo(id).Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "id deleted"})
	}
}
