package answare

import (
	"Tugas-Mini-Project/domains/auth"
	"Tugas-Mini-Project/domains/question"
	"gorm.io/gorm"
)

type Answare struct {
	gorm.Model
	QuestionId  int               `gorm:"not null"`
	Question    question.Question `gorm:"ForeignKey:QuestionId;References:ID;null"`
	UserId      int               `gorm:"not null"`
	User        auth.User         `gorm:"ForeignKey:UserId;References:ID"`
	Questions   string            `json:"questions"`
	AnswareUser string            `json:"answare_user" validate:"required"`
	Name        string            `json:"name" validate:"required"`
}
