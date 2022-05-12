package role

import (
	"Tugas-Mini-Project/domains"
	"Tugas-Mini-Project/entities"
	"github.com/labstack/echo/v4"
	"net/http"
)

type RoleHandler struct {
	Svc domains.RoleService
}

func (h *RoleHandler) CreateRole(c echo.Context) error {
	roles := entities.Role{}
	err := c.Bind(&roles)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status":  "Data Not Found",
			"message": err.Error(),
		})
	}

	e := h.Svc.CreateRoleService(roles)

	if e != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": e.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Success Register",
		"Data":    roles,
	})
}
