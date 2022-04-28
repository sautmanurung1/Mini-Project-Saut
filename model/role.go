package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Id   int    `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}
