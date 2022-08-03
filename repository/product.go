package repository

import (
	"mini-pos/database"
	"mini-pos/dto"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetOutletProductByID(ID int) (dto.OutletProduct, error)
}

type productRepo struct {
	DB *gorm.DB
}

func InitProductRepository() ProductRepository {
	return &productRepo{
		DB: database.DB,
	}
}

func (repo *productRepo) GetOutletProductByID(ID int) (data dto.OutletProduct, err error) {
	err = repo.DB.Preload("Product").Find(&data, ID).Error
	return
}
