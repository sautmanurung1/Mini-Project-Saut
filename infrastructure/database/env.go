package database

import (
	"github.com/joho/godotenv"
	"os"
)

func ENVDatabase() Config {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	return Config{
		DB_USERNAME: os.Getenv("DB_USERNAME"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_NAME:     os.Getenv("DB_NAME"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_HOST:     os.Getenv("DB_HOST"),
	}
}