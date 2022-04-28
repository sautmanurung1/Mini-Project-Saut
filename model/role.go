package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name string `json:"name"`
}
