package model

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	Id           int        `json:"id" gorm:"primary_key"`
	AssignmentId int        `json:"assignment_id"`
	Assignment   Assignment `gorm:"foreignkey:AssignmentId"`
	UserId       int        `json:"user_id"`
	User         User       `gorm:"foreignkey:UserId"`
	QuestionUser string     `json:"question_user" validate:"required"`
	Name         string     `json:"name" validate:"required"`
}
