package model

import "gorm.io/gorm"

type AssignmentPoint struct {
	gorm.Model
	AssignmentID int        `gorm:"not null"`
	Assignment   Assignment `gorm:"ForeignKey:AssignmentID;References:ID"`
	Point        int        `json:"point" validate:"required"`
	UserId       int        `gorm:"not null"`
	User         User       `gorm:"ForeignKey:UserId;References:ID"`
}
