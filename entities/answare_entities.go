package entities

type Answare struct {
	ID          uint     `gorm:"primaryKey" json:"id"`
	QuestionId  int      `gorm:"not null" json:"question_id"`
	Question    Question `gorm:"ForeignKey:QuestionId;References:ID;null" json:"-"`
	UserId      int      `gorm:"not null" json:"user_id"`
	User        User     `gorm:"ForeignKey:UserId;References:ID" json:"-"`
	Questions   string   `gorm:"not null"`
	AnswareUser string   `json:"answare_user" validate:"required"`
	Name        string   `json:"name" validate:"required"`
}

type AnswareResponse struct {
	UserAnsware string `json:"answare_user" validate:"required"`
	IdUser      int    `json:"user_id" validate:"required"`
	IdQuestion  int    `json:"question_id" validate:"required"`
}

func (*Answare) TableName() string {
	return "answares"
}
