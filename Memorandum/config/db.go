package config

import (
	"Memorandum/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB(dsn string) {
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("err = ", err)
		return
	}

	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.Task{})
}
