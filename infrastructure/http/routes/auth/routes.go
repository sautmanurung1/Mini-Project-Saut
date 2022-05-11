package auth

import (
	"Tugas-Mini-Project/infrastructure/database"
	auth2 "Tugas-Mini-Project/interface/auth"
	"Tugas-Mini-Project/repository/auth"
	service "Tugas-Mini-Project/service/auth"
	"github.com/labstack/echo/v4"
)

func Routes(echo *echo.Echo, conf database.Config) {
	db := database.InitDB(conf)

	repo := auth.NewAuthRepository(db)
	svc := service.NewAuthService(repo, conf)

	controller := auth2.AuthHandler{
		Svc: svc,
	}

	echo.POST("/register", controller.RegisterHandler)
	echo.POST("/login", controller.LoginHandler)
}
