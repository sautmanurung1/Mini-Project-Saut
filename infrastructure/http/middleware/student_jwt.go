package middleware

import (
	"Tugas-Mini-Project/infrastructure/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTStudentMiddleware() echo.MiddlewareFunc {
	secret := database.Config{}
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    []byte(secret.Login_Student),
		SigningMethod: "HS256",
	})
}
