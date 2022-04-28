package model

import "gorm.io/gorm"

type Discussions struct {
	gorm.Model
	Id          int      `json:"id" gorm:"primary_key"`
	UserId      int      `json:"user_id"`
	User        User     `gorm:"foreignkey:UserId"`
	Name        string   `json:"name" validate:"required"`
	QuestionId  int      `json:"question_id"`
	Question    Question `gorm:"foreignkey:QuestionId"`
	AnswareId   int      `json:"answare_id"`
	Answare     Answare  `gorm:"foreignkey:AnswareId"`
	AnswereUser string   `json:"answareuser" validate:"required"`
}
