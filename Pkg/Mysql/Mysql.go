package Mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error

	DBName := "root:@tcp(127.0.0.1:3306)/test_nutecth?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(DBName), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}
