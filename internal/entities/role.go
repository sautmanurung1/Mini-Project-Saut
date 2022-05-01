package entities

type Role struct {
	ID   int    `json:"id" gorm:"PrimaryKey;AUTO_INCREMENT;column:id"`
	Name string `json:"name"`
}
