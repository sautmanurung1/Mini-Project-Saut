package auth

import (
	"Tugas-Mini-Project/domains/auth"
	"Tugas-Mini-Project/model"
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

func (r *repository) Register(user model.User) error {
	response := r.DB.Create(&user)

	if user.RoleId == 1 {
		role := model.Role{
			ID:   1,
			Name: "teacher",
		}
		r.DB.Create(&role)
	} else if user.RoleId == 2 {
		role := model.Role{
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

func (r *repository) Login(credential model.User) error {
	response := r.DB.Where("username = ? AND password = ?", credential.Username, credential.Password).Find(&credential)

	if response.RowsAffected < 1 {
		return fmt.Errorf("error to login")
	}

	return nil
}
