package auth

import (
	"Tugas-Mini-Project/domains/auth"
	"Tugas-Mini-Project/domains/role"
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

func (r *repository) Register(user auth.User) error {
	roles := role.Role{}
	response := r.DB.Create(&user)
	user.RoleId = roles.ID

	if user.RoleId == 1 {
		roles := role.Role{
			ID:   1,
			Name: "teacher",
		}
		user.Role = roles
	} else {
		roles := role.Role{
			ID:   2,
			Name: "student",
		}
		user.Role = roles
	}

	if response.Error != nil {
		return response.Error
	}
	return nil
}

func (r *repository) Login(credential auth.User) error {
	response := r.DB.Raw("SELECT * FROM users WHERE username = ? AND password = ? AND role_id = ?", credential.Username, credential.Password, credential.RoleId).Scan(&credential)

	if response.RowsAffected < 1 {
		return fmt.Errorf("error to login")
	}

	return nil
}
