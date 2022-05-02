package discussions

import (
	"Tugas-Mini-Project/domains/answare"
	"Tugas-Mini-Project/domains/auth"
	"Tugas-Mini-Project/domains/question"
	"gorm.io/gorm"
)

type Discussions struct {
	gorm.Model
	UserId       int               `gorm:"not null"`
	User         auth.User         `gorm:"ForeignKey:UserId;References:ID;null"`
	Name         string            `json:"name" validate:"required"`
	QuestionId   int               `gorm:"not null"`
	Question     question.Question `gorm:"ForeignKey:QuestionId;References:ID;null"`
	QuestionUser string            `json:"question_user" validate:"required"`
	AnswareId    int               `gorm:"not null"`
	Answare      answare.Answare   `gorm:"ForeignKey:AnswareId;References:ID;null"`
	AnswereUser  string            `json:"answare_user" validate:"required"`
}