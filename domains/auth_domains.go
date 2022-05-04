package domains

import (
	"Tugas-Mini-Project/entities"
	"github.com/labstack/echo/v4"
)

type AuthHandler interface {
	LoginHandler(c echo.Context) error
	RegisterHandler(c echo.Context) error
}

type AuthRepository interface {
	Login(credential entities.User) error
	Register(user entities.User) error
}
