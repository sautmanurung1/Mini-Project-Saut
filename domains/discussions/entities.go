package discussions

import (
	"Tugas-Mini-Project/domains/answare"
	"Tugas-Mini-Project/domains/auth"
	"Tugas-Mini-Project/domains/question"
	"gorm.io/gorm"
)

type Discussions struct {
	gorm.Model
	UserId       int               `gorm:"not null" json:"user_id"`
	User         auth.User         `gorm:"ForeignKey:UserId;References:ID;null" json:"-"`
	Name         string            `json:"name" validate:"required"`
	QuestionId   int               `gorm:"not null" json:"question_id"`
	Question     question.Question `gorm:"ForeignKey:QuestionId;References:ID;null" json:"-"`
	QuestionUser string            `json:"question_user" validate:"required"`
	AnswareId    int               `gorm:"not null" json:"answare_id"`
	Answare      answare.Answare   `gorm:"ForeignKey:AnswareId;References:ID;null" json:"-"`
	AnswereUser  string            `json:"answare_user" validate:"required"`
}
