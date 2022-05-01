package question

import "Tugas-Mini-Project/internal/entities"

type QuestionRepository interface {
	CreateQuestion(question entities.Question) error
	GetQuestionByID(id int) (entities.Question, error)
	GetAllQuestion() ([]entities.Question, error)
	UpdateQuestion(id int, question entities.Question) (entities.Question, error)
	DeleteQuestion(id int) (entities.Question, error)
}
