package auth

import (
	"Tugas-Mini-Project/domains"
	"Tugas-Mini-Project/entities"
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

func (r *repository) Login(username string) (credential entities.User, err error) {
	err = r.DB.Raw("SELECT * FROM users WHERE username = ?", username).Scan(&credential).Error

	return credential, nil
}
