package database

import (
	"errors"
	"mini-pos/dto"
	"mini-pos/util"
)

func initSeeder() error {
	userSeeder()
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
