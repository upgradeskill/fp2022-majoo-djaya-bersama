package repository

import (
	"mini-pos/database"
	"mini-pos/dto"
	"time"

	"gorm.io/gorm"
)

type OutletRepository interface {
	GetAll(dto.Outlet, dto.Pagination) ([]dto.Outlet, error)
	GetByID(uint) (dto.Outlet, error)
	Insert(dto.Outlet) (dto.Outlet, error)
	Update(dto.Outlet) (dto.Outlet, error)
	DeleteByID(ID uint) (data dto.Outlet, err error)
}

type outletRepo struct {
	DB *gorm.DB
}

func InitOutletRepository() OutletRepository {
	return &outletRepo{
		DB: database.DB,
	}
}

func (repo *outletRepo) GetAll(payload dto.Outlet, pagination dto.Pagination) (data []dto.Outlet, err error) {
	err = pagination.Apply(repo.DB).Find(&data, payload).Error
	return
}

func (repo *outletRepo) GetByID(id uint) (data dto.Outlet, err error) {
	err = repo.DB.Find(&data, id).Error
	return
}

func (repo *outletRepo) Insert(payload dto.Outlet) (data dto.Outlet, err error) {
	err = repo.DB.Create(&payload).Error
	return payload, err
}

func (repo *outletRepo) Update(payload dto.Outlet) (data dto.Outlet, err error) {
	if err = repo.DB.First(&data, payload.Id).Error; err != nil {
		return
	}

	data.Name = payload.Name
	data.Address = payload.Address
	data.IsActive = payload.IsActive
	data.UpdatedAt = time.Now()

	err = repo.DB.Save(&data).Error
	return
}

func (repo *outletRepo) DeleteByID(Id uint) (data dto.Outlet, err error) {
	if err = repo.DB.First(&data, Id).Error; err != nil {
		return
	}

	err = repo.DB.Delete(&data, Id).Error
	return
}
