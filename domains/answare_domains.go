package domains

import (
	"Tugas-Mini-Project/entities"
	"github.com/labstack/echo/v4"
)

type AnswareRepository interface {
	CreateAnsware(answare entities.Answare) error
	GetAnswareById(id int) (entities.Answare, error)
	GetAllAnsware() ([]entities.Answare, error)
	UpdateAnsware(id int, answare entities.Answare) (entities.Answare, error)
	DeleteAnsware(id int) (entities.Answare, error)
}

type AnswareHandler interface {
	CreateAnswareHandler(c echo.Context) error
	GetAnswareByIdHandler(c echo.Context) error
	GetAllAnswareHandler(c echo.Context) error
	UpdateAnswareHandler(c echo.Context) error
	DeleteAnswareHandler(c echo.Context) error
}

type AnswareService interface {
}
