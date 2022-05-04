package domains

import (
	"Tugas-Mini-Project/entities"
	"github.com/labstack/echo/v4"
)

type QuestionRepository interface {
	CreateQuestion(question entities.Question) error
	GetQuestionByID(id int) (entities.Question, error)
	GetAllQuestion() ([]entities.Question, error)
	UpdateQuestion(id int, question entities.Question) (entities.Question, error)
	DeleteQuestion(id int) (entities.Question, error)
}

type QuestionHandler interface {
	CreateQuestionHandler(c echo.Context) error
	GetQuestionByIdHandler(c echo.Context) error
	GetAllQuestionsHandler(c echo.Context) error
	UpdateQuestionHandler(c echo.Context) error
	DeleteQuestionHandler(c echo.Context) error
}
