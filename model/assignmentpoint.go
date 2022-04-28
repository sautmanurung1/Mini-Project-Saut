package model

import "gorm.io/gorm"

type AssignmentPoint struct {
	gorm.Model
	Id           int        `json:"id" gorm:"primary_key"`
	AssignmentId int        `json:"assignment_id"`
	Assignment   Assignment `gorm:"foreignkey:AssignmentId"`
	UserId       int        `json:"user_id"`
	User         User       `gorm:"foreignkey:UserId"`
}
