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

	teacher.POST("/questions", controller.CreateQuestion, m.JWTTokenMiddleware())
	echo.GET("/questions", controller.GetAllQuestions)
	echo.GET("/questions/:id", controller.GetQuestionById)
	teacher.PUT("/questions/:id", controller.UpdateQuestion, m.JWTTokenMiddleware())
	teacher.DELETE("/questions/:id", controller.DeleteQuestion, m.JWTTokenMiddleware())
}
