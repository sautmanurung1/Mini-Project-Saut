package domains

import (
	"Tugas-Mini-Project/entities"
)

type QuestionRepository interface {
	CreateQuestion(question entities.Question) error
	GetQuestionByID(id int) (entities.Question, error)
	GetAllQuestion() ([]entities.Question, error)
	UpdateQuestion(id int, question entities.Question) (entities.Question, error)
	DeleteQuestion(id int, question entities.Question) error
}

type QuestionService interface {
	CreateQuestionService(question entities.Question) (string, error)
	GetQuestionByIDService(id int) (entities.Question, error)
	GetAllQuestionService() (question []entities.Question, err error)
	UpdateQuestionService(id int, question entities.Question) (entities.Question, error)
	DeleteQuestionService(id int, question entities.Question) (string, error)
}
