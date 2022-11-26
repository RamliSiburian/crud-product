package Repositories

import (
	"nutecth/Models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindProduct() ([]Models.Product, error)
	GetProductById(ID int) (Models.Product, error)
	GetProductByUser(UserID int) ([]Models.Product, error)
	CreateProduct(product Models.Product) (Models.Product, error)
	DeleteProduct(product Models.Product) (Models.Product, error)
	UpdateProduct(product Models.Product) (Models.Product, error)
}

func RepositoryProduct(db *gorm.DB) *users {
	return &users{db}
}

func (r *users) FindProduct() ([]Models.Product, error) {
	var products []Models.Product
	err := r.db.Preload("User").Order("id desc").Find(&products).Error

	return products, err
}

func (r *users) GetProductByUser(UserID int) ([]Models.Product, error) {
	var product []Models.Product
	err := r.db.Where("user_id=?", UserID).Order("id desc").Find(&product).Error

	return product, err
}

func (r *users) GetProductById(ID int) (Models.Product, error) {
	var product Models.Product
	err := r.db.Preload("User").First(&product, ID).Error

	return product, err
}

func (r *users) CreateProduct(product Models.Product) (Models.Product, error) {
	err := r.db.Create(&product).Error

	return product, err
}

func (r *users) DeleteProduct(product Models.Product) (Models.Product, error) {
	err := r.db.Delete(&product).Error

	return product, err
}

func (r *users) UpdateProduct(product Models.Product) (Models.Product, error) {
	err := r.db.Save(&product).Error

	return product, err
}
