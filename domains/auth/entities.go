package auth

import (
	"Tugas-Mini-Project/domains/role"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string    `json:"name" validate:"required"`
	Username string    `json:"username" validate:"required"`
	Password string    `json:"password" validate:"required"`
	RoleId   int       `gorm:"not null" json:"role_id"`
	Role     role.Role `gorm:"ForeignKey:RoleId;references:ID;null" json:"-"`
}

func (*User) TableName() string {
	return "User"
}
