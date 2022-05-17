package entities

type Assignment struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	UserId      int    `gorm:"null" json:"user_id"`
	User        User   `gorm:"ForeignKey:UserId;References:ID;null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	Description string `json:"description" validate:"required"`
	Questions   string `json:"questions" validate:"required"`
	Name        string `json:"name" validate:"required" gorm:"column:name"`
	Title       string `json:"title" validate:"required" gorm:"column:title"`
}

type AssignmentResponse struct {
	ID          int    `json:"id"`
	Questions   string `json:"questions" validate:"required"`
	Description string `json:"description" validate:"required"`
	UserId      int    `json:"user_id" validate:"required"`
	Title       string `json:"title" validate:"required"`
}

func (*Assignment) TableName() string {
	return "assignments"
}
