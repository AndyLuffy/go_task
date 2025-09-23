package main

import (
	"github.com/learn/go_task/learn/task_3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:admin@123456@tcp(192.168.3.59:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	task_3.Insert(db)
}
