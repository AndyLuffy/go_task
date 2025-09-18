package task_3

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
题目1：基本CRUD操作
假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
要求 ：
编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
*/
type Student struct {
	gorm.Model
	Name  string
	Age   int
	Grade string
}

func Insert() {

	/* sutdent:= Student{
		Name:  "张三",
		Age:   20,
		Grade: "三年级",
	}*/

	dsn := "root:admin@123456@tcp(192.168.3.59:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Student{})

}
