package entities

type Question struct {
	ID              uint       `gorm:"primaryKey"`
	AssignmentId    int        `gorm:"not null" json:"assignment_id"`
	AssignmentTitle string     `gorm:"not null" json:"assignment_title"`
	AssignmentsId   Assignment `gorm:"ForeignKey:AssignmentId;References:ID;null" json:"-"`
	UserId          int        `gorm:"not null" json:"user_id"`
	User            User       `gorm:"ForeignKey:UserId;References:ID" json:"-"`
	QuestionUser    string     `json:"question_user" validate:"required"`
	Name            string     `json:"name" validate:"required" gorm:"column:name"`
}

type QuestionResponse struct {
	QuestionUser string `json:"question_user" validate:"required"`
	UserId       int    `json:"user_id" validate:"required"`
	AssignmentId int    `json:"assignment_id" validate:"required"`
}

func (*Question) TableName() string {
	return "questions"
}
