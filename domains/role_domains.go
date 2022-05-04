package domains

import (
	"Tugas-Mini-Project/entities"
	"github.com/labstack/echo/v4"
)

type RoleHandler interface {
	GetRole(c echo.Context) error
	CreateRole(c echo.Context) error
}

type RoleRepository interface {
	GetRole(id int) (entities.Role, error)
	CreateRole(role entities.Role) error
}

type RoleService interface {
}
