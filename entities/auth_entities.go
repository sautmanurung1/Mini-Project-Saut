package entities

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	RoleId   int    `gorm:"not null" json:"role_id"`
	Role     Role   `gorm:"ForeignKey:RoleId;references:ID;null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
}

type LoginUser struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterUser struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	RoleId   int    `json:"role_id" required:"required"`
}

func (*User) TableName() string {
	return "users"
}
