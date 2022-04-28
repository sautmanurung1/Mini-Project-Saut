package model

import "gorm.io/gorm"

type Assignment struct {
	gorm.Model
	Id           int    `json:"id" gorm:"primary_key"`
	UserId       int    `json:"user_id"`
	User         User   `gorm:"foreignkey:UserId"`
	Description  string `json:"description" validate:"required"`
	QuestionUser string `json:"question" validate:"required"`
	Name         string `json:"name" validate:"required"`
}
