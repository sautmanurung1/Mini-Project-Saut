package entities

type Discussions struct {
	ID           uint     `gorm:"primaryKey"`
	UserId       int      `gorm:"null" json:"user_id"`
	User         User     `gorm:"ForeignKey:UserId;References:ID;null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	Name         string   `json:"name" validate:"required"`
	QuestionId   int      `gorm:"null" json:"question_id"`
	Question     Question `gorm:"ForeignKey:QuestionId;References:ID;null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	QuestionUser string   `json:"question_user" validate:"required"`
	AnswareId    int      `gorm:"null" json:"answare_id"`
	Answare      Answare  `gorm:"ForeignKey:AnswareId;References:ID;null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	AnswereUser  string   `json:"answare_user" validate:"required"`
}

type DiscussionsResponse struct {
	UserId     int `json:"user_id" validate:"required"`
	QuestionId int `json:"question_id" validate:"required"`
	AnswareId  int `json:"answare_id" validate:"required"`
}

func (*Discussions) TableName() string {
	return "discussions"
}
