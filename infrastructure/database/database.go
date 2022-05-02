package database

import (
	"Tugas-Mini-Project/domains/answare"
	"Tugas-Mini-Project/domains/assignment"
	"Tugas-Mini-Project/domains/assignment_point"
	"Tugas-Mini-Project/domains/auth"
	"Tugas-Mini-Project/domains/discussions"
	"Tugas-Mini-Project/domains/question"
	"Tugas-Mini-Project/domains/role"
	entities2 "Tugas-Mini-Project/domains/total_student"
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

	e := DB.AutoMigrate(&auth.User{}, &role.Role{}, &entities2.TotalStudent{}, &assignment.Assignment{}, &assignment_point.AssignmentPoint{}, &answare.Answare{}, discussions.Discussions{}, &question.Question{})
	if e != nil {
		panic(e)
	}

	return DB
}
