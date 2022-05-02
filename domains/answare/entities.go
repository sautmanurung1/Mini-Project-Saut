package answare

import (
	"Tugas-Mini-Project/domains/assignment"
	"Tugas-Mini-Project/domains/auth"
	"gorm.io/gorm"
)

type Answare struct {
	gorm.Model
	AssignmentId int                   `gorm:"not null"`
	Assignment   assignment.Assignment `gorm:"ForeignKey:AssignmentId;References:ID;null"`
	UserId       int                   `gorm:"not null"`
	User         auth.User             `gorm:"ForeignKey:UserId;References:ID"`
	Questions    string                `json:"questions"`
	Name         string                `json:"name"`
}
