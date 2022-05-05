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
	// roles := entities.Role{}
	response := r.DB.Create(&user)
	// user.RoleId = roles.ID
	// user.Role = roles
	/*
		if user.RoleId == 1 {
			roles := entities.Role{
				ID:   1,
				Name: "teacher",
			}
			r.DB.Create(&roles)
		} else if user.RoleId == 2 {
			roles := entities.Role{
				ID:   2,
				Name: "student",
			}
			r.DB.Create(&roles)
		}
	*/
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
