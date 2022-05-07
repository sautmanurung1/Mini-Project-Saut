package entities

import (
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	AssignmentId    int        `gorm:"not null" json:"assignment_id"`
	AssignmentTitle string     `gorm:"not null" json:"assignment_title"`
	AssignmentsId   Assignment `gorm:"ForeignKey:AssignmentId;References:ID;null" json:"-"`
	UserId          int        `gorm:"not null" json:"user_id"`
	User            User       `gorm:"ForeignKey:UserId;References:ID" json:"-"`
	QuestionUser    string     `json:"question_user" validate:"required"`
	Name            string     `json:"name" validate:"required" gorm:"column:name"`
}

func (*Question) TableName() string {
	return "questions"
}
