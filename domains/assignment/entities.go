package assignment

import (
	"Tugas-Mini-Project/domains/auth"
	"gorm.io/gorm"
)

type Assignment struct {
	gorm.Model
	UserId      int       `gorm:"not null" json:"user_id"`
	User        auth.User `gorm:"ForeignKey:UserId;References:ID;null"`
	Description string    `json:"description" validate:"required"`
	Questions   string    `json:"questions" validate:"required"`
	Name        string    `json:"name" validate:"required" gorm:"column:name"`
	Title       string    `json:"title" validate:"required" gorm:"column:title"`
}
