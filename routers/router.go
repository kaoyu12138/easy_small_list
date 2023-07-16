package routers

import (
	"bubble/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*") //告诉gin框架这个index.html文件在哪里找
	//告诉gin框架那个html文件里需要的js文件在哪里找
	r.Static("static", "static")
	r.GET("/", controller.IndexHandler)
	//实现网页的渲染，但此时的网页还是一个网页，没有实现具体的功能

	v1Group := r.Group("v1")
	{
		v1Group.POST("/todo", controller.CreateTodo)

		v1Group.GET("/todo", controller.GetTodo)

		v1Group.PUT("/todo/:id", controller.UpdateTodo)

		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
	}

	return r
}
