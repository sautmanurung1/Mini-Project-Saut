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

// Register godoc
// @Summary Register User
// @Description Create Register User
// @Tags Auth
// @accept json
// @Produce json
// @Router /register [post]
// @Param data body entities.RegisterUser true "required"
// @Success 200 {object} entities.User
// @Failure 401 {object} entities.User
// @Failure 500 {object} entities.User
func (h *AuthHandler) Register(c echo.Context) error {
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

// Login godoc
// @Summary Login User
// @Description Login User
// @Tags Auth
// @accept json
// @Produce json
// @Router /login [post]
// @Param data body entities.LoginUser true "required"
// @Success 200 {object} entities.User
// @Success 400 {object} entities.User
// @Failure 401 {object} entities.User
// @Failure 500 {object} entities.User
func (h *AuthHandler) Login(c echo.Context) error {
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
			"message":    "your username and password wrong",
			"statusCode": statusCode,
		})
	} else if statusCode == http.StatusInternalServerError {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":    "internal server error",
			"statusCode": statusCode,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"data":    token,
	})
}
