package database

import (
	"errors"
	"math/rand"
	"mini-pos/dto"
	"mini-pos/util"

	"github.com/icrowley/fake"
)

func initSeeder() error {
	userSeeder()
	productSeeder()
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
		IsActive:    1,
	}

	err := DB.Create(&user).Error // add user data to database
	if err != nil {
		return errors.New("failed to seed user data")
	}
	return nil
}

func productSeeder() error {
	products := []dto.Product{}

	err := fake.SetLang("en")
	if err != nil {
		panic(err)
	}

	for i := 0; i < 100; i++ {
		products = append(products, dto.Product{
			CategoryId: uint(rand.Intn(5)),
			Name: fake.ProductName(),
			Description: fake.Words(),
			ImagePath: fake.Word(),
		})
	}

	DB.Create(&products)
	return nil
}
