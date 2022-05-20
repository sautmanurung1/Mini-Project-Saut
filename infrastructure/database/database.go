package database

import (
	"Tugas-Mini-Project/entities"
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
	JWT_KEY     string
}

func InitDB(conf Config) *gorm.DB {
	conf = ENVDatabase()
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.DB_USERNAME,
		conf.DB_PASSWORD,
		conf.DB_HOST,
		conf.DB_PORT,
		conf.DB_NAME,
	)
	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		// DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}

	e := DB.AutoMigrate(&entities.User{}, &entities.Role{}, &entities.Assignment{}, &entities.Answare{}, entities.Discussions{}, &entities.Question{})
	if e != nil {
		panic(e)
	}

	return DB
}
