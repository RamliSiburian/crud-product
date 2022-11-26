package Database

import (
	"fmt"
	"nutecth/Models"
	"nutecth/Pkg/Mysql"
)

func Migration() {
	err := Mysql.DB.AutoMigrate(&Models.User{}, Models.Product{})

	if err != nil {
		panic("Migration failed")
	}
	fmt.Println("Migration success")
}
