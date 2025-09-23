package main

import (
	"github.com/gin-gonic/gin"
	"github.com/learn/go_task/learn/task_4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:admin@123456@tcp(192.168.3.59:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	router := gin.Default()
	router.GET("/hello", func(c *gin.Context) {
		task_4.Register(c, db)
	})

	err2 := router.Run(":8080")
	if err2 != nil {
		panic(err2)
	}

}
