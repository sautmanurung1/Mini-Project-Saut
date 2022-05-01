package assignment

import (
	"Tugas-Mini-Project/internal/entities"
)

type AssignmentRepository interface {
	CreateAssignment(assignment entities.Assignment) error
	GetAssignmentById(id int) (entities.Assignment, error)
	GetAllAssignment() ([]entities.Assignment, error)
	UpdateAssignment(id int, assignment entities.Assignment) (entities.Assignment, error)
	DeleteAssignment(id int) (entities.Assignment, error)
}
