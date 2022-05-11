package question

import (
	"Tugas-Mini-Project/infrastructure/database"
	m "Tugas-Mini-Project/infrastructure/http/middleware"
	questions "Tugas-Mini-Project/interface/question"
	"Tugas-Mini-Project/repository/question"
	service "Tugas-Mini-Project/service/question"
	"github.com/labstack/echo/v4"
)

func Routes(echo *echo.Echo, conf database.Config) {
	db := database.InitDB(conf)

	repo := question.NewQuestionRepository(db)
	svc := service.NewQuestionService(repo, conf)
	teacher := echo.Group("/teacher")
	controller := questions.QuestionHandler{
		Svc: svc,
	}

	teacher.POST("/questions", controller.CreateQuestionHandler, m.JWTTeacherMiddleware())
	echo.GET("/questions", controller.GetAllQuestionsHandler)
	echo.GET("/questions/:id", controller.GetQuestionByIdHandler)
	teacher.PUT("/questions/:id", controller.UpdateQuestionHandler, m.JWTTeacherMiddleware())
	teacher.DELETE("/questions/:id", controller.DeleteQuestionHandler, m.JWTTeacherMiddleware())
}
