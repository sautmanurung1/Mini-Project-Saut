package auth

import (
	"Tugas-Mini-Project/internal/entities"
	"Tugas-Mini-Project/service/auth"
	"fmt"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) auth.AuthRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) Register(user entities.User) error {
	role := entities.Role{}
	response := r.DB.Create(&user)
	user.RoleId = role.ID
	user.Role = role

	if user.RoleId == 1 {
		role := entities.Role{
			ID:   1,
			Name: "teacher",
		}
		r.DB.Create(&role)
	} else if user.RoleId == 2 {
		role := entities.Role{
			ID:   2,
			Name: "student",
		}
		r.DB.Create(&role)
	}

	if response.Error != nil {
		return response.Error
	}
	return nil
}

func (r *repository) Login(credential entities.User) error {
	response := r.DB.Raw("SELECT * FROM users WHERE username = ? AND password = ? AND role_id = ?", credential.Username, credential.Password, credential.RoleId).Scan(&credential)

	if response.RowsAffected < 1 {
		return fmt.Errorf("error to login")
	}

	return nil
}
