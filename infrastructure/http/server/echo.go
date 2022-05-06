package server

import (
	"Tugas-Mini-Project/infrastructure/database"
	"Tugas-Mini-Project/interface/answare"
	"Tugas-Mini-Project/interface/assignment"
	"Tugas-Mini-Project/interface/auth"
	"Tugas-Mini-Project/interface/question"
	"Tugas-Mini-Project/interface/role"
	"github.com/labstack/echo/v4"
)

func Server() *echo.Echo {
	app := echo.New()
	conf := database.Config{}

	auth.Routes(app, conf)
	role.Routes(app, conf)
	assignment.Routes(app, conf)
	question.Routes(app, conf)
	answare.Routes(app, conf)
	return app
}
