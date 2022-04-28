package model

import "gorm.io/gorm"

type Assignment struct {
	gorm.Model
	UserId      int    `gorm:"not null"`
	User        User   `gorm:"ForeignKey:UserId;References:ID"`
	Description string `json:"description" validate:"required"`
	Questions   string `json:"questions" validate:"required"`
	Name        string `json:"name" validate:"required"`
}
