package question

import (
	"Tugas-Mini-Project/domains/assignment"
	"Tugas-Mini-Project/domains/auth"
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	AssignmentId     int                   `gorm:"not null" json:"assignment_id"`
	AssignmentTitle  string                `gorm:"not null" json:"assignment_title"`
	AssignmentsId    assignment.Assignment `gorm:"ForeignKey:AssignmentId;References:ID;null" json:"-"`
	AssignmentsTitle assignment.Assignment `gorm:"ForeignKey:AssignmentTitle;References:title;null" json:"-"`
	UserId           int                   `gorm:"not null" json:"user_id"`
	User             auth.User             `gorm:"ForeignKey:UserId;References:ID" json:"-"`
	QuestionUser     string                `json:"question_user" validate:"required"`
	Name             string                `json:"name" validate:"required" gorm:"column:name"`
}
