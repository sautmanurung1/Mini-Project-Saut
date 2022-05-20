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

// CreateRole godoc
// @Summary Create Role
// @Description Create Role user
// @Tags Role
// @accept json
// @Produce json
// @Router /role [post]
// @Param data body entities.Role true "required"
// @Success 200 {object} entities.Role
// @Success 400 {object} entities.Role
// @Failure 404 {object} entities.Role
func (h *RoleHandler) CreateRole(c echo.Context) error {
	roles := entities.Role{}
	err := c.Bind(&roles)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status":  "data not found",
			"message": err.Error(),
		})
	}

	e := h.Svc.CreateRoleService(roles)

	if e != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error to create role",
			"message": e.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success to create the role",
		"Data":    roles,
	})
}
