package role

import "Tugas-Mini-Project/model"

type RoleRepository interface {
	GetRole(credential model.Role) (model.Role, error)
	CreateRole(role model.Role) error
}
