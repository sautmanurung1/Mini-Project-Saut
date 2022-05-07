package assignment

import (
	"Tugas-Mini-Project/infrastructure/database"
	m "Tugas-Mini-Project/infrastructure/http/middleware"
	"Tugas-Mini-Project/repository/assignment"
	service "Tugas-Mini-Project/service/assignment"
	"github.com/labstack/echo/v4"
)

func Routes(echo *echo.Echo, conf database.Config) {
	db := database.InitDB(conf)

	repo := assignment.NewAssignmentRepository(db)
	svc := service.NewAssignmentService(repo, conf)
	teacher := echo.Group("/teacher")
	controller := AssignmentHandler{
		svc: svc,
	}

	teacher.POST("/assignment", controller.CreateAssignmentHandler, m.JWTTeacherMiddleware())
	echo.GET("/assignment", controller.GetAllAssignmentHandler)
	echo.GET("/assignment/:id", controller.GetAssignmentByIdHandler)
	teacher.PUT("/assignment/:id", controller.UpdateAssignmentHandler, m.JWTTeacherMiddleware())
	teacher.DELETE("/assignment/:id", controller.DeleteAssignmentHandler, m.JWTTeacherMiddleware())
}
