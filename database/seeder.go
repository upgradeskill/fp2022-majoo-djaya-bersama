package database

import (
	"errors"
	"mini-pos/dto"
	"mini-pos/util"
	"strconv"
	"time"

	"github.com/shopspring/decimal"
)

func initSeeder() error {
	userSeeder()
	outletSeeder()
	categorySeeder()
	productSeeder()
	outletProductSeeder()
	transactionSeeder()
	return nil
}

func userSeeder() error {
	pass, _ := util.HashPassword("test123")
	user := dto.User{
		OutletId:    1,
		Name:        "User",
		Username:    "user@mail.com",
		PhoneNumber: "0812365444",
		Password:    pass,
		IsRole:      1,
		IsActive:    dto.IsActive{IsActive: 1},
	}

	err := DB.Create(&user).Error // add user data to database
	if err != nil {
		return errors.New("failed to seed user data")
	}
	return nil
}

func outletSeeder() error {
	outlet := dto.Outlet{
		Name:     "Outlet Name",
		Address:  "Address",
		IsActive: dto.IsActive{IsActive: 1},
	}
	err := DB.Create(&outlet).Error
	if err != nil {
		return errors.New("failed to seed outlet data")
	}
	return nil
}

func categorySeeder() error {
	category := dto.Category{
		Name:        "Item Category Test",
		Description: "",
		IsActive:    dto.IsActive{IsActive: 1},
	}
	err := DB.Create(&category).Error
	if err != nil {
		return errors.New("failed to seed user data")
	}

	return nil
}

func productSeeder() error {
	product := dto.Product{
		CategoryId:  1,
		Name:        "Product Test",
		Description: "",
		ImagePath:   "",
		IsActive:    dto.IsActive{IsActive: 1},
	}
	err := DB.Create(&product).Error
	if err != nil {
		return errors.New("failed to seed product data")
	}

	return nil
}

func outletProductSeeder() error {
	outletProduct := dto.OutletProduct{
		OutletID:  1,
		ProductID: 1,
		Stock:     1000,
		Price:     decimal.NewFromInt(5000),
		IsActive:  dto.IsActive{IsActive: 1},
	}
	err := DB.Create(&outletProduct).Error
	if err != nil {
		return errors.New("failed to seed outlet product data")
	}

	return nil
}

func transactionSeeder() error {
	for i := 1; i <= 10; i++ {
		transaction := dto.Transaction{
			OutletID:     1,
			UserID:       1,
			OrderNumber:  strconv.Itoa(i),
			OrderDate:    time.Now(),
			OrderNominal: decimal.NewFromInt((5000)),
		}
		err := DB.Create(&transaction).Error
		if err != nil {
			return errors.New("failed to seed transaction data")
		}

		detailTransaction := dto.TransactionDetail{
			TransactionID:   transaction.Id,
			OutletProductID: 1,
			ProductName:     "Product Test",
			Quantity:        1,
			Price:           decimal.NewFromInt((5000)),
		}

		err = DB.Create(&detailTransaction).Error
		if err != nil {
			return errors.New("failed to seed transaction detail data")
		}

	}

	return nil
}
