package server

import (
	docs "Tugas-Mini-Project/docs"
	"Tugas-Mini-Project/infrastructure/database"
	"Tugas-Mini-Project/infrastructure/http/routes/answare"
	"Tugas-Mini-Project/infrastructure/http/routes/assignment"
	"Tugas-Mini-Project/infrastructure/http/routes/auth"
	"Tugas-Mini-Project/infrastructure/http/routes/discussions"
	"Tugas-Mini-Project/infrastructure/http/routes/question"
	"Tugas-Mini-Project/infrastructure/http/routes/role"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"os"
)

// @title School Discussions Room API Documentation
// @description This is School Discussions Room API
// @version 2.0
// @host localhost:8080
// @BasePath
// @schemes http https
// @securityDefinitions.apiKey JWT
// @in header
// @name Authorization
func Server() *echo.Echo {
	app := echo.New()
	conf := database.Config{}

	auth.Routes(app, conf)
	role.Routes(app, conf)
	assignment.Routes(app, conf)
	question.Routes(app, conf)
	answare.Routes(app, conf)
	discussions.Routes(app, conf)
	app.GET("/swagger/*", echoSwagger.WrapHandler)

	docs.SwaggerInfo.Host = os.Getenv("APP_HOST")
	return app
}
