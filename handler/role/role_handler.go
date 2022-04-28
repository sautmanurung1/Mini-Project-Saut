package role

import (
	"Tugas-Mini-Project/domains/role"
	"Tugas-Mini-Project/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

type handler struct {
	repository role.RoleRepository
}

func NewRoleHandler(repository role.RoleRepository) role.RoleHandler {
	return &handler{
		repository: repository,
	}
}

func (h *handler) GetRole(c echo.Context) error {
	roles := model.Role{}
	err := c.Bind(&roles)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	e, _ := h.repository.GetRole(roles)

	return c.JSON(http.StatusOK, e)
}

func (h *handler) CreateRole(c echo.Context) error {
	roles := model.Role{}
	err := c.Bind(&roles)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	e := h.repository.CreateRole(roles)

	if e != nil {
		return c.JSON(http.StatusInternalServerError, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Success Register",
		"Data":    roles,
	})
}
