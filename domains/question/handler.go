package question

import "github.com/labstack/echo/v4"

type QuestionHandler interface {
	CreateQuestionHandler(c echo.Context) error
	GetQuestionByIdHandler(c echo.Context) error
	GetAllQuestionsHandler(c echo.Context) error
	UpdateQuestionHandler(c echo.Context) error
	DeleteQuestionHandler(c echo.Context) error
}
