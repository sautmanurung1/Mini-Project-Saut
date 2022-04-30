package assignment

import "Tugas-Mini-Project/model"

type AssignmentRepository interface {
	CreateAssignment(assignment model.Assignment) error
	GetAssignmentById(id int) (model.Assignment, error)
	GetAllAssignment() ([]model.Assignment, error)
	UpdateAssignment(id int, assignment model.Assignment) (model.Assignment, error)
	DeleteAssignment(id int) (model.Assignment, error)
}
