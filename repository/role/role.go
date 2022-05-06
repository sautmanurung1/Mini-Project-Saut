package role

import (
	"Tugas-Mini-Project/domains"
	"Tugas-Mini-Project/entities"
	"fmt"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewRoleRepository(db *gorm.DB) domains.RoleRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) CreateRole(role entities.Role) error {
	response := r.DB.Create(&role)
	if response.RowsAffected < 1 {
		return fmt.Errorf("failed to create role")
	}
	return nil
}
