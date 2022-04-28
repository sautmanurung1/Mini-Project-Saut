package auth

import (
	"Tugas-Mini-Project/domains/auth"
	"Tugas-Mini-Project/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

type handler struct {
	repository auth.AuthRepository
}

func NewAuthHandler(repository auth.AuthRepository) auth.AuthHandler {
	return &handler{
		repository: repository,
	}
}

func (h *handler) LoginHandler(c echo.Context) error {
	user := model.User{}
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	e, _ := h.repository.Login(user)

	return c.JSON(http.StatusOK, e)
}

func (h *handler) RegisterHandler(c echo.Context) error {
	user := model.User{}
	role := model.Role{}
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	e := h.repository.Register(user)

	if user.RoleId == 1 {
		role.Id = 1
		role.Name = "teacher"
	}

	if user.RoleId == 2 {
		role.Id = 2
		role.Name = "student"
	}

	if e != nil {
		return c.JSON(http.StatusInternalServerError, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Success Register",
		"Data":    user,
	})
}
