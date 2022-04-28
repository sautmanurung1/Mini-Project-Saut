package model

import "gorm.io/gorm"

type Answare struct {
	gorm.Model
	Id          int      `json:"id" gorm:"primary_key"`
	QuestionId  int      `json:"question_id"`
	Question    Question `gorm:"foreignkey:QuestionId"`
	UserId      int      `json:"user_id"`
	User        User     `gorm:"foreignkey:UserId"`
	Name        string   `json:"name" validate:"required"`
	AnswareUser string   `json:"answare_user" validate:"required"`
}
