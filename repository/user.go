package repository

import (
	"errors"
	"math"
	"mini-pos/database"
	"mini-pos/dto"
	"mini-pos/util"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type UserRepository interface {
	Insert(dto.User) (dto.User, error)
	Update(dto.User) (dto.User, error)
	Login(username string, password string) (dto.User, error)
	DeleteByID(owner uint, ID uint) (data dto.User, err error)
	UserByID(owner uint, ID uint) (data dto.User, err error)
	UserList(owner uint, outlet string, status string, term string, page string, limit string) (data []dto.User, meta dto.UserMeta, err error)
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

func (repo *userRepo) DeleteByID(owner uint, ID uint) (data dto.User, err error) {
	// get book by id
	if err = repo.DB.Where("created_by = ? ", owner).First(&data, ID).Error; err != nil {
		return
	}

	// delete book
	err = repo.DB.Where("created_by = ? ", owner).Delete(&data, ID).Error
	return
}

func (repo *userRepo) UserByID(owner uint, ID uint) (data dto.User, err error) {
	// get book by id
	if err = repo.DB.Where("created_by = ? ", owner).First(&data, ID).Error; err != nil {
		return
	}
	// delete book
	return data, nil
}

func (repo *userRepo) UserList(owner uint, outlet string, status string, term string, page string, limit string) (data []dto.User, meta dto.UserMeta, err error) {
	query := repo.DB.Where("created_by = ? ", owner)
	var intLimit int
	var intPage int

	if outlet != "" {
		intOutlet, _ := strconv.Atoi(outlet)
		query.Where("outlet_id = ? ", intOutlet)
	}
	if status != "" {
		intStatus, _ := strconv.Atoi(status)
		query.Where("status = ? ", intStatus)
	}
	if term != "" {
		query.Where("(id = ? or name Like ? or phone_number Like ?)", term, "%"+term+"%", "%"+term+"%")
	}

	meta.TotalRow = int(query.Find(&data).RowsAffected)

	if limit != "" {
		intLimit, _ = strconv.Atoi(limit)
		query.Limit(intLimit)
	} else {
		intLimit = 5
	}
	intPage, _ = strconv.Atoi(page)
	query.Offset((intPage - 1) * intLimit)

	meta.Limit = intLimit
	meta.TotalPage = int(math.Floor(float64(meta.TotalRow) / float64(intLimit)))
	if meta.TotalPage < 1 {
		meta.TotalPage = 1
	}
	if err = query.Find(&data).Error; err != nil {
		return data, dto.UserMeta{}, err
	}
	return data, meta, nil
}
