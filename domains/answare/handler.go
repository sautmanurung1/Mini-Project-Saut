package answare

import "github.com/labstack/echo/v4"

type AnswareHandler interface {
	CreateAnswareHandler(c echo.Context) error
	GetAnswareByIdHandler(c echo.Context) error
	GetAllAnswareHandler(c echo.Context) error
	UpdateAnswareHandler(c echo.Context) error
	DeleteAnswareHandler(c echo.Context) error
}
