package repository

import (
	"mini-pos/database"
	"mini-pos/dto"

	"gorm.io/gorm"
)

type AuthorizeRepository interface {
	OwnerAuthorize(id uint, outlet_id uint) error
	StaffAuthorize(id uint, outlet_id uint) error
}

type authorizeRepo struct {
	DB *gorm.DB
}

func InitAuthorizeRepository() AuthorizeRepository {
	return &authorizeRepo{
		DB: database.DB,
	}
}

func (repo *authorizeRepo) OwnerAuthorize(id uint, outlet_id uint) error {

	outlet := dto.Outlet{}
	result := repo.DB.First(&outlet, "id = ? and created_by = ?", outlet_id, id)
	err := result.Error
	if err != nil {
		return err
	}
	if result.RowsAffected < 1 {
		return nil
	}

	return nil
}

func (repo *authorizeRepo) StaffAuthorize(id uint, outlet_id uint) error {

	user := dto.User{}
	user.Id = id
	user.OutletId = outlet_id
	err := repo.DB.Find(&user, user).Error

	if err != nil {
		return err
	}

	return nil
}
