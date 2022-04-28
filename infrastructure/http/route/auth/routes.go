package auth

import (
	ur "Tugas-Mini-Project/handler/auth"
	"Tugas-Mini-Project/infrastructure/database"
	"Tugas-Mini-Project/repository/auth"
	"github.com/labstack/echo/v4"
)

func Routes(echo *echo.Echo) {
	db := database.InitDB()

	repo := auth.NewAuthRepository(db)
	handler := ur.NewAuthHandler(repo)

	teacher := echo.Group("/teacher")
	student := echo.Group("/student")

	student.POST("/register", handler.RegisterHandler)
	teacher.POST("/register", handler.RegisterHandler)
	student.GET("/login", handler.LoginHandler)
	teacher.GET("/login", handler.LoginHandler)
}
