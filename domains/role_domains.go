package domains

import (
	"Tugas-Mini-Project/entities"
)

type RoleRepository interface {
	CreateRole(role entities.Role) error
}

type RoleService interface {
	CreateRoleService(role entities.Role) error
}
