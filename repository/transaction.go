package repository

import (
	"errors"
	"mini-pos/database"
	"mini-pos/dto"
	"time"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Insert(dto.Transaction) (dto.Transaction, error)
	InsertDetail(dto.TransactionDetail) (dto.TransactionDetail, error)
	GetAll(dto.Transaction, dto.Pagination) ([]dto.Transaction, error)
	GetByID(uint) (dto.Transaction, error)
	GetDetail(dto.TransactionPayload) ([]dto.TransactionDetail, error)
	Update(dto.Transaction) (dto.Transaction, error)
	SavePayment(dto.PaymentPayload) (dto.Transaction, error)
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

func (repo *transactionRepo) Update(payload dto.Transaction) (data dto.Transaction, err error) {
	if err = repo.DB.First(&data, payload.Id).Error; err != nil {
		return
	}

	if data.Id == 0 {
		return data, errors.New("transaction not found")
	}

	// if time updated at is not same, possibly the data already updated by another user
	if payload.UpdatedAt != data.UpdatedAt && !payload.UpdatedAt.IsZero() {
		return data, errors.New("please refresh to sync the data first")
	}

	// set data from payload
	data = payload

	// save data
	err = repo.DB.Save(&data).Error
	return payload, err
}

func (repo *transactionRepo) SavePayment(payload dto.PaymentPayload) (data dto.Transaction, err error) {
	if err = repo.DB.First(&data, payload.TransactionID).Error; err != nil {
		return
	}

	if data.Id == 0 {
		return data, errors.New("transaction not found")
	}

	// set data from payload
	data.PaymentNumber = payload.PaymentNumber
	data.PaymentDate = time.Now()
	data.PaymentNominal = payload.PaymentNominal
	data.PaymentMethod = payload.PaymentMethod
	data.PaymentNote = payload.PaymentNote
	data.IsStatus = 1

	// save data
	err = repo.DB.Save(&data).Error
	return data, err
}
