package auth

import (
	"Tugas-Mini-Project/infrastructure/database"
	ur "Tugas-Mini-Project/internal/handler/auth"
	"Tugas-Mini-Project/internal/repository/auth"
	"github.com/labstack/echo/v4"
)

func Routes(echo *echo.Echo) {
	db := database.InitDB()

	repo := auth.NewAuthRepository(db)
	handler := ur.NewAuthHandler(repo)

	echo.POST("/register", handler.RegisterHandler)
	echo.POST("/login", handler.LoginHandler)
}
