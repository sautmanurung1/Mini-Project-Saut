package auth

import (
	"Tugas-Mini-Project/domains"
	"Tugas-Mini-Project/entities"
	"fmt"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) domains.AuthRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) Register(user entities.User) error {
	response := r.DB.Create(&user)

	if response.Error != nil {
		return response.Error
	}
	return nil
}

func (r *repository) Login(username, password string, roleId int) (credential entities.User, err error) {
	response := r.DB.Raw("SELECT * FROM users WHERE username = ? AND password = ? AND role_id = ?", username, password, roleId).Scan(&credential)

	if response.RowsAffected < 1 {
		return credential, fmt.Errorf("username or password is incorrect")
	}

	return credential, nil
}
