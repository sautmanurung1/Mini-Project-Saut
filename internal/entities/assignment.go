package entities

import "gorm.io/gorm"

type Assignment struct {
	gorm.Model
	UserId      int    `gorm:"not null" json:"user_id"`
	User        User   `gorm:"ForeignKey:UserId;References:ID;null"`
	Description string `json:"description" validate:"required"`
	Questions   string `json:"questions" validate:"required"`
	Name        string `json:"name" validate:"required" gorm:"column:name"`
	Title       string `json:"title" validate:"required" gorm:"column:title"`
}
