package repository

import (
	"mini-pos/database"
	"mini-pos/dto"
	"time"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetAll(dto.Category, dto.Pagination) ([]dto.Category, error)
	GetByID(uint) (dto.Category, error)
	Insert(dto.Category) (dto.Category, error)
	Update(dto.Category) (dto.Category, error)
	DeleteByID(ID uint) (data dto.Category, err error)
}

type categoryRepo struct {
	DB *gorm.DB
}

func InitCategoryRepository() CategoryRepository {
	return &categoryRepo{
		DB: database.DB,
	}
}

func (repo *categoryRepo) GetAll(payload dto.Category, pagination dto.Pagination) (data []dto.Category, err error) {
	err = pagination.Apply(repo.DB).Find(&data, payload).Error
	return
}

func (repo *categoryRepo) GetByID(id uint) (data dto.Category, err error) {
	err = repo.DB.Find(&data, id).Error
	return
}

func (repo *categoryRepo) Insert(payload dto.Category) (data dto.Category, err error) {
	err = repo.DB.Create(&payload).Error
	return payload, err
}

func (repo *categoryRepo) Update(payload dto.Category) (data dto.Category, err error) {
	if err = repo.DB.First(&data, payload.Id).Error; err != nil {
		return
	}

	data.Name = payload.Name
	data.Description = payload.Description
	data.IsActive = payload.IsActive
	data.UpdatedAt = time.Now()

	err = repo.DB.Save(&data).Error
	return
}

func (repo *categoryRepo) DeleteByID(Id uint) (data dto.Category, err error) {
	if err = repo.DB.First(&data, Id).Error; err != nil {
		return
	}

	err = repo.DB.Delete(&data, Id).Error
	return
}
