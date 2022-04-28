package role

import "github.com/labstack/echo/v4"

type RoleHandler interface {
	GetRole(c echo.Context) error
	CreateRole(c echo.Context) error
}
