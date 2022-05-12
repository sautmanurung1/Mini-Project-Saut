package domains

import (
	"Tugas-Mini-Project/entities"
)

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
