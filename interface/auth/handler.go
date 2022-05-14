package auth

import (
	"Tugas-Mini-Project/domains"
	"Tugas-Mini-Project/entities"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthHandler struct {
	Svc domains.AuthService
}

// Register User godoc
// @Summary Register User
// @Description Create Register User
// @Tags Auth
// @accept json
// @Produce json
// @Router /register [post]
// @Param data body entities.User true "required"
// @Success 200 {object} entities.User
// @Failure 401 {object} entities.User
// @Failure 500 {object} entities.User
func (h *AuthHandler) RegisterHandler(c echo.Context) error {
	user := entities.User{}
	err := c.Bind(&user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
	}

	er := h.Svc.RegisterService(user)

	if er != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"Message": "User UnRegistered",
			"Data":    er.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Success Register",
		"Data":    user,
	})
}

// Login User godoc
// @Summary Login User
// @Description Login User
// @Tags Auth
// @accept json
// @Produce json
// @Router /login [post]
// @Param data body entities.User true "required"
// @Success 200 {object} entities.User
// @Success 400 {object} entities.User
// @Failure 401 {object} entities.User
// @Failure 500 {object} entities.User
func (h *AuthHandler) LoginHandler(c echo.Context) error {
	userLogin := entities.User{}

	err := c.Bind(&userLogin)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "Bad Request",
			"error":  err.Error(),
		})
	}

	token, statusCode := h.Svc.LoginService(userLogin.Username, userLogin.Password, userLogin.RoleId)

	if statusCode == http.StatusUnauthorized {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message":    "Your Username and Password Wrong",
			"status":     "Unauthorized",
			"statusCode": statusCode,
		})
	} else if statusCode == http.StatusInternalServerError {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":    "Internal Server Error",
			"status":     "Server Error",
			"statusCode": statusCode,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Login",
		"status":  "Success",
		"data":    token,
	})
}
