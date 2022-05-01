package entities

import "gorm.io/gorm"

type TotalStudent struct {
	gorm.Model
	AssignmentId int        `gorm:"not null"`
	Assignment   Assignment `gorm:"ForeignKey:AssignmentId;References:ID;null"`
	UserId       int        `gorm:"not null"`
	User         User       `gorm:"ForeignKey:UserId;References:ID;null"`
	TotalUser    int        `json:"total_user" validate:"required"`
}
