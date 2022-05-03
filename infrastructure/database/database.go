package database

import (
	"Tugas-Mini-Project/domains/answare"
	"Tugas-Mini-Project/domains/assignment"
	"Tugas-Mini-Project/domains/auth"
	"Tugas-Mini-Project/domains/discussions"
	"Tugas-Mini-Project/domains/question"
	"Tugas-Mini-Project/domains/role"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DB_USERNAME   string
	DB_PASSWORD   string
	DB_NAME       string
	DB_PORT       string
	DB_HOST       string
	Login_Teacher string
	Login_Student string
}

func InitDB() *gorm.DB {
	conf := ENVDatabase()
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

	e := DB.AutoMigrate(&auth.User{}, &role.Role{}, &assignment.Assignment{}, &answare.Answare{}, discussions.Discussions{}, &question.Question{})
	if e != nil {
		panic(e)
	}

	return DB
}
