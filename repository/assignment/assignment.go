package assignment

import (
	"Tugas-Mini-Project/domains"
	"Tugas-Mini-Project/entities"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewAssignmentRepository(db *gorm.DB) domains.AssignmentRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) CreateAssignment(assignment entities.Assignment) error {
	var user entities.User
	r.DB.Where("id = ?", assignment.UserId).First(&user)
	assignment.Name = user.Name
	r.DB.Create(&assignment)
	return nil
}

func (r *repository) GetAssignmentById(id int) (entities.Assignment, error) {
	var assign entities.Assignment
	var user entities.User
	r.DB.Joins("JOIN users ON users.id = assignments.user_id").Where("assignments.id = ?", id).First(&assign)
	r.DB.Where("id = ?", assign.UserId).First(&user)
	assign.Name = user.Name
	return assign, nil
}

func (r *repository) GetAllAssignment() ([]entities.Assignment, error) {
	var assignments []entities.Assignment
	r.DB.Find(&assignments)
	return assignments, nil
}

func (r *repository) UpdateAssignment(id int, assignment entities.Assignment) (entities.Assignment, error) {
	var user entities.User
	r.DB.Model(&assignment).Where("id = ?", id).Updates(assignment)
	r.DB.Joins("JOIN users ON users.id = assignments.user_id").Where("assignments.id = ?", id).First(&assignment)
	r.DB.Where("id = ?", assignment.UserId).First(&user)
	assignment.Name = user.Name
	return assignment, nil
}

func (r *repository) DeleteAssignment(id int, assign entities.Assignment) error {
	r.DB.Where("id = ?", id).Delete(&assign)
	return nil
}
