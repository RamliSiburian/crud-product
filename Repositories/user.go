package Repositories

import (
	"nutecth/Models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUser(ID int) (Models.User, error)
	Login(username string) (Models.User, error)
}

type users struct {
	db *gorm.DB
}

func RepositoryUser(db *gorm.DB) *users {
	return &users{db}
}

func (r *users) GetUser(ID int) (Models.User, error) {
	var user Models.User
	err := r.db.First(&user, ID).Error
	return user, err
}
func (r *users) Login(username string) (Models.User, error) {
	var user Models.User
	err := r.db.First(&user, "username=?", username).Error
	return user, err
}
