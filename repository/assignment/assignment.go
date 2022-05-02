package assignment

import (
	"Tugas-Mini-Project/domains/assignment"
	"Tugas-Mini-Project/domains/auth"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewAssignmentRepository(db *gorm.DB) assignment.AssignmentRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) CreateAssignment(assignment assignment.Assignment) error {
	// query to get the name from table user to insert name to table assignment
	var user auth.User
	r.DB.Where("id = ?", assignment.UserId).First(&user)
	assignment.Name = user.Name
	r.DB.Create(&assignment)
	return nil
}

func (r *repository) GetAssignmentById(id int) (assignment.Assignment, error) {
	var assign assignment.Assignment
	var user auth.User
	r.DB.Joins("JOIN users ON users.id = assignments.user_id").Where("assignments.id = ?", id).First(&assign)
	r.DB.Where("id = ?", assign.UserId).First(&user)
	assign.Name = user.Name
	return assign, nil
}

func (r *repository) GetAllAssignment() ([]assignment.Assignment, error) {
	assignments := []assignment.Assignment{}
	r.DB.Find(&assignments)
	return assignments, nil
}

func (r *repository) UpdateAssignment(id int, assignment assignment.Assignment) (assignment.Assignment, error) {
	var user auth.User
	r.DB.Model(&assignment).Updates(assignment)
	r.DB.Joins("JOIN users ON users.id = assignments.user_id").Where("assignments.id = ?", id).First(&assignment)
	r.DB.Where("id = ?", assignment.UserId).First(&user)
	assignment.Name = user.Name
	return assignment, nil
}

func (r *repository) DeleteAssignment(id int) (assignment.Assignment, error) {
	var assign assignment.Assignment
	r.DB.Where("id = ?", id).Delete(&assign)
	return assign, nil
}
