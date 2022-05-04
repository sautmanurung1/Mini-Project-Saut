package entities

import (
	"gorm.io/gorm"
)

type Answare struct {
	gorm.Model
	QuestionId  int      `gorm:"not null" json:"question_id"`
	Question    Question `gorm:"ForeignKey:QuestionId;References:ID;null" json:"-"`
	UserId      int      `gorm:"not null" json:"user_id"`
	User        User     `gorm:"ForeignKey:UserId;References:ID" json:"-"`
	Questions   string   `gorm:"not null"`
	AnswareUser string   `json:"answare_user" validate:"required"`
	Name        string   `json:"name" validate:"required"`
}

func (*Answare) TableName() string {
	return "Answare"
}
