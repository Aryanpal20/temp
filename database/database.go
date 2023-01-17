package database

import (
	"fmt"

	entity "api/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// here we want to create Database
var Database *gorm.DB

// Username:Password@tcp(127.0.0.1:3306)/Database_Name
var urlDSN = "root:Java1234!@#$@tcp(127.0.0.1:3306)/detail?charset=utf8mb4&parseTime=True&loc=Local"
var err error

func DataMigration() {

	Database, err = gorm.Open(mysql.Open(urlDSN), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())

		panic("connection failed")
	}
	Database.AutoMigrate(entity.User{}, entity.Task{}, entity.Otp{})
}
