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
	user := model.User{}
	if assignment.UserId == 1 {
		r.DB.Create(&assignment)
		assignment.User = user
	} else {
		return nil
	}
	return nil
}

func (r *repository) GetAssignmentById(id int) (model.Assignment, error) {
	var assign model.Assignment
	var user model.User
	r.DB.Joins("JOIN users ON users.id = assignments.user_id").Where("assignments.id = ?", id).First(&assign)
	r.DB.Where("id = ?", assign.UserId).First(&user)
	assign.User = user
	return assign, nil
}

func (r *repository) GetAllAssignment() ([]model.Assignment, error) {
	var assignments []model.Assignment
	r.DB.Joins("JOIN users ON users.id = assignments.user_id").Find(&assignments)
	return assignments, nil
}

func (r *repository) UpdateAssignment(assignment model.Assignment) error {
	var assign model.Assignment
	var user model.User
	r.DB.Joins("JOIN users ON users.id = assignments.user_id").Where("assignments.id = ?", assignment.ID).First(&assign)
	r.DB.Where("id = ?", assign.UserId).First(&user)
	if assign.UserId == 1 {
		r.DB.Model(&assign).Updates(assignment)
		r.DB.Where("id = ?", assign.UserId).First(&user)
		assign.User = user
	} else {
		return nil
	}
	return nil
}

func (r *repository) DeleteAssignment(id int) error {
	var assign model.Assignment
	r.DB.Joins("JOIN users ON users.id = assignments.user_id").Where("assignments.id = ?", id).First(&assign)
	if assign.UserId == 1 {
		r.DB.Where("id = ?", id).Delete(&assign)
	} else {
		return nil
	}
	return nil
}
