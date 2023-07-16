package dao

//存放数据库相关的初始化等操作

import "github.com/jinzhu/gorm"

var (
	DB *gorm.DB
)

func InitMysql() (err error) {
	DB, err = gorm.Open("mysql", "root:root1234@tcp(127.0.0.1:13306)/todo?charset=utf8mb4&parseTime=True&loc=Local")
	return err
}

func Close() {
	DB.Close()
}
