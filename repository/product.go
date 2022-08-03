package repository

import (
	"mini-pos/database"
	"mini-pos/dto"
	"time"

	"gorm.io/gorm"
)

type ProductRepository interface {
	List(page int, pageSize int) ([]dto.Product, error)
	Show(id uint) (dto.Product, error)
	Insert(dto.Product) (dto.Product, error)
	Update(dto.Product) (dto.Product, error)
	DeleteByID(ID uint) (data dto.Product, err error)
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

func (repo *productRepo) List(page int, pageSize int) (data []dto.Product, err error)  {
	if page == 0 {
		page = 1
	}

	switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
    }

	offset := (page - 1) * pageSize
	err = repo.DB.Offset(offset).Limit(pageSize).Find(&data).Error
	return data, err
}

func (repo *productRepo) Show(id uint) (data dto.Product, err error)  {
	err = repo.DB.First(&data, id).Error
	return data, err
}

func (repo *productRepo) Insert(payload dto.Product) (data dto.Product, err error) {
	err = repo.DB.Create(&payload).Error
	return payload, err
}

func (repo *productRepo) Update(payload dto.Product) (data dto.Product, err error) {

	// get book by id
	if err = repo.DB.First(&data, payload.Id).Error; err != nil {
		return
	}

	// update value
	data.CategoryId = payload.CategoryId
	data.Name = payload.Name
	data.Description = payload.Description
	data.ImagePath = payload.ImagePath
	data.IsActive = payload.IsActive
	data.UpdatedAt = time.Now()

	// update book data
	err = repo.DB.Save(&data).Error
	return
}

func (repo *productRepo) DeleteByID(ID uint) (data dto.Product, err error) {
	// get book by id
	if err = repo.DB.First(&data, ID).Error; err != nil {
		return
	}

	// delete book
	err = repo.DB.Delete(&data, ID).Error
	return
}

func (repo *productRepo) GetOutletProductByID(ID int) (data dto.OutletProduct, err error) {
	err = repo.DB.Preload("Product").Find(&data, ID).Error
	return
}
