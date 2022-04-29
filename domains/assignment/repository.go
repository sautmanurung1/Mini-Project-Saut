package assignment

import "Tugas-Mini-Project/model"

type AssignmentRepository interface {
	CreateAssignment(assignment model.Assignment) error
	GetAssignmentById(id int) (model.Assignment, error)
	GetAllAssignment() ([]model.Assignment, error)
	UpdateAssignment(assignment model.Assignment) error
	DeleteAssignment(id int) error
}
