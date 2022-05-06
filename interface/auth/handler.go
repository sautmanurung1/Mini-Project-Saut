package auth

import (
	"Tugas-Mini-Project/domains"
	"Tugas-Mini-Project/entities"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthHandler struct {
	svc domains.AuthService
}

func (h *AuthHandler) RegisterHandler(c echo.Context) error {
	user := entities.User{}
	err := c.Bind(&user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
	}

	er := h.svc.RegisterService(user)

	if er != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"Status":  http.StatusNotFound,
			"Message": "User UnRegistered",
			"Data":    er.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Success Register",
		"Data":    user,
	})
}

func (h *AuthHandler) LoginHandler(c echo.Context) error {
	userLogin := make(map[string]interface{})

	err := c.Bind(&userLogin)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Bad Request",
			"error":   err.Error(),
		})
	}

	token, statusCode := h.svc.LoginService(userLogin["username"].(string), userLogin["password"].(string), userLogin["role_id"].(int))

	if statusCode == http.StatusUnauthorized {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Unauthorized",
		})
	} else if statusCode == http.StatusInternalServerError {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Login",
		"data":    token,
	})
}
