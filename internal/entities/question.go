package entities

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	AssignmentId int        `gorm:"not null"`
	Assignment   Assignment `gorm:"ForeignKey:AssignmentId;References:ID"`
	UserId       int        `gorm:"not null"`
	User         User       `gorm:"ForeignKey:UserId;References:ID"`
	QuestionUser string     `json:"question_user" validate:"required"`
	Name         string     `json:"name" validate:"required"`
}
