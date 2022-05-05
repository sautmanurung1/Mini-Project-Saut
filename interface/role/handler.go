package role

import (
	"Tugas-Mini-Project/domains"
	"Tugas-Mini-Project/entities"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type handler struct {
	repository domains.RoleRepository
}

func NewRoleHandler(repository domains.RoleRepository) domains.RoleHandler {
	return &handler{
		repository: repository,
	}
}

func (h *handler) GetRole(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
	}

	roles, e := h.repository.GetRole(id)

	if e != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": e.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Get Role Success",
		"data":    roles,
	})
}

func (h *handler) CreateRole(c echo.Context) error {
	roles := entities.Role{}
	err := c.Bind(&roles)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
	}

	e := h.repository.CreateRole(roles)

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
