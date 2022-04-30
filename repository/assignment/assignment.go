package assignment

import (
	"Tugas-Mini-Project/domains/assignment"
	"Tugas-Mini-Project/model"
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

func (r *repository) CreateAssignment(assignment model.Assignment) error {
	// query to get the name from table user to insert name to table assignment
	var user model.User
	r.DB.Where("id = ?", assignment.UserId).First(&user)
	assignment.Name = user.Name
	r.DB.Create(&assignment)
	// make query to get user name from user id in assignment
	// r.DB.Where("id = ?", assignment.UserId).First(&user)
	// assignment.User = user
	return nil
}

func (r *repository) GetAssignmentById(id int) (model.Assignment, error) {
	var assign model.Assignment
	var user model.User
	r.DB.Joins("JOIN users ON users.id = assignments.user_id").Where("assignments.id = ?", id).First(&assign)
	r.DB.Where("id = ?", assign.UserId).First(&user)
	assign.Name = user.Name
	return assign, nil
}

func (r *repository) GetAllAssignment() ([]model.Assignment, error) {
	assignments := []model.Assignment{}
	r.DB.Find(&assignments)
	return assignments, nil
}

func (r *repository) UpdateAssignment(id int, assignment model.Assignment) (model.Assignment, error) {
	var user model.User
	r.DB.Model(&assignment).Updates(assignment)
	r.DB.Joins("JOIN users ON users.id = assignments.user_id").Where("assignments.id = ?", id).First(&assignment)
	r.DB.Where("id = ?", assignment.UserId).First(&user)
	assignment.Name = user.Name
	return assignment, nil
}

func (r *repository) DeleteAssignment(id int) (model.Assignment, error) {
	// r.DB.Joins("JOIN users ON users.id = assignments.user_id").Where("assignments.id = ?", id).First(&assign)
	// delete assignment from table assignment where id = id parameter
	var assign model.Assignment
	r.DB.Where("id = ?", id).Delete(&assign)
	return assign, nil
}
