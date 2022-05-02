package auth

import (
	"Tugas-Mini-Project/infrastructure/database"
	"Tugas-Mini-Project/repository/auth"
	"github.com/labstack/echo/v4"
)

func Routes(echo *echo.Echo) {
	db := database.InitDB()

	repo := auth.NewAuthRepository(db)
	handler := NewAuthHandler(repo)

	echo.POST("/register", handler.RegisterHandler)
	echo.POST("/login", handler.LoginHandler)
}