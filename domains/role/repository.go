package role

import "Tugas-Mini-Project/entities"

type RoleRepository interface {
	GetRole(id int) (entities.Role, error)
	CreateRole(role entities.Role) error
}
