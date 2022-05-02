package assignment

type AssignmentRepository interface {
	CreateAssignment(assignment Assignment) error
	GetAssignmentById(id int) (Assignment, error)
	GetAllAssignment() ([]Assignment, error)
	UpdateAssignment(id int, assignment Assignment) (Assignment, error)
	DeleteAssignment(id int) (Assignment, error)
}
