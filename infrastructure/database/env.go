package database

import (
	"os"
)

func ENVDatabase() Config {
	return Config{
		Login_Student: os.Getenv("LOGIN_STUDENT"),
		Login_Teacher: os.Getenv("LOGIN_TEACHER"),
		DB_USERNAME:   os.Getenv("DB_USERNAME"),
		DB_PASSWORD:   os.Getenv("DB_PASSWORD"),
		DB_NAME:       os.Getenv("DB_NAME"),
		DB_PORT:       os.Getenv("DB_PORT"),
		DB_HOST:       os.Getenv("DB_HOST"),
	}
}
