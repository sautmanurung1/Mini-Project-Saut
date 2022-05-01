package auth

import (
	"Tugas-Mini-Project/domains/auth"
	"Tugas-Mini-Project/internal/entities"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
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
	user := entities.User{}

	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  http.StatusBadRequest,
			"Message": err.Error(),
		})
	}

	er := h.repository.Login(user)

	if er != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status":  http.StatusInternalServerError,
			"Message": er.Error(),
		})
	}
	if user.RoleId == 1 {
		claims := jwt.MapClaims{}
		claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		jwtToken, e := token.SignedString([]byte("loginTeacher"))

		if e != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"Status":  http.StatusInternalServerError,
				"Message": e.Error(),
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"Message": "Success Create JWT in Teacher Login",
			"Token":   jwtToken,
		})
	} else {
		claims := jwt.MapClaims{}
		claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		jwtToken, e := token.SignedString([]byte("loginStudent"))

		if e != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"Status":  http.StatusInternalServerError,
				"Message": e.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Message": "Success Create JWT in Student Login",
			"Token":   jwtToken,
		})
	}
}

func (h *handler) RegisterHandler(c echo.Context) error {
	user := entities.User{}
	err := c.Bind(&user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
	}

	users := h.repository.Register(user)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Success Register",
		"Data":    users,
	})
}
