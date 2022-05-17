package auth

import (
	"Tugas-Mini-Project/infrastructure/database"
	auths "Tugas-Mini-Project/interface/auth"
	"Tugas-Mini-Project/repository/auth"
	service "Tugas-Mini-Project/service/auth"
	"github.com/labstack/echo/v4"
)

func Routes(echo *echo.Echo, conf database.Config) {
	db := database.InitDB(conf)

	repo := auth.NewAuthRepository(db)
	svc := service.NewAuthService(repo, conf)

	controller := auths.AuthHandler{
		Svc: svc,
	}

	echo.POST("/register", controller.Register)
	echo.POST("/login", controller.Login)
}
