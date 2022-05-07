package domains

import (
	"Tugas-Mini-Project/entities"
)

/*
type AssignmentHandler interface {
	CreateAssignmentHandler(c echo.Context) error
	GetAssignmentByIdHandler(c echo.Context) error
	GetAllAssignmentHandler(c echo.Context) error
	UpdateAssignmentHandler(c echo.Context) error
	DeleteAssignmentHandler(c echo.Context) error
}
*/

type AssignmentRepository interface {
	CreateAssignment(assignment entities.Assignment) error
	GetAssignmentById(id int, assign entities.Assignment) error
	GetAllAssignment() ([]entities.Assignment, error)
	UpdateAssignment(id int, assignment entities.Assignment) error
	DeleteAssignment(id int, assign entities.Assignment) error
}

type AssignmentService interface {
	CreateAssignmentService(assignment entities.Assignment) (string, error)
	GetAssignmentByIdService(id int, assignment entities.Assignment) (string, error)
	GetAllAssignmentService() (assignment []entities.Assignment, err error)
	UpdateAssignmentService(id int, assignment entities.Assignment) (string, error)
	DeleteAssignmentService(id int, assignment entities.Assignment) (string, error)
}
