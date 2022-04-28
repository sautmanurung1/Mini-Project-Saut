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
	if response.Error != nil {
		return response.Error
	}

	if user.RoleId == 1 {
		role := model.Role{
			Name: "teacher",
		}
		r.DB.Create(&role)
	} else if user.RoleId == 2 {
		role := model.Role{
			Name: "student",
		}
		r.DB.Create(&role)
	}
	return nil
}

func (r *repository) Login(credential model.User) error {
	response := r.DB.Where("username = ? AND password = ?", credential.Username, credential.Password).Find(&credential)

	if response.RowsAffected < 1 {
		return fmt.Errorf("Error To Login")
	}

	return nil
}
