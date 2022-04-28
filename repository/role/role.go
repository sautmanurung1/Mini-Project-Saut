package role

import (
	"Tugas-Mini-Project/domains/role"
	"Tugas-Mini-Project/model"
	"fmt"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewRoleRepository(db *gorm.DB) role.RoleRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) CreateRole(role model.Role) error {
	response := r.DB.Create(&role)
	if response.RowsAffected < 1 {
		return fmt.Errorf("failed to create role")
	}
	return nil
}

func (r *repository) GetRole(credential model.Role) (model.Role, error) {
	roles := model.Role{}
	response := r.DB.Where("name = ?", credential).Find(&roles)

	// make if role teacher
	if credential.Name == "teacher" {
		roles.Id = 1
		roles.Name = "teacher"
	}

	// make if role student
	if credential.Name == "student" {
		roles.Id = 2
		roles.Name = "student"
	}

	return roles, response.Error
}
