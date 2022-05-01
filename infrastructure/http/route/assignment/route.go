package assignment

import (
	"Tugas-Mini-Project/infrastructure/database"
	m "Tugas-Mini-Project/infrastructure/http/middleware"
	ur "Tugas-Mini-Project/internal/handler/assignment"
	"Tugas-Mini-Project/internal/repository/assignment"
	"github.com/labstack/echo/v4"
)

func Routes(echo *echo.Echo) {
	db := database.InitDB()

	repo := assignment.NewAssignmentRepository(db)
	handler := ur.NewAssignmentHandler(repo)

	teacher := echo.Group("/teacher")

	teacher.POST("/assignment", handler.CreateAssignmentHandler, m.JWTTeacherMiddleware())
	echo.GET("/assignment", handler.GetAllAssignmentHandler)
	echo.GET("/assignment/:id", handler.GetAssignmentByIdHandler)
	teacher.PUT("/assignment/:id", handler.UpdateAssignmentHandler, m.JWTTeacherMiddleware())
	teacher.DELETE("/assignment/:id", handler.DeleteAssignmentHandler, m.JWTTeacherMiddleware())
}
