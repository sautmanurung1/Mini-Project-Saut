package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	RoleId   int    `gorm:"not null" json:"roleId"`
	Role     Role   `gorm:"ForeignKey:RoleId;references:ID"`
}