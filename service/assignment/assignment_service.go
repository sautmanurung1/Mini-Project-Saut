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
		return "You're Not Teacher", nil
	} else {
		return "Teacher can Create Assignment", s.repo.CreateAssignment(assignment)
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
		return "You Can't Delete The Assignment", nil
	} else {
		return "You Can Delete The Assignment By ID", s.repo.DeleteAssignment(id, assignment)
	}
}
