package role

import "Tugas-Mini-Project/model"

type RoleRepository interface {
	GetRole(id int) (model.Role, error)
	CreateRole(role model.Role) error
}
