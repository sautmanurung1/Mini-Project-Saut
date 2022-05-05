package auth

import (
	"Tugas-Mini-Project/infrastructure/database"
	"Tugas-Mini-Project/repository/auth"
	"github.com/labstack/echo/v4"
)

func Routes(echo *echo.Echo, conf database.Config) {
	db := database.InitDB(conf)

	repo := auth.NewAuthRepository(db)
	handler := NewAuthHandler(repo)
	// svc := service.NewAuthService(repo, conf)
	echo.POST("/register", handler.RegisterHandler)
	echo.POST("/login", handler.LoginHandler)
}
