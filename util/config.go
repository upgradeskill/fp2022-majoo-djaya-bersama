package util

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_HOST     string
	DB_PORT     string
	JWT_SECRET  string
	JWT_ISSUER  string
}

var GlobalConfig Config

func LoadConfig() {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatalf("Error load env: %s", err)
	}
	GlobalConfig.DB_USER = os.Getenv("DB_USER")
	GlobalConfig.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	GlobalConfig.DB_NAME = os.Getenv("DB_NAME")
	GlobalConfig.DB_HOST = os.Getenv("DB_HOST")
	GlobalConfig.DB_PORT = os.Getenv("DB_PORT")
	GlobalConfig.JWT_SECRET = os.Getenv("JWT_SECRET")
	GlobalConfig.JWT_ISSUER = os.Getenv("JWT_ISSUER")

}
