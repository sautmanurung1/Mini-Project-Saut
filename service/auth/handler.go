package auth

import "github.com/labstack/echo/v4"

type AuthHandler interface {
	LoginHandler(c echo.Context) error
	RegisterHandler(c echo.Context) error
}
