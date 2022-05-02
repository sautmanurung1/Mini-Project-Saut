package question

import (
	"Tugas-Mini-Project/domains/assignment"
	"Tugas-Mini-Project/domains/auth"
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	AssignmentId int                   `gorm:"not null"`
	Assignment   assignment.Assignment `gorm:"ForeignKey:AssignmentId;References:ID;null"`
	UserId       int                   `gorm:"not null"`
	User         auth.User             `gorm:"ForeignKey:UserId;References:ID"`
	QuestionUser string                `json:"question_user" validate:"required"`
	Name         string                `json:"name" validate:"required" gorm:"column:name"`
}
