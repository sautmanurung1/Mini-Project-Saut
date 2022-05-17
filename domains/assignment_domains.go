package domains

import (
	"Tugas-Mini-Project/entities"
)

type AssignmentRepository interface {
	CreateAssignment(assignment entities.Assignment) error
	GetAssignmentById(id int) (entities.Assignment, error)
	GetAllAssignment() ([]entities.Assignment, error)
	UpdateAssignment(id int, assignment entities.Assignment) (entities.Assignment, error)
	DeleteAssignment(id int, assign entities.Assignment) error
}

type AssignmentService interface {
	CreateAssignmentService(assignment entities.Assignment) (string, error)
	GetAssignmentByIdService(id int) (entities.Assignment, error)
	GetAllAssignmentService() (assignment []entities.Assignment, err error)
	UpdateAssignmentService(id int, assignment entities.Assignment) (entities.Assignment, error)
	DeleteAssignmentService(id int, assignment entities.Assignment) (string, error)
}
