package entities

import "gorm.io/gorm"

type Answare struct {
	gorm.Model
	AssignmentId int        `gorm:"not null"`
	Assignment   Assignment `gorm:"ForeignKey:AssignmentId;References:ID;null"`
	UserId       int        `gorm:"not null"`
	User         User       `gorm:"ForeignKey:UserId;References:ID"`
	Questions    string     `json:"questions"`
	Name         string     `json:"name"`
}
