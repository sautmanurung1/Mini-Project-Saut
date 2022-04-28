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

	if response.RowsAffected < 1 {
		return fmt.Errorf("Error To Register")
	}
	return nil
}

func (r *repository) Login(credential model.User) (model.User, error) {
	user := model.User{}
	err := r.DB.Where("username = ?, password = ? AND role_id = ?", credential.Username, credential.RoleId).First(&user).Error
	return user, err
}
