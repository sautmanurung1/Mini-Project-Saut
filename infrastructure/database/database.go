package database

import (
	entities2 "Tugas-Mini-Project/internal/entities"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
	DB_PORT     string
	DB_HOST     string
}

func InitDB() *gorm.DB {
	conf := Config{
		DB_USERNAME: "root",
		DB_PASSWORD: "Sautmanurung234",
		DB_NAME:     "go_rest_api",
		DB_PORT:     "3306",
		DB_HOST:     "localhost",
	}
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.DB_USERNAME,
		conf.DB_PASSWORD,
		conf.DB_HOST,
		conf.DB_PORT,
		conf.DB_NAME,
	)

	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}

	e := DB.AutoMigrate(&entities2.User{}, &entities2.Role{}, &entities2.TotalStudent{}, &entities2.Assignment{}, &entities2.AssignmentPoint{}, &entities2.Answare{}, entities2.Discussions{}, &entities2.Question{})
	if e != nil {
		panic(e)
	}

	return DB
}
