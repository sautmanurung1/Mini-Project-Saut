package role

import (
	"Tugas-Mini-Project/domains/role"
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

func (r *repository) CreateRole(role role.Role) error {
	response := r.DB.Create(&role)
	if response.RowsAffected < 1 {
		return fmt.Errorf("failed to create role")
	}
	return nil
}

func (r *repository) GetRole(id int) (role.Role, error) {
	roles := role.Role{}
	response := r.DB.Where("ID = ?", id).Find(&roles)

	if response.RowsAffected < 1 {
		return roles, response.Error
	}
	return roles, nil
}
