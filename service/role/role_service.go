package role

import (
	"Tugas-Mini-Project/domains"
	"Tugas-Mini-Project/entities"
	"Tugas-Mini-Project/infrastructure/database"
)

type svcRole struct {
	c    database.Config
	repo domains.RoleRepository
}

func NewRoleService(repo domains.RoleRepository, c database.Config) domains.RoleService {
	return &svcRole{
		c:    c,
		repo: repo,
	}
}

func (s *svcRole) CreateRoleService(role entities.Role) error {
	return s.repo.CreateRole(role)
}
