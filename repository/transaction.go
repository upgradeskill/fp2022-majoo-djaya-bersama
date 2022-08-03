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
	GetByID(uint) (dto.Transaction, error)
	GetDetail(dto.TransactionPayload) ([]dto.TransactionDetail, error)
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
	err = pagination.Apply(repo.DB).Find(&data, payload).Error
	return
}

func (repo *transactionRepo) GetByID(id uint) (data dto.Transaction, err error) {
	err = repo.DB.Find(&data, id).Error
	return
}

func (repo *transactionRepo) GetDetail(payload dto.TransactionPayload) (data []dto.TransactionDetail, err error) {
	err = repo.DB.Find(&data, payload).Error
	return
}
