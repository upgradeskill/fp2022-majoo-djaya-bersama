package database

import (
	"errors"
	"fmt"
	"log"
	"mini-pos/dto"
	"mini-pos/util"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupModels() *gorm.DB {
	conName := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", util.GlobalConfig.DB_USER, util.GlobalConfig.DB_PASSWORD, util.GlobalConfig.DB_HOST, util.GlobalConfig.DB_PORT, util.GlobalConfig.DB_NAME)
	fmt.Println("conname is\t", conName)
	var err error
	DB, err = gorm.Open(mysql.Open(conName), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	err = DB.AutoMigrate(
		&dto.User{},
	)

	if err != nil {
		log.Println(err)
	}

	// Check if there already have a data (check from user)
	if err := DB.First(&dto.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		// if no data, insert seed data
		err := initSeeder()
		if err != nil {
			log.Println(err)
			return DB
		}
	}

	return DB
}
