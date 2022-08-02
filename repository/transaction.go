package repository

import (
	"mini-pos/database"
	"mini-pos/dto"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Insert(dto.Transaction) (dto.Transaction, error)
	InsertDetail(dto.TransactionDetail) (dto.TransactionDetail, error)
	GetAll(dto.Transaction, dto.Pagination) ([]dto.Transaction, error)
	// Update(dto.Transaction) (dto.Transaction, error)
	// Login(username string, password string) (dto.User, error)
	// DeleteByID(ID uint) (data dto.User, err error)
}

type transactionRepo struct {
	DB *gorm.DB
}

func InitTransactionRepository() TransactionRepository {
	return &transactionRepo{
		DB: database.DB,
	}
}

func (repo *transactionRepo) Insert(payload dto.Transaction) (data dto.Transaction, err error) {
	err = repo.DB.Create(&payload).Error
	return payload, err
}

func (repo *transactionRepo) InsertDetail(payload dto.TransactionDetail) (data dto.TransactionDetail, err error) {
	err = repo.DB.Create(&payload).Error
	return payload, err
}

func (repo *transactionRepo) GetAll(payload dto.Transaction, pagination dto.Pagination) (data []dto.Transaction, err error) {
	err = pagination.Apply(repo.DB).Preload("Outlet").Find(&data, payload).Error
	return
}

func (repo *transactionRepo) GetDetail(ID int) (data dto.TransactionPayload, err error) {
	err = repo.DB.Preload("Outlet").Find(&data, ID).Error
	return
}

// func (repo *transactionRepo) Update(payload dto.User) (data dto.User, err error) {

// 	// get book by id
// 	if err = repo.DB.First(&data, payload.Id).Error; err != nil {
// 		return
// 	}

// 	// update value
// 	data.Name = payload.Name
// 	data.Username = payload.Username
// 	if payload.Password != "" {
// 		data.Password, err = util.HashPassword(payload.Password)
// 		if err != nil {
// 			return dto.User{}, err
// 		}
// 	}
// 	data.PhoneNumber = payload.PhoneNumber
// 	data.IsRole = payload.IsRole
// 	data.IsActive = payload.IsActive
// 	data.UpdatedAt = time.Now()

// 	// update book data
// 	err = repo.DB.Save(&data).Error
// 	return
// }

// func (repo *transactionRepo) DeleteByID(ID uint) (data dto.User, err error) {
// 	// get book by id
// 	if err = repo.DB.First(&data, ID).Error; err != nil {
// 		return
// 	}

// 	// delete book
// 	err = repo.DB.Delete(&data, ID).Error
// 	return
// }
