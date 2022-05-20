package assignment

import (
	"Tugas-Mini-Project/domains"
	"Tugas-Mini-Project/entities"
	"Tugas-Mini-Project/infrastructure/database"
)

type svcAssignment struct {
	c    database.Config
	repo domains.AssignmentRepository
}

func NewAssignmentService(repo domains.AssignmentRepository, c database.Config) domains.AssignmentService {
	return &svcAssignment{
		c:    c,
		repo: repo,
	}
}

func (s *svcAssignment) CreateAssignmentService(assignment entities.Assignment) (string, error) {
	if assignment.UserId != 1 {
		return "you're not teacher", nil
	} else {
		return "teacher can create assignment", s.repo.CreateAssignment(assignment)
	}
}

func (s *svcAssignment) GetAssignmentByIdService(id int) (entities.Assignment, error) {
	return s.repo.GetAssignmentById(id)
}

func (s *svcAssignment) GetAllAssignmentService() (assignment []entities.Assignment, err error) {
	return s.repo.GetAllAssignment()
}

func (s *svcAssignment) UpdateAssignmentService(id int, assignment entities.Assignment) (entities.Assignment, error) {
	if assignment.UserId != 1 {
		return assignment, nil
	} else {
		return s.repo.UpdateAssignment(id, assignment)
	}
}

func (s *svcAssignment) DeleteAssignmentService(id int, assignment entities.Assignment) (string, error) {
	if assignment.UserId != 1 {
		return "you can't delete the assignment", nil
	} else {
		return "you can delete the assignment by id", s.repo.DeleteAssignment(id, assignment)
	}
}
