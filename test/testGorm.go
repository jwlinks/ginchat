package main

import (
	"fmt"
	"ginchat/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:sup157@tcp(localhost:23306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//迁移
	db.AutoMigrate(&models.UserBasic{})
	user := &models.UserBasic{}
	user.Name = "test"
	db.Create(user)
	fmt.Println(db.First(user, 1))

	db.Model(user).Update("PassWord", "123456")

}