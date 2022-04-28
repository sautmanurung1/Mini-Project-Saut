package model

import "gorm.io/gorm"

type TotalStudent struct {
	gorm.Model
	Id           int        `json:"id" gorm:"primary_key"`
	AssignmentId int        `json:"assignment_id"`
	Assignment   Assignment `gorm:"foreignkey:AssignmentId"`
	UserId       int        `json:"user_id"`
	User         User       `json:"user" gorm:"foreignkey:UserId"`
	TotalUser    int        `json:"total_user" validate:"required"`
}
