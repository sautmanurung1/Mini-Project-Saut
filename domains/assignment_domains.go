package domains

import (
	"Tugas-Mini-Project/entities"
	"github.com/labstack/echo/v4"
)

type AssignmentHandler interface {
	CreateAssignmentHandler(c echo.Context) error
	GetAssignmentByIdHandler(c echo.Context) error
	GetAllAssignmentHandler(c echo.Context) error
	UpdateAssignmentHandler(c echo.Context) error
	DeleteAssignmentHandler(c echo.Context) error
}

type AssignmentRepository interface {
	CreateAssignment(assignment entities.Assignment) error
	GetAssignmentById(id int) (entities.Assignment, error)
	GetAllAssignment() ([]entities.Assignment, error)
	UpdateAssignment(id int, assignment entities.Assignment) (entities.Assignment, error)
	DeleteAssignment(id int) (entities.Assignment, error)
}

type AssignmentService interface {
}
