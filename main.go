package main

import (
	"bubble/dao"
	"bubble/models"
	"bubble/routers"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	err := dao.InitMysql()
	if err != nil {
		panic(err)
	}
	defer dao.Close()
	dao.DB.AutoMigrate(&models.Todo{})

	r := routers.SetupRouter()
	r.Run()
}
