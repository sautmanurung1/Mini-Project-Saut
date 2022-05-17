package assignment

import (
	"Tugas-Mini-Project/infrastructure/database"
	m "Tugas-Mini-Project/infrastructure/http/middleware"
	assignments "Tugas-Mini-Project/interface/assignment"
	"Tugas-Mini-Project/repository/assignment"
	service "Tugas-Mini-Project/service/assignment"
	"github.com/labstack/echo/v4"
)

func Routes(echo *echo.Echo, conf database.Config) {
	db := database.InitDB(conf)

	repo := assignment.NewAssignmentRepository(db)
	svc := service.NewAssignmentService(repo, conf)
	teacher := echo.Group("/teacher")
	controller := assignments.AssignmentHandler{
		Svc: svc,
	}

	teacher.POST("/assignment", controller.CreateAssignment, m.JWTTeacherMiddleware())
	echo.GET("/assignment", controller.GetAllAssignment)
	echo.GET("/assignment/:id", controller.GetAssignmentById)
	teacher.PUT("/assignment/:id", controller.UpdateAssignment, m.JWTTeacherMiddleware())
	teacher.DELETE("/assignment/:id", controller.DeleteAssignment, m.JWTTeacherMiddleware())
}
