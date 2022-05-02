package role

type RoleRepository interface {
	GetRole(id int) (Role, error)
	CreateRole(role Role) error
}
