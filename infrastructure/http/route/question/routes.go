package question

import (
	"Tugas-Mini-Project/infrastructure/database"
	m "Tugas-Mini-Project/infrastructure/http/middleware"
	ur "Tugas-Mini-Project/internal/handler/question"
	"Tugas-Mini-Project/internal/repository/question"
	"github.com/labstack/echo/v4"
)

func Routes(echo *echo.Echo) {
	db := database.InitDB()

	repo := question.NewQuestionRepository(db)
	handler := ur.NewQuestionHandler(repo)

	teacher := echo.Group("/teacher")
	teacher.POST("/questions", handler.CreateQuestionHandler, m.JWTTeacherMiddleware())
	echo.GET("/questions", handler.GetAllQuestionsHandler)
	echo.GET("/questions/:id", handler.GetQuestionByIdHandler)
	teacher.PUT("/questions/:id", handler.UpdateQuestionHandler, m.JWTTeacherMiddleware())
	teacher.DELETE("/questions/:id", handler.DeleteQuestionHandler, m.JWTTeacherMiddleware())
}
