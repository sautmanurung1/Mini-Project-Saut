package assignment

import (
	"github.com/labstack/echo/v4"
)

type AssignmentHandler interface {
	CreateAssignmentHandler(c echo.Context) error
	GetAssignmentByIdHandler(c echo.Context) error
	GetAllAssignmentHandler(c echo.Context) error
	UpdateAssignmentHandler(c echo.Context) error
	DeleteAssignmentHandler(c echo.Context) error
}
