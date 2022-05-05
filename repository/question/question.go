package question

import (
	"Tugas-Mini-Project/domains"
	"Tugas-Mini-Project/entities"
	"database/sql"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) domains.QuestionRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) CreateQuestion(question entities.Question) error {
	var assignments entities.Assignment
	var users entities.User
	r.DB.Where("id = ?", question.AssignmentId).First(&assignments)
	r.DB.Where("id = ?", question.UserId).First(&users)
	r.DB.Raw("SELECT title FROM assignments WHERE title = @title",
		sql.Named("title", assignments.Title)).Find(&assignments)
	question.Name = assignments.Name
	question.AssignmentTitle = assignments.Title
	r.DB.Create(&question)
	return nil
}

func (r *repository) GetQuestionByID(id int) (entities.Question, error) {
	var assign entities.Assignment
	var questions entities.Question
	r.DB.Joins("JOIN assignments ON assignments.id = questions.assignment_id").Where("questions.id = ?", id).First(&questions)
	r.DB.Where("id = ?", questions.AssignmentId).First(&assign)
	questions.Name = assign.Name
	return questions, nil
}

func (r *repository) GetAllQuestion() ([]entities.Question, error) {
	questions := []entities.Question{}
	r.DB.Find(&questions)
	return questions, nil
}

func (r *repository) UpdateQuestion(id int, question entities.Question) (entities.Question, error) {
	var assignments entities.Assignment
	r.DB.Model(&question).Where("id = ?", id).Updates(&question)
	r.DB.Joins("JOIN assignments ON assignments.id = questions.assignment_id").Where("questions.id = ?", id).First(&question)
	r.DB.Where("id = ?", question.AssignmentId).First(&assignments)
	question.Name = assignments.Name
	return question, nil
}

func (r *repository) DeleteQuestion(id int) (entities.Question, error) {
	var questions entities.Question
	r.DB.Where("id = ?", id).Delete(&questions)
	return questions, nil
}
