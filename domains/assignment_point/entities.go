package assignment_point

import (
	"Tugas-Mini-Project/domains/assignment"
	"Tugas-Mini-Project/domains/auth"
	"gorm.io/gorm"
)

type AssignmentPoint struct {
	gorm.Model
	AssignmentID int                   `gorm:"not null"`
	Assignment   assignment.Assignment `gorm:"ForeignKey:AssignmentID;References:ID;null"`
	Point        int                   `json:"point" validate:"required"`
	UserId       int                   `gorm:"not null"`
	User         auth.User             `gorm:"ForeignKey:UserId;References:ID"`
}
