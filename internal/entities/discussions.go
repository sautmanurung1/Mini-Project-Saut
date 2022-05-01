package entities

import "gorm.io/gorm"

type Discussions struct {
	gorm.Model
	UserId       int      `gorm:"not null"`
	User         User     `gorm:"ForeignKey:UserId;References:ID"`
	Name         string   `json:"name" validate:"required"`
	QuestionId   int      `gorm:"not null"`
	Question     Question `gorm:"ForeignKey:QuestionId;References:ID"`
	QuestionUser string   `json:"question_user" validate:"required"`
	AnswareId    int      `gorm:"not null"`
	Answare      Answare  `gorm:"ForeignKey:AnswareId;References:ID"`
	AnswereUser  string   `json:"answare_user" validate:"required"`
}
