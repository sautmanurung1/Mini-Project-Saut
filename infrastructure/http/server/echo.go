package server

import (
	"Tugas-Mini-Project/interface/answare"
	"Tugas-Mini-Project/interface/assignment"
	"Tugas-Mini-Project/interface/auth"
	"Tugas-Mini-Project/interface/question"
	"Tugas-Mini-Project/interface/role"
	"github.com/labstack/echo/v4"
)

func Server() *echo.Echo {
	app := echo.New()
	auth.Routes(app)
	role.Routes(app)
	assignment.Routes(app)
	question.Routes(app)
	answare.Routes(app)
	return app
}
