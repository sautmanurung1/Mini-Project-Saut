package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int    `json:"id" gorm:"primary_key"`
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	RoleId   int    `json:"role_id"`
	Role     Role   `json:"role" gorm:"foreignkey:RoleId"`
}
