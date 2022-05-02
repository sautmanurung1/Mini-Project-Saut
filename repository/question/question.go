package question

import (
	"Tugas-Mini-Project/domains/assignment"
	"Tugas-Mini-Project/domains/auth"
	"Tugas-Mini-Project/domains/question"
	"database/sql"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) question.QuestionRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) CreateQuestion(question question.Question) error {
	var assignments assignment.Assignment
	var users auth.User
	r.DB.Where("id = ?", question.AssignmentId).First(&assignments)
	r.DB.Where("id = ?", question.UserId).First(&users)
	r.DB.Raw("SELECT title FROM assignments WHERE title = @title",
		sql.Named("title", assignments.Title)).Find(&assignments)
	question.Name = assignments.Name
	question.AssignmentTitle = assignments.Title
	r.DB.Create(&question)
	return nil
}

func (r *repository) GetQuestionByID(id int) (question.Question, error) {
	var assign assignment.Assignment
	var questions question.Question
	r.DB.Joins("JOIN assignments ON assignments.id = questions.assignment_id").Where("questions.id = ?", id).First(&questions)
	r.DB.Where("id = ?", questions.AssignmentId).First(&assign)
	questions.Name = assign.Name
	return questions, nil
}

func (r *repository) GetAllQuestion() ([]question.Question, error) {
	questions := []question.Question{}
	r.DB.Find(&questions)
	return questions, nil
}

func (r *repository) UpdateQuestion(id int, question question.Question) (question.Question, error) {
	var assignments assignment.Assignment
	r.DB.Model(&question).Where("id = ?", id).Updates(&question)
	r.DB.Joins("JOIN assignments ON assignments.id = questions.assignment_id").Where("questions.id = ?", id).First(&question)
	r.DB.Where("id = ?", question.AssignmentId).First(&assignments)
	question.Name = assignments.Name
	return question, nil
}

func (r *repository) DeleteQuestion(id int) (question.Question, error) {
	var questions question.Question
	r.DB.Where("id = ?", id).Delete(&questions)
	return questions, nil
}
